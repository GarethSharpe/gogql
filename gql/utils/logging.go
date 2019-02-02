package utils

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"time"
	"bytes"
)

const (
	INFO    string 	= " \033[1;34m%s\033[0m\n"
	NOTICE  string 	= " \033[1;36m%s\033[0m\n"
	WARNING string	= " \033[1;33m%s\033[0m\n"
	ERROR   string	= " \033[1;31m%s\033[0m\n"
	DEBUG   string 	= " \033[0;36m%s\033[0m\n"
)

type logEntry struct {
	OperationName 	string
	Variables		string
	//Query			string
	RemoteAddr		string
	Host			string
	Proto			string
	ResponseTime	string
}

type graphQLRequest struct {
	OperationName 	string
	Variables		string
	//Query			string
}

func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the content
		var bodyBytes []byte
		var log logEntry
		if r.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(r.Body)
			bodyString := string(bodyBytes)
			var gql graphQLRequest
			json.Unmarshal([]byte(bodyString), &gql)
			log = logEntry{
				OperationName: fmt.Sprintf(INFO, gql.OperationName),
				Variables: fmt.Sprintf(INFO, gql.Variables),
				// Query: fmt.Sprintf(INFO, gql.Query),
				RemoteAddr: fmt.Sprintf(INFO, r.RemoteAddr),
				Host: fmt.Sprintf(INFO, r.Host),
				Proto: fmt.Sprintf(INFO, r.Proto),
			}
		}

		// Restore the io.ReadCloser to its original state
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		t1 := time.Now()
		next.ServeHTTP(w, r)
		log.ResponseTime = fmt.Sprintf(WARNING, time.Since(t1).String())
		fmt.Printf(ERROR, "***********************")
		fmt.Printf(DEBUG, "-----------------------")
		fmt.Printf(DEBUG, "Log")
		fmt.Printf(DEBUG, "-----------------------")
		fmt.Printf("%+v\n", log)
		fmt.Printf(ERROR, "***********************")
	}
}