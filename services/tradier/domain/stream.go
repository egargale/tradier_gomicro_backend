package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"nhooyr.io/websocket"
	// "nhooyr.io/websocket/wsjson"
	// "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"tradier_gomicro_backend/services/tradier/subscriber"
	"github.com/micro/micro/v3/service"
)

var ctx = context.Background()

const (
	// fasthhtp Maximum number of connections per each host which may be established.
	MaxConnsPerHost = 100
	//  fasthhtp Maximum number of attempts for idempotent calls
	MaxIdemponentCallAttempts = 10
	//
	defaultRetries = 3
	// Header indicating the number of requests remaining.
	rateLimitAvailable = "X-Ratelimit-Available"
	// Header indicating the time at which our rate limit will renew.
	rateLimitExpiry = "X-Ratelimit-Expiry"
	// Error returned by Tradier if we make too big of a request.
	ErrBodyBufferOverflow = "protocol.http.TooBigBody"
	//
	SandboxEndpoint = "https://sandbox.tradier.com"
	APIEndpoint     = "https://api.tradier.com"
	StreamEndpoint  = "https://stream.tradier.com"
)

var (
	// ErrNoAccountSelected is returned if account-specific methods
	// are attempted to be used without selecting an account first.
	ErrNoAccountSelected = errors.New("no account selected")
)
var SocketConfig Stream

type Stream struct {
	SessionId string
	Url       string
}

type ClientParams struct {
	Endpoint   string
	AuthToken  string
	Client     *fasthttp.Client
	RetryLimit int
	Account    string
}

// DefaultParams returns ClientParams initialized with default values.
func DefaultParams(authToken string) ClientParams {
	return ClientParams{
		Endpoint:   APIEndpoint,
		AuthToken:  authToken,
		Client:     &fasthttp.Client{},
		RetryLimit: defaultRetries,
	}
}

// Client provides methods for making requests to the Tradier API.
type Client struct {
	client     *fasthttp.Client
	endpoint   string
	authHeader string
	retryLimit int
	account    string
}

func NewClient(params ClientParams) *Client {
	return &Client{
		client:     params.Client,
		endpoint:   params.Endpoint,
		authHeader: fmt.Sprintf("Bearer %s", params.AuthToken),
		retryLimit: params.RetryLimit,
		account:    params.Account,
	}
}

func (tc *Client) SelectAccount(account string) {
	tc.account = account
}

func (tc *Client) GetSessionID() (Stream, error) {

	// First create a streaming session.
	createSessionUrl := tc.endpoint + "/v1/markets/events/session"

	a := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(a)

	r := fiber.AcquireResponse()
	defer fiber.ReleaseResponse(r)

	a.Reuse()

	//Set request
	req := a.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.Header.Set("Authorization", tc.authHeader)
	req.Header.Set("accept", "application/json")
	req.SetRequestURI(createSessionUrl)

	if err := a.Parse(); err != nil {
		fmt.Println(err.Error())
	}

	// var retCode int
	var retBody []byte
	var errs []error

	// Send out the HTTP request
	var sessionResp struct {
		Stream struct {
			SessionId string
			Url       string
		}
	}

	if _, retBody, errs = a.Struct(&sessionResp); len(errs) > 0 {
		log.Printf("received: %v", string(retBody))
		log.Printf("could not send HTTP request: %v", errs)
		return sessionResp.Stream, fmt.Errorf("Error in socket")
	}
	log.Println(sessionResp.Stream.SessionId)
	SocketConfig.SessionId = sessionResp.Stream.SessionId
	SocketConfig.Url = sessionResp.Stream.Url

	// go OpenStreamSocket(sessionResp.Stream.SessionId)
	return sessionResp.Stream, nil
}

func OpenStreamSocket(id string) {
	p := service.NewEvent(subscriber.Topic)
	
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	// defer cancel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, _, err := websocket.Dial(ctx, "wss://ws.tradier.com/v1/markets/events", nil)
	if err != nil {
		//log.Printf("handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial:", err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	type streamPayload struct {
		Symbols   []string `json:"symbols"`
		SessionId string   `json:"sessionid"`
		Linebreak bool     `json:"linebreak"`
	}

	p2 := streamPayload{
		Symbols:   []string{"SPY", "GM"},
		SessionId: id,
		Linebreak: true,
	}

	u, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}
	// log.Println(string(u))

	fakepayload := fmt.Sprintf("{\"symbols\": [\"SPY\"], \"sessionid\": \"%s\", \"linebreak\": true}", id)
	// OK
	// err = c.Write(ctx, websocket.MessageText, []byte(fakepayload))
	// OK
	err = c.Write(ctx, websocket.MessageText, u)
	// NO to be fixed
	//err = wsjson.Write(ctx, c, p2)
	if err != nil {
		log.Printf("Error in subscribing stream %s", err)
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.Read(ctx)
			if err != nil {
				log.Println("Error reading socket:", err)
				return
			}
			if err := p.Publish(ctx, message); err != nil {
				panic(err)
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			c.Close(websocket.StatusNormalClosure, "")
			return
		case <-ticker.C:
			// log.Printf("time: %s", t.String())
			err = c.Write(ctx, websocket.MessageText, []byte(fakepayload))
			// err := c.Write(ctx, websocket.MessageText, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}
