package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func NewProxy(target *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy, url *url.URL, endpoint string) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("[broxy] Request received at %s at %s\n", req.URL, time.Now())
		// Update the headers to allow for SSL redirection
		req.URL.Host = url.Host
		req.URL.Scheme = url.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = url.Host

		// trim reverseProxyRouterPrefix
		path := req.URL.Path
		req.URL.Path = strings.TrimLeft(path, endpoint)

		// Note that ServeHttp is non blocking and uses a go routine under the hood
		fmt.Printf("[broxy] Proxying request to %s at %s\n", req.URL, time.Now())
		proxy.ServeHTTP(res, req)
	}
}
