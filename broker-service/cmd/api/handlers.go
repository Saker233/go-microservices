package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type requestAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestBrokerPayload struct {
	Action string      `json:"action"`
	Auth   requestAuth `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *Config) handleSubmission(w http.ResponseWriter, r *http.Request) {
	var req RequestBrokerPayload

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch req.Action {
	case "auth":
		app.authenticate(w, req.Auth)
	case "log":
		app.logItem(w, req.Log)
	default:
		app.errorJSON(w, errors.New("unknown Action"))
	}

}
func (app *Config) authenticate(w http.ResponseWriter, a requestAuth) {
	// create some json which will be sent to auth service
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service:8000/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
	}
	// request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back to the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we will read response body
	var jsonFromService jsonResponse

	// decode the json from the service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	// now we have a valid login
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated!"
	payload.Data = jsonFromService

	app.writeJSON(w, http.StatusAccepted, payload)
}


func (app *Config) logItem(w http.ResponseWriter, entry LogPayload) {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	
	logServiceURL := "http://logger-service:8000/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	respose, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer respose.Body.Close()

	if respose.StatusCode != http.StatusAccepted {
		app.errorJSON(w, err)
		return
	}

	var resp jsonResponse
	resp.Error = false
	resp.Message = "logged"

	app.writeJSON(w, http.StatusAccepted, resp)
}