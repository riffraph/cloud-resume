package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	app.Initialize()
	result := m.Run()

	os.Exit(result)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestVisits(t *testing.T) {
	request := httptest.NewRequest("GET", "/visits", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	var statsObj SiteStatistics
	json.Unmarshal(response.Body.Bytes(), &statsObj)

	if statsObj.Visits < 0 {
		t.Errorf("Expect the visits to be a number greater than zero, instead got %v", statsObj.Visits)
	}
}

func TestAddVisit(t *testing.T) {
	request := httptest.NewRequest("GET", "/addvisit", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	var statsObj SiteStatistics
	json.Unmarshal(response.Body.Bytes(), &statsObj) // how do we check for errors here?

	// expect to get the updated count of visits in the response
	if statsObj.Visits < 0 {
		t.Errorf("Expect the visits to be a number greater than zero, instead got %v", statsObj.Visits)
	}
}

// test that add visit is really working, by comparing the site count before and after adding a visit
func TestGetSiteCountAndAddVisit(t *testing.T) {
	t.Error("not implemented")
}
