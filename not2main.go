package not2main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	hostTarget = map[string]string{
		"/arctic": "http://localhost:7000",
		"/icebox": "http://localhost:4000",
	}
	hostProxy map[string]*httputil.ReverseProxy
)

type baseHandle struct{}

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    host := r.URL.Path
	log.Print(host)
    if target, ok := hostTarget[host]; ok {
        remoteUrl, err := url.Parse(target)
        if err != nil {
            log.Println("target parse fail:", err)
            return
        }

        proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
        proxy.ServeHTTP(w, r)
        return
    }
    w.Write([]byte("403: Host forbidden " + host))
}

func main() {

	h := &baseHandle{}
	http.Handle("/", h)

	server := &http.Server{
		Addr:    ":5050",
		Handler: h,
	}
	log.Fatal(server.ListenAndServe())
}
