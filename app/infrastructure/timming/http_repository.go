package timming

import (
	"bytes"
	"errors"
	"github.com/Timming/app/infrastructure/timming/client"
	"github.com/labstack/gommon/log"
	"net/http"
)

type HttpRepository struct {
	client  client.HttpClient
	url     string
	method  string
	headers map[string]string
	body    *bytes.Buffer
}

func NewHttpRepository(client client.HttpClient) *HttpRepository {
	if client == nil {
		panic("You must create client")
	}
	return &HttpRepository{
		client: client,
	}
}

func (r *HttpRepository) URL(url string) *HttpRepository {
	r.url = url
	return r
}
func (r *HttpRepository) Method(method string) *HttpRepository {
	r.method = method
	return r
}

func (r *HttpRepository) Headers(headers map[string]string) *HttpRepository {
	r.headers = headers
	return r
}

func (r *HttpRepository) Body(body *bytes.Buffer) *HttpRepository {
	r.body = body
	return r
}

func (r *HttpRepository) Do() (*http.Response, error) {

	if r.url == "" {
		return nil, errors.New("Empty URL")
	}
	if r.method == "" {
		return nil, errors.New("Empty Method")
	}
	if r.method != http.MethodPost && r.method != http.MethodGet && r.method != http.MethodPut && r.method != http.MethodDelete && r.method != http.MethodPatch {
		return nil, errors.New("Method not allowed")
	}

	request, err := http.NewRequest(r.method, r.url, nil)
	if err != nil {
		log.Errorf("error in creating request")
		return nil, err
	}
	if r.body != nil {
		request, _ = http.NewRequest(r.method, r.url, r.body)
	}

	if r.headers != nil {
		for key, value := range r.headers {
			request.Header.Add(key, value)
		}
	}
	resp, err := r.client.Do(request)
	return resp, err
}
