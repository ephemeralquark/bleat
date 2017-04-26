package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxyUrl, err := url.Parse("http://www.humana.io")

	if err != nil {
		log.Fatal("Could not parse the proxy url")
	}
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	proxy.Director = func(req *http.Request) {
		req.URL = proxyUrl
		req.Host = proxyUrl.Host
	}

	http.Handle("/", proxy)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
