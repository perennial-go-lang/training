package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rohitpavaskar/training/api/models/posts"
)

type Post struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// posts := []Post{
	// 	Post{1, "First Post", "Post Description"},
	// 	Post{2, "Second Post", "Post Description"},
	// 	Post{3, "Third Post", "Post Description"},
	// }
	posts := posts.Get()
	encoder := json.NewEncoder(w)
	encoder.Encode(&posts)
}

func Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	posts := posts.First(idInt)
	encoder := json.NewEncoder(w)
	encoder.Encode(&posts)
}

func Store(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["title"] = r.FormValue("title")
	data["description"] = r.FormValue("description")
	posts.Insert(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	data := make(map[string]string)
	data["title"] = r.FormValue("title")
	data["description"] = r.FormValue("description")
	posts.Update(data, idInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	posts.Delete(idInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
