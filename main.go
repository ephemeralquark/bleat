/*
* example from: https://golang.org/pkg/net/http/httputil/#example_ReverseProxy
 */
package main

import (
	// "fmt"
	// "io/ioutil"
	"io"
	"log"
	"net/http"
	// "net/http/httptest"
	// "net/http/httputil"
	// "net/url"
)

// backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// fmt.Fprintln(w, "this call was relayed by the reverse proxy")
// }))
// defer backendServer.Close()

// rpURL, err := url.Parse(backendServer.URL)
// if err != nil {
// log.Fatal(err)
// }
// frontendProxy := httptest.NewServer(httputil.NewSingleHostReverseProxy(rpURL))
// defer frontendProxy.Close()

// resp, err := http.Get(frontendProxy.URL)
// if err != nil {
// log.Fatal(err)
// }

// b, err := ioutil.ReadAll(resp.Body)
// if err != nil {
// log.Fatal(err)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
