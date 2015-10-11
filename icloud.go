package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/tgreiser/icloud/contacts"
	"github.com/tgreiser/icloud/engine"
)

var (
	appleId  = flag.String("apple_id", "", "Apple ID to log in")
	password = flag.String("password", "", "iCloud password")
)

func NewEngine(apple_id, password string) (*engine.ICloudEngine, error) {
	return engine.NewEngine(&http.Client{}, apple_id, password)
}

/*
func RemindersApp(e *ICloudEngine) (*ICloudRemindersApp, error){
	return reminders.NewApp(e)
}
*/

func main() {
	flag.Parse()
	if *appleId == "" || *password == "" {
		log.Fatal("Usage: go run icloud.go -apple_id=<apple ID> -password=<password>")
	}
	eng, e := NewEngine(*appleId, *password)
	if e == nil {
		fmt.Printf("Got engine: %v version: %v", eng, eng.Version)
		cr, _ := contacts.Get(eng)
		fmt.Printf("Got contacts: %v\n", cr.Contacts)
	} else {
		fmt.Printf("Got err: %v", e)
	}
}
