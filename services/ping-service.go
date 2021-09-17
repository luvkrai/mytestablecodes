package services

import (
	"fmt"
)

const (
	pong = "pong"
)


// Create a public gobal variable of type "pingService" (interface) and assign it to the structure which implements this interface i.e "pingServiceImplementation"
var (
	PingService pingService = &pingServiceImplementation{}
)

// Any structure which has HandlePing() (string, error) implemented will be of this "pingService" kind interface
type pingService interface {
	HandlePing() (string, error)
}


// Structure for our service (i.e this must implicitly implement pingService interface)
type pingServiceImplementation struct {

}


func ( pingServer *pingServiceImplementation) HandlePing() (string, error) {    // Now structure "pingServiceImplementation" implements "pingService" interface 
	fmt.Println("doing some complex things...")
	return pong, nil
}