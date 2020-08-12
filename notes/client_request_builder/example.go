package example

import (
	"fmt"
	"net/http"
	"net/url"
)

type ReadMessageRequest struct {
	OrgGUID   string
	Mailbox   string
	MessageID string
}

func (r ReadMessageRequest) Request(scheme, host string) *http.Request {
	return &http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: scheme,
			Host:   host,
			Path:   fmt.Sprintf("/api/v1/%s", r.OrgGUID),
			RawQuery: url.Values{
				"mailbox":    []string{r.Mailbox},
				"message_id": []string{r.MessageID},
			}.Encode(),
		},
	}
}

func ExampleMain() {
	req := ReadMessageRequest{
		OrgGUID:   "foo",
		Mailbox:   "bar",
		MessageID: "baz",
	}

	http.DefaultClient.Do(req.Request("http", "localhost:9090"))
}
