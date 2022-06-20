package resources

import (
	"fmt"
    "net/url"
	
	"github.com/waeasyapi/waeasyapi-go/constants"
	"github.com/waeasyapi/waeasyapi-go/requests"
)

//Message ...
type Message struct {
	Request *requests.Request
}

// send a message
func (p *Payment) sendMessage(number string, message string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/message", constants.MESSAGE_URL, url.PathEscape(number))
	if data == nil {
		data = make(map[string]interface{})
	}
	data["message"] = message

	return p.Request.Post(url, data, extraHeaders)
}

// send a template message
func (p *Payment) sendTemplate(number string, template string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/template", constants.MESSAGE_URL, url.PathEscape(number))
	if data == nil {
		data = make(map[string]interface{})
	}
	data["template"] = template

	return p.Request.Post(url, data, extraHeaders)
}

// send a media message
func (p *Payment) sendMedia(number string, template string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/media", constants.MESSAGE_URL, url.PathEscape(number))
	if data == nil {
		data = make(map[string]interface{})
	}
	data["template"] = template

	return p.Request.Post(url, data, extraHeaders)
}