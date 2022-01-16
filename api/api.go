package main
import (
    "encoding/json"
    "log"
    "net/http"
)
type director struct {
    Name string `json:"name"`
}
type movie struct {
    Universe    string   `json:"universe"`
    Title       string   `json:"title"`
    Year 	 	int      `json:"year"`
}

var movies = [...]movie{
    {"Marvel", "Iron Man", 2008},
    {"DC", "The Dark Knight", 2008},
    {"Marvel", "The Avengers", 2012},
    {"DC", "Man of Steel", 2013},
    {"Marvel", "Black Panther", 2018},
    {"DC", "Justice League", 2017},
}
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        j, _ := json.Marshal(movies)
        w.Header().Set("Content-Type", "application/json")
        w.Write(j)
    })
    log.Fatal(http.ListenAndServe(":80", nil))
}