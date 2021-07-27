package server

import (
	"io/ioutil"
	"net/http"

	"github.com/ash3798/url-shortner/task"
)

//HandleRequest is Request handler function for server
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read the body of the request sent", http.StatusBadRequest)
		return
	}

	//Get request for the URL
	if r.Method == "GET" {
		task.GetShortURL("testURL")
	}

	//create URL Request
	if r.Method == "POST" {
		err = task.CreateShortURL(data)
		if err != nil {
			http.Error(w, "Error while processing the request."+err.Error(), http.StatusBadRequest)
			return
		}
	}

	//invalid method
	http.Error(w, "Wrong http method used. Please use GET for existing URL and POST for creating new URL", http.StatusMethodNotAllowed)
}
