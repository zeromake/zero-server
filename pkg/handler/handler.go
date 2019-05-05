package handler

import (
	"io"
	"net/http"
)

type Handler struct {
	transport http.RoundTripper
}

func (h *Handler) Handler(resp http.ResponseWriter, req *http.Request) {
	res, err := h.transport.RoundTrip(req)
	if err != nil {
		resp.WriteHeader(502)
		return
	}
	header := resp.Header()
	copyHeader(header, res.Header)
	if res.Body != nil {
		_, _ = io.Copy(resp, res.Body)
	}
}

func New(transport http.RoundTripper) *Handler {
	return &Handler{
		transport,
	}
}

func copyHeader(dst http.Header, src http.Header) {
	for k, v := range src {
		val := make([]string, len(v))
		copy(val, v)
		dst[k] = val
	}
}
