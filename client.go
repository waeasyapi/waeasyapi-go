package waeasyapi

import (
	"net/http"
	"time"

	"github.com/waeasyapi/waeasyapi-go/constants"
	"github.com/waeasyapi/waeasyapi-go/requests"
	"github.com/waeasyapi/waeasyapi-go/resources"
)

//Request ...
var Request *requests.Request

//Client provides various helper methods to make HTTP requests to WAEasyAPI's APIs.
type Client struct {
	Message        *resources.Message
}

// NewClient creates and returns a new WAEasyAPI client. key and secret
// are used to authenticate the requests made to WAEasyAPI's APIs.
func NewClient(key string, secret string) *Client {
	auth := requests.Auth{Key: key, Secret: secret}
	httpClient := &http.Client{Timeout: requests.TIMEOUT * time.Second}
	Request = &requests.Request{Auth: auth, HTTPClient: httpClient,
		Version: getVersion(), SDKName: getSDKName(),
		BaseURL: constants.BASE_URL}

	message := resources.Message{Request: Request}
	client := Client{
		Message:        &message,
	}
	return &client
}

// AddHeaders adds additional headers to WAEasyAPI's client. All requests
// made using the client will contain these additional headers in the HTTP request.
func (client *Client) AddHeaders(headers map[string]string) {
	Request.AddHeaders(headers)
}

// SetTimeout sets the timeout of WAEasyAPI's Client. The default timeout will
// be overridden for all HTTP requests made using this client.
func (client *Client) SetTimeout(timeout int16) {
	Request.SetTimeout(timeout)
}

func getVersion() string {
	return SDKVersion
}

func getSDKName() string {
	return SDKName
}