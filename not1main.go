package not1main

import (
    "log"
    "net/http"
    "regexp"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello world!"))
}

func adaptFileServer(fs http.Handler, mux http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, req *http.Request) {
        staticRegex := regexp.MustCompile("^/static-page-[0-9]+/")
        if matches := staticRegex.FindStringSubmatch(req.URL.Path); matches != nil {
            log.Printf("Match: %v, %v", req.URL.Path, matches[0])
            fsHandler := http.StripPrefix(matches[0], fs)
            fsHandler.ServeHTTP(w, req)
        } else if mux != nil {
            log.Printf("Doesn't match, pass to other MUX: %v", req.URL.Path)
            mux.ServeHTTP(w, req)
        } else {
            http.Error(w, "Page Not Found", http.StatusNotFound)
        }
    }
    return http.HandlerFunc(fn)
}

func init() {
    //Usual routing definition with MUX
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", helloHandler)

    //"Dynamic" static file server.
    fs := http.FileServer(http.Dir("web"))
    http.Handle("/", adaptFileServer(fs, mux))
}

func main() {
    log.Fatal(http.ListenAndServe(":5050", nil))
}
