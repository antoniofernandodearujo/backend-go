package controllers

import (
	"challenge-api/internal/helpers"
	"challenge-api/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var album services.Album

// GetAllAlbums godoc
// @Summary Get all albums
// @Description Retrieves a list of all album entries
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {object} services.AlbumsList
// @Router /albums [get]
func GetAllAlbums( w http.ResponseWriter, r *http.Request) {
	albums, err := album.GetAllAlbums()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting all albums: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"albums": albums}, nil)
}

// GetAlbumByID godoc
// @Summary Get album by ID
// @Description Retrieves a specific album by its ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} services.Album
// @Router /albums/{id} [get]
func GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	album, err := album.GetAlbumByID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting album by id: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"album": album}, nil)
}

// CreateAlbum godoc
// @Summary Create an album
// @Description Creates a new album entry
// @Tags albums
// @Accept json
// @Produce json
// @Param albumData body services.AlbumPayload true "Album Data"
// @Success 201 {object} services.Album
// @Router /albums [post]
func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var albumData services.Album
	err := json.NewDecoder(r.Body).Decode(&albumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error parsing album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	albumCreated, err := album.CreateAlbum(albumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error creating album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	helpers.WriteJSON(w, http.StatusCreated, helpers.Envelop{"album": albumCreated}, nil)
}

// UpdateAlbum godoc
// @Summary Update an album
// @Description Updates an existing album identified by ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Param albumData body services.AlbumPayload true "Album Data"
// @Success 200 {object} services.Album
// @Router /albums/{id} [put]
func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var albumData services.Album
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&albumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error parsing album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	albumUpdated, err := album.UpdateAlbum(id, albumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error updating album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"album": albumUpdated}, nil)
}

// DeleteAlbum godoc
// @Summary Delete an album
// @Description Deletes an album entry identified by ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} Message
// @Router /albums/{id} [delete]
func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := album.DeleteAlbum(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error deleting album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"message": "Album deleted successfully"}, nil)
}

// GetAlbumsByUserID godoc
// @Summary Get albums by user ID
// @Description Retrieves a list of albums associated with a user ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} services.AlbumsList
// @Router /users/{id}/albums [get]
func GetAlbumsByUserID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	albums, err := album.GetUserAlbums(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting user albums: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"albums": albums}, nil)
}

// AddAlbumToUser godoc
// @Summary Add album to user
// @Description Associates an album with a user
// @Tags albums
// @Accept json
// @Produce json
// @Param userAlbumData body services.UserAlbumPayload true "User Album Data"
// @Success 201 {object} services.UserAlbum
// @Router /albums/save [post]
func AddAlbumToUser(w http.ResponseWriter, r *http.Request) {
	var userAlbumData services.UserAlbum
	err := json.NewDecoder(r.Body).Decode(&userAlbumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error decoding user album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	albumAdded, err := album.AddAlbumToUser(userAlbumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error adding album to user: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	helpers.WriteJSON(w, http.StatusCreated, helpers.Envelop{"album": albumAdded}, nil)
}

// RemoveAlbumFromUser godoc
// @Summary Remove album from user
// @Description Removes an album association from a user
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param album_id path string true "Album ID"
// @Success 200 {object} Message
// @Router /users/{id}/albums/{album_id} [delete]
func RemoveAlbumFromUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	albumID:= chi.URLParam(r, "album_id")
	err := album.RemoveAlbumFromUser(userID, albumID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error removing album from user: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"message": "Album removed from user successfully"}, nil)
}
