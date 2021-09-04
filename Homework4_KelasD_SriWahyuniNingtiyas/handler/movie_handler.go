package handler

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/rysmaadit/go-template/common/responder"
    "github.com/rysmaadit/go-template/model"
    "gorm.io/gorm"
    "net/http"
)

func CreateMovie(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        var movie model.Movie
        err := json.NewDecoder(r.Body).Decode(&movie)
        if err != nil {
            responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
            return
        }
        db.Create(&movie)
        responder.NewHttpResponse(r, w, http.StatusCreated, &movie, nil)
    }
}

func GetMovie(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        var movie model.Movie

        params := mux.Vars(r)
        slug := params["slug"]
        db.Where(&model.Movie{Slug: slug}).First(&movie)

        responder.NewHttpResponse(r, w, http.StatusOK, &movie, nil)
        err := json.NewDecoder(r.Body).Decode(&movie)
        if err != nil {
            responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
            return
        }
    }
}

func GetMovies(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var movie []model.Movie
        db.Find(&movie)
        responder.NewHttpResponse(r, w, http.StatusOK, &movie, nil)

        err := json.NewDecoder(r.Body).Decode(&movie)
        if err != nil {
            responder.NewHttpResponse(r, w, http.StatusNotFound, nil, err)
            return
        }
    }
}

func DeleteMovie(db *gorm.DB) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request){

        var movie model.Movie
        params := mux.Vars(r)
        slug := params["slug"]
        db.Where(&model.Movie{Slug: slug}).First(&movie)
        db.Delete(&movie)

        responder.NewHttpResponse(r,w, http.StatusOK, &movie,nil)

        err := json.NewDecoder(r.Body).Decode(&movie)
        if err != nil {
            responder.NewHttpResponse(r, w, http.StatusNotFound, nil, err)
            return
        }
    }
}

func UpdateMovie(db *gorm.DB) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request){
        var movie model.Movie
        err := json.NewDecoder(r.Body).Decode(&movie)
        if err!= nil{
            responder.NewHttpResponse(r,w,http.StatusBadRequest,nil,err)
            return
        }

        params := mux.Vars(r)
        slug := params["slug"]

        db.Where(&model.Movie{Slug: slug}).Save(&movie)

        responder.NewHttpResponse(r,w, http.StatusOK, &movie,nil)
    }
}


