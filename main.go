package main

import (
	"auth-server/api"
	"auth-server/auth"
	"auth-server/config"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var c = loadConfig()
var wg sync.WaitGroup
var StopFlag = new(bool)

var testUsername = "michael"
var testPassword = "bobbitt"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds | log.Ldate)
	log.Println("Starting")

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	*StopFlag = false

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	testFuncs(StopFlag)
	//}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		api.Start(StopFlag)
	}()

	<-stopChan
	*StopFlag = true

	log.Println("Waiting for threads to finish...")
	wg.Wait()
	log.Println("Stopped")
	os.Exit(0)
}

func testFuncs(flag *bool) {
	err := auth.Init(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Initialized")

	err = auth.CreateUser(c, "testUsername", testPassword)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User created", testUsername, testPassword)

	user, err := auth.GetUser(c, testUsername)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User found", user)

	err = auth.DeleteUser(c, user.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User deleted", user.Id)

	user, err = auth.GetUser(c, testUsername)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("User not found")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("User found", user)
	}

	for !*flag {
		log.Println("Waiting")
		time.Sleep(5 * time.Second)
	}
}

func loadConfig() config.Config {
	return config.Config{
		DatabasePath: "./test.db",
	}
}
