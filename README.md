# WA Easy API Go Client

Golang bindings for interacting with the WA Easy API

This is primarily meant for merchants who wish to perform interactions with the WA Easy API programatically

Read up here for getting started and understanding the messages flow with WA Easy API: <https://docs.waeasyapi.com/>

## Documentation

Documentation of WA Easy API's API and their usage is available at <https://docs.waeasyapi.com>

## Usage
You need to setup your key and secret using the following:
You can find your keys at <https://app.waeasyapi.com/>.

```go

go get github.com/waeasyapi/waeasyapi-go

```

Or
    

```go

import (waeasyapi "github.com/waeasyapi/waeasyapi-go")

```

```go

client := waeasyapi.NewClient("<YOUR_ACC_KEY>", "<YOUR_ACC_SECRET>")

```

Note: All methods return a `map[string]interface{}` and `error`

```go

// number must start with the country's dialing code

// send a message
number := "1982388224"
message := "Hello World"

body, err := client.message.SendMessage(number, message, nil)

```

```go

// number must start with the country's dialing code

// send a template message
number := "1982388224"
template := "template"

params := map[string]interface{}{
  "key1": "value1",
  "key2": "value2",
}
body, err := client.message.SendTemplate(number, template, params, nil)

```

```go

// number must start with the country's dialing code

// send a media message
number := "1982388224"
template := "template"

params := map[string]interface{}{
  "key1": "value1",
  "key2": "value2",
}
body, err := client.message.SendMedia(number, template, params, nil)
```

## License

The WA Easy API Go SDK is released under the MIT License. See [LICENSE](LICENSE) file for more details.
