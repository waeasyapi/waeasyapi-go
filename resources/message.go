package resources

import (
	"github.com/waeasyapi/waeasyapi-go/constants"
	"github.com/waeasyapi/waeasyapi-go/requests"
)

//Message ...
type Message struct {
	Request *requests.Request
}

// send a text message
func (p *Message) SendTextMessage(number string, message string, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/text"
	data := map[string]interface{}{
		"number": number,
		"message": message,
	}

	return p.Request.Post(url, data, extraHeaders)
}

// send a url message
func (p *Message) SendURLMessage(number string, url string, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/url"
	data := map[string]interface{}{
		"number": number,
		"url": url,
	}

	return p.Request.Post(url, data, extraHeaders)
}

// send a template message
func (p *Message) SendTemplateMessage(number string, template string, params map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/template"
	data := map[string]interface{}{
		"number": number,
		"template": template,
		"params": params,
	}
	return p.Request.Post(url, data, extraHeaders)
}

// send a image message
func (p *Message) SendImageMessage(number string, params map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/image"
	data := map[string]interface{}{
		"number": number,
		"params": params,
	}
	return p.Request.Post(url, data, extraHeaders)
}

// send a video message
func (p *Message) SendVideoMessage(number string, params map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/video"
	data := map[string]interface{}{
		"number": number,
		"params": params,
	}
	return p.Request.Post(url, data, extraHeaders)
}

// send a audio message
func (p *Message) SendAudioMessage(number string, params map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/audio"
	data := map[string]interface{}{
		"number": number,
		"params": params,
	}
	return p.Request.Post(url, data, extraHeaders)
}

// send a document message
func (p *Message) SendDocumentMessage(number string, params map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/document"
	data := map[string]interface{}{
		"number": number,
		"params": params,
	}
	return p.Request.Post(url, data, extraHeaders)
}

// send a contact message
func (p *Message) SendContactMessage(number string, params map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/contacts"
	data := map[string]interface{}{
		"number": number,
		"params": params,
	}
	return p.Request.Post(url, data, extraHeaders)
}

// send a contact message
func (p *Message) SendLocationMessage(number string, params map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/location"
	data := map[string]interface{}{
		"number": number,
		"params": params,
	}
	return p.Request.Post(url, data, extraHeaders)
}

// send an interactive message
func (p *Message) SendLocationMessage(number string, params map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := constants.MESSAGE_URL + "/interactive"
	data := map[string]interface{}{
		"number": number,
		"params": params,
	}
	return p.Request.Post(url, data, extraHeaders)
}