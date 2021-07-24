package story

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type option struct {
	Text string
	Arc  string
}

type Chapter struct {
	Title   string
	Story   []string
	Options []option
}

//DecodeJSON decodes the json file into a map of chapters
func DecodeJSON(filename string) (map[string]Chapter, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var chapters map[string]Chapter
	err = json.Unmarshal(content, &chapters)
	return chapters, err
}

//BuildTemplate sets up the http template
func buildTemplate(chapters map[string]Chapter, filename string) http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles(filename))
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if len(path) == 1 {
			path = "/intro"
		}

		chapter, ok := chapters[path[1:]]
		if ok {
			tmpl.Execute(w, chapter)
		} else {
			fmt.Fprintln(w, "Chapter not found!")
		}
	}
}

//StartGame sets up the template and server for the game
func StartGame(chapters map[string]Chapter) {
	http.HandleFunc("/", buildTemplate(chapters, "story.html"))
	http.ListenAndServe(":80", nil)

}
