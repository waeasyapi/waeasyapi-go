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
func (p *Message) SendMessage(number string, message string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/message"
	fmt.Println("HERE: ", url)
	if data == nil {
		data = make(map[string]interface{})
	}
	data["message"] = message

	return p.Request.Post(url, data, extraHeaders)
}

// send a template message
func (p *Message) SendTemplate(number string, template string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/template"
	if data == nil {
		data = make(map[string]interface{})
	}
	data["template"] = template

	return p.Request.Post(url, data, extraHeaders)
}

// send a media message
func (p *Message) SendMedia(number string, template string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/media"
	if data == nil {
		data = make(map[string]interface{})
	}
	data["template"] = template

	return p.Request.Post(url, data, extraHeaders)
}
