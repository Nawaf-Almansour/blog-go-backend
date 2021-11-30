package main

import (
	"backend/models"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
	}
	app.logger.Println("id is ", id)

	movie := models.Movie{
		ID: id,
		Title: "some title",
		Description: "Some description",
		Year: 2010,
		ReleaseDate: time.Date(2010,01,01,01,0,0,0,time.Local),
		Runtime: 100,
		Rating: 5,
		MPAARating: "PG-13",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
err = app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

	app.logger.Println("all is ")
}