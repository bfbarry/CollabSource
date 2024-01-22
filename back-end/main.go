package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"log"
	"github.com/bfbarry/CollabSource/back-end/controllers"
	"github.com/bfbarry/CollabSource/back-end/connections"
)


func main() {
	const port_num = 3333
	log.SetOutput(os.Stdout)	
	client, db := connections.InitDB()
	defer connections.CloseDB(client)()
	mux := http.NewServeMux()
	env := &controllers.Env{DB: db}
	mux.HandleFunc("/get_projects", env.GetSampleProjects)
	log.Printf("listening on :%d\n", port_num)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port_num), mux)
	log.Println("closed")
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("http.ErrServerClosed: server shut down \n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}