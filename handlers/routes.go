package handlers

import (
	"api-booking/models"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewHandler(db *sql.DB) http.Handler {
	router := chi.NewRouter()
	router.Get("/", Router(db))
	router.Post("/post", User(db))
	return router
}

func User(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		var user models.User
		user.Username = "phong"
		user.Password = "admin"
		var err = user.Insert(ctx, db, boil.Infer())
		response, _ := json.Marshal("ok")
		if err != nil {
			response, _ = json.Marshal("no ok")
		}
		w.Write(response)
	}
}

func Router(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		account, _ := models.Users().All(ctx, db)
		for _, data := range account {
			fmt.Println("data", data.Email1.String)
		}
		response, _ := json.Marshal(account)
		w.Write(response)
	}
}
