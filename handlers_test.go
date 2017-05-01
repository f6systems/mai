//filename=iam/handlers_test.go
package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	log.Printf("[handers_test:init()] <INFO> Simple log message to note Tests starting.(TODO:(hopley) - An internet resource to capture/log event.)\n")
}

//TestDude
func TestDude(t *testing.T) {
	log.Printf("[TestDude] - Starting test.\n")
	req := httptest.NewRequest("GET", "http://sc.ipcloud.com", nil)

	//create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Dude)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusTemporaryRedirect)
	}
	t.Logf("\t[%v] - Dude() tested. Got expected %d = %d.\n", "<OKAY>", http.StatusTemporaryRedirect, rr.Code)

}

func TestIP(t *testing.T) {
	log.Printf("[TestDude] - Starting test.\n")
	req := httptest.NewRequest("GET", "http://sc.ipcloud.com", nil)

	//create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IP)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusTemporaryRedirect)
	}
	t.Logf("\t[%v] - IP() tested. Got expected %d = %d.\n", "<OKAY>", http.StatusTemporaryRedirect, rr.Code)

}
