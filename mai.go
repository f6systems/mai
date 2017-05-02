//filename=mai.go
//see=hopley@ipcloud.net
package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	log.Printf("init()\n")
}

func main() {
	log.Printf("Starting. (Testing drone v0.6.0-rc1)\n")
	srvPort := "8888"
	http.Handle("/", http.HandlerFunc(sayHello))
	http.Handle("/ip", http.HandlerFunc(IP))
	http.Handle("/dude", http.HandlerFunc(Dude))
	log.Printf("Starting. Server mai to start on port = %q\n",srvPort)
	err := http.ListenAndServe(srvPort, nil)
	if err != nil {
		fmt.Printf("ListenAndServe Error. (err=%s)\n", err)
	}

}
