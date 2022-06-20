package resources

import (
	"github.com/waeasyapi/waeasyapi-go/constants"
	"github.com/waeasyapi/waeasyapi-go/requests"
)

//Message ...
type Message struct {
	Request *requests.Request
}

// send a message
func (p *Message) SendMessage(number string, message string, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/message"
	data := map[string]interface{}{
		"number": number,
		"message": message,
	}

	return p.Request.Post(url, data, extraHeaders)
}

// send a template message
func (p *Message) SendTemplate(number string, template string, params map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/template"
	data := map[string]interface{}{
		"number": number,
		"template": template,
		"params": params,
	}
	return p.Request.Post(url, data, extraHeaders)
}

// send a media message
func (p *Message) SendMedia(number string, template string, params map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/media"
	data := map[string]interface{}{
		"number": number,
		"template": template,
		"params": params,
	}
	return p.Request.Post(url, data, extraHeaders)
}
