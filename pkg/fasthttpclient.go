package pkg

import (
	"time"

	"github.com/attapon-th/template-fiber-api/helper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

// FasthttpClient Fasthttp client
type FasthttpClient struct {
	BaseURI       string
	TransactionID string            // default time.Now().String()
	UserAgent     string            // default testAgent
	ContentType   string            // default application/json
	Headers       map[string]string // default empty
	Accept        string            // default application/json
	TimeOut       time.Duration     // default 10s
	log           zerolog.Logger    // default log
	Debug         bool              // default false
}

// NewFasthttpClient Get and setup a default fasthttp client
func NewFasthttpClient() *FasthttpClient {
	// From a zapcore.Core, it's easy to construct a Logger.
	return &FasthttpClient{
		TransactionID: time.Now().String(),
		UserAgent:     "Fasthttp-Agent",
		ContentType:   "application/json; charset=utf-8",
		Accept:        "application/json",
		Headers:       make(map[string]string),
		TimeOut:       10 * time.Second,
		Debug:         false,
		log:           log.Logger,
	}
}

// FasthttpByte  do  POST request via fasthttp
func (w *FasthttpClient) FasthttpByte(requestURI string, method string, body []byte) (*fasthttp.Response, error) {
	t1 := time.Now()
	w.TransactionID = t1.String()
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	req.SetRequestURI(requestURI)
	req.Header.SetContentType(w.ContentType)
	req.Header.Add("User-Agent", w.UserAgent)
	req.Header.Add("TransactionID", w.TransactionID)
	req.Header.Add("Accept", w.Accept)

	req.Header.SetMethod(method)
	req.SetBody(body)
	// fmt.Println("---------- req --------------")
	ldebug := log.Debug().Str("transactionID", w.TransactionID)
	if w.Debug {
		ldebug.Str("method", method)
		req.Header.VisitAll(func(key, value []byte) {
			ldebug.
				Str("key", helper.B2S(key)).
				Str("value", helper.B2S(value))
			req.Header.VisitAll(func(key, value []byte) {
				ldebug.Str("key", helper.B2S(key)).Str("value", helper.B2S(value))
			})
		})
	}
	log.Debug().Msg("request")

	timeOut := 3 * time.Second
	if w.TimeOut != 0 {
		timeOut = w.TimeOut
	}
	err := fasthttp.DoTimeout(req, resp, timeOut)
	if err != nil {
		log.Error().Err(err).Msg("post request error")
		return nil, err
	}
	// list all response for debug
	ldebug = log.Debug().Str("transactionID", w.TransactionID)
	if w.Debug {
		elapsed := time.Since(t1)
		ldebug.Dur("elapsed", elapsed).Int("http status code", resp.StatusCode())
		resp.Header.VisitAll(func(key, value []byte) {
			ldebug.Str("key", helper.B2S(key)).Str("value", helper.B2S(value))
		})
		ldebug.Str("http payload", helper.B2S(resp.Body()))
	}
	ldebug.Msg("response")

	out := fasthttp.AcquireResponse()
	resp.CopyTo(out)

	return out, nil
}

// FastGet do GET request via fasthttp
func (w *FasthttpClient) FastGet(requestURI string) (*fasthttp.Response, error) {
	return w.FasthttpByte(requestURI, fasthttp.MethodGet, nil)
}

// FastPost do GET request via fasthttp
func (w *FasthttpClient) FastPost(requestURI string, body []byte) (*fasthttp.Response, error) {
	return w.FasthttpByte(requestURI, fasthttp.MethodGet, body)
}

// FastPatch do PATCH request via fasthttp
func (w *FasthttpClient) FastPatch(requestURI string, body []byte) (*fasthttp.Response, error) {
	return w.FasthttpByte(requestURI, fasthttp.MethodPatch, body)
}

// FastPut do PUT request via fasthttp
func (w *FasthttpClient) FastPut(requestURI string, body []byte) (*fasthttp.Response, error) {
	return w.FasthttpByte(requestURI, fasthttp.MethodPut, body)
}

// FastDelete do DELETE request via fasthttp
func (w *FasthttpClient) FastDelete(requestURI string) (*fasthttp.Response, error) {
	return w.FasthttpByte(requestURI, fasthttp.MethodDelete, nil)
}
