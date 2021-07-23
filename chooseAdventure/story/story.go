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
func DecodeJSON() (map[string]Chapter, error) {
	content, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		return nil, err
	}
	var chapters map[string]Chapter
	err = json.Unmarshal(content, &chapters)
	return chapters, err
}

//StartGame sets up the template and server for the game
func StartGame(chapters map[string]Chapter) {
	tmpl := template.Must(template.ParseFiles("story.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})
	http.ListenAndServe(":80", nil)

}
