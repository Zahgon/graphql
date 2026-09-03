package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/graphql-go/graphql"
)

var schema graphql.Schema

const jsonDataFile = "data.json"

func handleSIGUSR1(c chan os.Signal) { _ = "STUB: not implemented"; return }

func filterUser(data []map[string]interface{}, args map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	_ = "STUB: not implemented"
	return nil
}

func importJSONDataFromFile(fileName string) error { _ = "STUB: not implemented"; return nil }

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go handleSIGUSR1(c)

	err := importJSONDataFromFile(jsonDataFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(name:\"Dan\"){id,surname}}'")
	fmt.Printf("Reload json file   : kill -SIGUSR1 %s\n", strconv.Itoa(os.Getpid()))
	http.ListenAndServe(":8080", nil)
}
