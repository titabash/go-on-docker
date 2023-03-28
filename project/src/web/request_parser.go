package web

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type RequestParser struct {
	body  map[string]interface{}
	query url.Values
}

func (rp *RequestParser) NewRequestParser(r *http.Request) {
	body := r.Body
	defer body.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, body)
	json.Unmarshal(buf.Bytes(), &rp.body)
	rp.query = r.URL.Query()
}
