package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type StatsResponse struct {
	Visits int `json:"Visits"`
}

func TestStatsHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "/stats", nil)
	responseRecorder := httptest.NewRecorder()
	statsHandler(responseRecorder, request)

	response := responseRecorder.Result()
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	var statsResponse StatsResponse
	json.Unmarshal(data, &statsResponse) // how do we check for errors here?

	if statsResponse.Visits < 0 {
		t.Errorf("Expect the visits to be a number greater than zero, instead got %v", statsResponse.Visits)
	}
}

func TestAddVisitHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "/addvisit", nil)
	responseRecorder := httptest.NewRecorder()
	statsHandler(responseRecorder, request)

	response := responseRecorder.Result()
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	// expect to get the updated count of visits in the response
	var statsResponse StatsResponse
	json.Unmarshal(data, &statsResponse) // how do we check for errors here?

	if statsResponse.Visits < 0 {
		t.Errorf("Expect the visits to be a number greater than zero, instead got %v", statsResponse.Visits)
	}
}

// test that add visit is really working, by comparing the site count before and after adding a visit
func TestGetSiteCountAndAddVisit(t *testing.T) {
	t.Error("not implemented")
}
