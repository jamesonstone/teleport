package not0main

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

func loadDynamicRoutes(r *mux.Router) {
	m := make(map[string]string)

	m["icebreaker"] = "localhost:4000"
	m["arctic"] = "localhost:7000"

	for name, url := range m {
		log.Print("key: ", name)
		log.Print("value: ", url)
		r.HandleFunc("/"+name, func(w http.ResponseWriter, r *http.Request) {
			handler(url, w, r)
		})
	}
}

func handler(serviceURL string, w http.ResponseWriter, r *http.Request) {

	log.Print("\n")
	log.Print(serviceURL)
	log.Print(w)
	log.Print(r)
	log.Print("\n")

	u, _ := url.Parse("http://" + serviceURL)
	proxy := httputil.NewSingleHostReverseProxy(u)
	ph := &proxyHandler{proxy}
	ph.ServeHTTP(w, r)
}

func loadRouter() *mux.Router {
	r := mux.NewRouter()

	loadDynamicRoutes(r)

	r.HandleFunc("/arctic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	return r
}

func main() {
	http.ListenAndServe(":5050", loadRouter())
}
