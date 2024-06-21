package services

import (
	"database/sql"
	"time"
)

var db*sql.DB

const dbTimeout = 5 * time.Second

type Models struct {
	Users User
	Albums Album
	Posts Post
	JsonResponse JsonResponseModel
}

func New(dbPool * sql.DB) Models{
	db = dbPool
	return Models{}
}