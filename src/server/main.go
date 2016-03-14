package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("../")))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/all", allHandler)

	certPath := "/etc/letsencrypt/live/repdom.umassmed.io/cert.pem"
	privKeyPath := "/etc/letsencrypt/live/repdom.umassmed.io/privkey.pem"

	// loadPages()

	// open a simple log file
	logFile, err := os.OpenFile("myLog.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(time.Now(), err)
	}
	defer logFile.Close()

	// set up the writer
	logWriter := bufio.NewWriter(logFile)
	defer logWriter.Flush()

	//
	err = http.ListenAndServeTLS(":443", certPath, privKeyPath, nil)
	if err != nil {
		fmt.Println(time.Now().String() + err.Error())
	}
}

// AllHandler returns all the data
func allHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
