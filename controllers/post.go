package controllers

import (
	"blog-platform/models"
	"blog-platform/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	resp := post.Create()
	utils.Respond(w, map[string]interface{}{"message": resp})
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := models.GetPosts()
	response := make(map[string]interface{})
	response["posts"] = posts
	utils.Respond(w, response)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	resp := models.DeletePost(postID)
	utils.Respond(w, map[string]interface{}{"message": resp})
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post := models.GetPostByID(postID)
	if post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	response := make(map[string]interface{})
	response["post"] = post
	utils.Respond(w, response)
}
