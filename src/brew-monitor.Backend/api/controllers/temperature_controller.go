package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	repository "github.com/gormaar/brew-monitor/api/repositories"
	response "github.com/gormaar/brew-monitor/api/responses"
)

func (server *Server) GetRecentTemperatureData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	brewId, err := strconv.ParseUint(vars["brew_id"], 10 ,32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tempModel := repository.Temperature{}
	temp,err := tempModel.GetRecentTemperature(server.DB, uint(brewId))
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, temp)
}

func (server *Server) GetAllTemperatureData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	brewId, err := strconv.ParseUint(vars["brew_id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tempModel := repository.Temperature{}
	temp, err := tempModel.GetAllTemperatures(server.DB, uint(brewId))
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, temp)
}

func (server *Server) CreateTemperature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	tempId, err := strconv.ParseUint(vars["brew_id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tempModel := repository.Temperature{}
	temp, err := tempModel.CreateTemperature(server.DB, r, uint(tempId))
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, temp)
}

func (server *Server) DeleteTemperature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	brewId, err := strconv.ParseUint(vars["brew_id"], 10, 32)
	tempId, err := strconv.ParseUint(vars["temp_id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tempModel := repository.Temperature{}
	temp, err := tempModel.DeleteTemperature(server.DB, uint(brewId), uint(tempId))
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, temp)
}

func (server *Server) PutTemperature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	brewId, err := strconv.ParseUint(vars["brew_id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tempModel := repository.Temperature{}
	temp, err := tempModel.PutTemperature(server.DB, uint(brewId))
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, temp)
}
