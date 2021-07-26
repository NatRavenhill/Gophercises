package urlshort

import (
	"encoding/json"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		path, ok := pathsToUrls[r.RequestURI]
		if ok {
			http.Redirect(rw, r, path, http.StatusFound)
		} else if fallback != nil {
			fallback.ServeHTTP(rw, r)
		}
	}
}

type Link struct {
	Path string
	Url  string
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(content []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var links []Link

	err := yaml.Unmarshal(content, &links)
	if err != nil {
		return nil, err
	}

	return MapHandler(createPaths(links), fallback), nil
}

//JSONHandler parses the provided JSON file and then returns
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
func JSONHandler(content []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var links []Link
	err := json.Unmarshal(content, &links)
	if err != nil {
		return nil, err
	}

	return MapHandler(createPaths(links), fallback), nil
}

//Creates the map from a link array
func createPaths(links []Link) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, val := range links {
		pathsToUrls[val.Path] = val.Url
	}

	return pathsToUrls
}
