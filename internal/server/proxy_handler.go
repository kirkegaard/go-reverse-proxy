package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProxy(target *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy, url *url.URL, endpoint string) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("<= %s/%s\n", req.Host, req.URL)

		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Origin-Host", url.Host)

		req.Host = url.Host

		fmt.Printf("=> %s/%s\n", url.Host, req.URL)
		proxy.ServeHTTP(res, req)
	}
}
