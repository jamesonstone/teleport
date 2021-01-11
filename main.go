package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handler struct{}

var (
	presetServiceMap = map[string]string{
		"/arctic":     "http://localhost:7000",
		"/icebreaker": "http://localhost:4000",
	}
)

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	servicePath := r.URL.Path
	log.Print(servicePath)

	if t, ok := presetServiceMap[servicePath]; ok {
		remoteUrl, err := url.Parse(t)
		if err != nil {
			log.Println("Failed to parse service path: ", err)
			return
		}

		log.Print("Preset Set")
		r.URL.Path = "/"
		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		proxy.ServeHTTP(w, r)
		return
	}

	w.Write([]byte("404: Application not found " + servicePath))
}

func main() {

	h := &handler{}
	http.Handle("/", h)

	server := &http.Server{
		Addr:    ":5050",
		Handler: h,
	}
	log.Fatal(server.ListenAndServe())
}
