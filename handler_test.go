package go_web_basic

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	// bikin handler func -> handlerFunc tapi cuman bisa buat 1 endpoint aja
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		_, _ = fmt.Fprint(writer, "Hello Deby")
	}

	// bikin server & set handler
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "Hi Deby")
	})

	// priority url in serve mux -> tapi klo bisa bikin url unique aja biar gk bingung
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "Images")
	})

	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "Images Thumbnails")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, request.Method)
		_, _ = fmt.Fprint(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
