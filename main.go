package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

type proxyHandler struct {
	p *httputil.ReverseProxy
}

func loadRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/{name}", handler)

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m := make(map[string]string)

	m["icebreaker"] = "localhost:4000"
	m["arctic"] = "localhost:7000"

	log.Print(m)

	log.Print(vars["name"])
	if val, ok := m[vars["name"]]; ok {
		u, _ := url.Parse("http://"+val)
		proxy := httputil.NewSingleHostReverseProxy(u)
		ph := &proxyHandler{proxy}
		log.Print(val)
		ph.ServeHTTP(w, r)
	} else {
		fmt.Fprintf(w, "Application: \"%v\" is not connected or does not exist.\n", vars["name"])
	}
}

func (ph *proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	ph.p.ServeHTTP(w, r)
}

func SessionHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r)
    })
}

func main() {
	http.ListenAndServe(":5050", loadRouter())
}

