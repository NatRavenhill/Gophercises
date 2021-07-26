//package main creates an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"Gophercises/urlShortener/urlshort"
)

func main() {
	mode := flag.String("mode", "yaml", "Sets the file parsing mode, accept yaml or json")

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	// yaml := `- path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution

	var handler http.HandlerFunc
	switch *mode {
	case "yaml":
		content, err := ioutil.ReadFile("paths.yaml")
		if err != nil {
			log.Fatal(err)
		}
		handler, err = urlshort.YAMLHandler(content, mapHandler)
		if err != nil {
			log.Fatal(err)
		}
	case "json":
		content, err := ioutil.ReadFile("paths.json")
		if err != nil {
			log.Fatal(err)
		}
		handler, err = urlshort.JSONHandler(content, mapHandler)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Mode flag not recognised")
	}

	http.ListenAndServe(":8080", handler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
