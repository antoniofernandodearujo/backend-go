package controllers

import (
	"challenge-api/internal/helpers"
	"challenge-api/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var post services.Post

// GetAllPosts godoc
// @Summary Get all posts
// @Description Retrieves a list of all posts
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {object} services.PostsList
// @Router /posts [get]
func GetAllPosts(w http.ResponseWriter, r *http.Request)  {
	posts, err := post.GetAllPosts()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting all posts: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"posts": posts}, nil)
}

// GetPostByID godoc
// @Summary Get post by ID
// @Description Retrieves a specific post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} services.Post
// @Router /posts/{id} [get]
func GetPostByID(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")
	post, err := post.GetPostByID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting post by id: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"post": post}, nil)
}

// CreatePost godoc
// @Summary Create a post
// @Description Creates a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param postData body services.PostPayload true "Post Data"
// @Success 201 {object} services.Post
// @Router /posts/create [post]
func CreatePost(w http.ResponseWriter, r *http.Request)  {
	var postData services.Post
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error parsing post: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postCreated, err := post.CreatePost(postData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error creating post: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusCreated, helpers.Envelop{"post": postCreated}, nil)
}

// UpdatePost godoc
// @Summary Update a post
// @Description Updates an existing post identified by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param postData body services.PostPayload true "Post Data"
// @Success 200 {object} services.Post
// @Router /posts/{id} [put]
func UpdatePost(w http.ResponseWriter, r *http.Request)  {
	var postData services.Post
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error parsing post: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postUpdated, err := post.UpdatePost(id, postData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error updating post: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"post": postUpdated}, nil)
}

// DeletePost godoc
// @Summary Delete a post
// @Description Deletes a post entry identified by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} Message
// @Router /posts/{id} [delete]
func DeletePost(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")
	err := post.DeletePost(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error deleting post: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"post": "deleted"}, nil)
}

// GetPostsByUserID godoc
// @Summary Get posts by user ID
// @Description Retrieves a list of posts associated with a user ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} services.PostsList
// @Router /users/{id}/posts [get]
func GetPostsByUserID(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")
	posts, err := post.GetPostsByUserID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting posts by user id: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"posts": posts}, nil)
}