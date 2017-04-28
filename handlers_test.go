//filename=iam/handlers_test.go
package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TestDude
func TestDude (t *testing.T) {
	log.Printf("[TestDude] - Starting test.\n")
	        req := httptest.NewRequest("GET","http://sc.ipcloud.com",nil)

       //create a ResponseRecorder to record the response
        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(Dude)
        //handler := rootH()
        //
    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
    // Check the status code is what we expect. -> StatusTemporaryRedirect (307)
    //if status := rr.Code; status != http.StatusTemporaryRedirect {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusTemporaryRedirect)
	}
        t.Logf("\t[%v] - dude() tested. Got expected %d = %d.\n","checkMark",http.StatusTemporaryRedirect,rr.Code)

}
