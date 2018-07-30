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
	posts, err := posts.Get()
	if err != nil {
		encoder := json.NewEncoder(w)
		encoder.Encode(&err)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(&posts)
}

func Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	posts, err := posts.First(idInt)
	if err != nil {
		encoder := json.NewEncoder(w)
		encoder.Encode(&err)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(&posts)
}

func Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := make(map[string]string)
	data["title"] = r.FormValue("title")
	data["description"] = r.FormValue("description")
	err := posts.Insert(data)
	if err != nil {
		encoder := json.NewEncoder(w)
		encoder.Encode(&err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	data := make(map[string]string)
	data["title"] = r.FormValue("title")
	data["description"] = r.FormValue("description")
	err := posts.Update(data, idInt)
	if err != nil {
		encoder := json.NewEncoder(w)
		encoder.Encode(&err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	err := posts.Delete(idInt)
	if err != nil {
		encoder := json.NewEncoder(w)
		encoder.Encode(&err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
