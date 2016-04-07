package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

/*
vars := mux.Vars(req)
		matchID := vars["id"]
		match, err := repo.getMatch(matchID)*/
func lastAlertHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, "tbd")
	}
}

func lastTelemetryHandler(formatter *render.Render, repo eventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		droneID := getDroneID(req)
		fmt.Printf("Looking up last telemetry event for drone %s", droneID)
		event, err := repo.GetTelemetryEvent(droneID)
		if err == nil {
			formatter.JSON(w, http.StatusOK, &event)
		} else {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func lastPositionHandler(formatter *render.Render, repo eventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		droneID := getDroneID(req)
		fmt.Printf("Looking up last position event for drone %s", droneID)
		event, err := repo.GetPositionEvent(droneID)
		if err == nil {
			formatter.JSON(w, http.StatusOK, &event)
		} else {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func getDroneID(req *http.Request) (droneID string) {
	vars := mux.Vars(req)
	droneID = vars["droneId"]
	return
}