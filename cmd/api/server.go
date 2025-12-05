package main

import (
	"fmt"
	"net/http"
	"strings"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func teacherHandler(w http.ResponseWriter, r *http.Request) {

	//****Accessing the path parameters*****
	//teachers/{id}

	fmt.Println(r.URL.Path)
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	UserID := strings.TrimSuffix(path, "/")

	fmt.Println("UserID:", UserID)

	w.Write([]byte("hello teachers route"))
	fmt.Println("r.Method", r.Method)
	//creating CRUD opertaions
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET method called on Teacher route"))
	case http.MethodPost:
		w.Write([]byte("POST method called on Teacher route"))
	//***** processing the form request(xxx-form-urlencoded, multipart/form-data)*****
	// err := r.ParseForm()
	// if err != nil {
	// 	// sendinf the error response to client
	// 	http.Error(w, "Error parsing the form", http.StatusBadRequest)
	// 	fmt.Println("error parsing form: ", err)
	// 	return
	// }
	// fmt.Println(" form :", r.Form)
	// // Prepare Responsr data
	// response := make(map[string]interface{})
	// for key, value := range r.Form {
	// 	response[key] = value[0]
	// }
	// fmt.Println("response data: ", response)
	// // response data:  map[Name:[John Snow] age :[30] city:[LA]]
	// //Processing RAW body data(JSON)
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Error reading body", http.StatusBadRequest)
	// 	fmt.Println("error reading body: ", err)
	// 	return
	// }
	// defer r.Body.Close()
	//fmt.Println("raw body data: ", string(body))

	//Unmarshalling JSON body data to struct
	// var userInstance user
	// err = json.Unmarshal(body, &userInstance)
	// if err != nil {
	// 	http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
	// 	fmt.Println("error unmarshalling JSON: ", err)
	// 	return
	// }
	//fmt.Println("Unmarshalled user instance: ", userInstance)
	case http.MethodPatch:
		w.Write([]byte("PATCH method called on Teacher route"))
	case http.MethodDelete:
		w.Write([]byte("DELETE method called on Teacher route"))
	case http.MethodPut:
		w.Write([]byte("PUT method called on Teacher route"))

	}

}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	{
		w.Write([]byte("hello route"))
		fmt.Println("hello root router")
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("GET method called on students route"))
			fmt.Println(" form :", r.Form)
		case http.MethodPost:
			w.Write([]byte("POST method called on students route"))
		case http.MethodPatch:
			w.Write([]byte("PATCH method called on students route"))
		case http.MethodDelete:
			w.Write([]byte("DELETE method called on students route"))
		case http.MethodPut:
			w.Write([]byte("PUT method called on students route"))

		}
	}

}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello route"))
	fmt.Println("hello root router")
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET method called on exec route"))
		fmt.Println(" form :", r.Form)
	case http.MethodPost:
		w.Write([]byte("POST method called on exec route"))
	case http.MethodPatch:
		w.Write([]byte("PATCH method called on exec route"))
	case http.MethodDelete:
		w.Write([]byte("DELETE method called on exec route"))
	case http.MethodPut:
		w.Write([]byte("PUT method called on exec route"))

	}
}
func main() {

	//creating handlers and routes

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello route"))
	// 	fmt.Println("hello root router")

	// })

	http.HandleFunc("/teachers/", teacherHandler)

	http.HandleFunc("/students/", studentHandler)

	http.HandleFunc("/execs/", execsHandler)

	//creating server
	port := ":3000"
	fmt.Println("server is starting at port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("error starting server: ", err)
	}

}
