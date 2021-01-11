package not3main

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

func (ph *proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	ph.p.ServeHTTP(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {

	if val, ok := m[r.URL.Path]; ok {
		log.Print(val)
		u, _ := url.Parse("http://" + val)
		proxy := httputil.NewSingleHostReverseProxy(u)
		ph := &proxyHandler{proxy}
		ph.ServeHTTP(w, r)
	} else {
		fmt.Fprintf(w, "Application: \"%v\" is not connected or does not exist.\n", r.URL.Path)
	}
}

func loadRouter() *mux.Router {
	r := mux.NewRouter()

	// urlPath := mux.Vars(r)["name"]
	m := make(map[string]string)

	m["icebreaker"] = "localhost:4000"
	m["arctic"] = "localhost:7000"

	for k, v := range m {
		r.HandleFunc(k)
	}

	r.HandleFunc("/arctic", handler)
	return r
}

func main() {
	http.ListenAndServe(":5050", loadRouter())
}

func dynamicHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.URL.Path)
		if r.URL.Path == "/arctic" {
			u, _ := url.Parse("http://localhost:7000")
			proxy := httputil.NewSingleHostReverseProxy(u)
			ph := &proxyHandler{proxy}
			ph.ServeHTTP(w, r)
			log.Print("arctic")
		} else {
			log.Print("nothing to see here")
		}
	}

	return http.HandlerFunc(fn)
}

