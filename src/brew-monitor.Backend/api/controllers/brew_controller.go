package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	repository "github.com/gormaar/brew-monitor/api/repositories"
	response "github.com/gormaar/brew-monitor/api/responses"
)

func (server *Server) GetBrew(w http.ResponseWriter, r *http.Request) {
	setCors(&w)
	vars := mux.Vars(r)

	brewId, err := strconv.ParseUint(vars["brew_id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	brewModel := repository.Brew{}
	brew, err := brewModel.GetBrew(server.DB, uint(brewId))
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, brew)
}

func (server *Server) GetBrews(w http.ResponseWriter, r *http.Request) {
	setCors(&w)
	brewModel := repository.Brew{}
	brews, err := brewModel.GetBrews(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, brews)
}

func (server *Server) CreateBrew(w http.ResponseWriter, r *http.Request) {
	setCors(&w)
	brewModel := repository.Brew{}
	brew, err := brewModel.CreateBrew(server.DB, r)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, brew)
}

func (server *Server) DeleteBrew(w http.ResponseWriter, r *http.Request) {
	setCors(&w)
	vars := mux.Vars(r)
	brewId, err := strconv.ParseUint(vars["brew_id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	brewModel := repository.Brew{}
	brew, err := brewModel.DeleteBrew(server.DB, uint(brewId))
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, brew)
}

func (server *Server) PutBrew(w http.ResponseWriter, r *http.Request) {
	setCors(&w)
	var b repository.Brew
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	brew, err := b.PutBrew(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, brew)
}