package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	http.Handle("/form", http.HandlerFunc(acceptForm))
	http.Handle("/form2", http.HandlerFunc(acceptForm2))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

func acceptForm(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n\n=========================================\n\n")
	fmt.Printf("Method: %s URL: %s\n", r.Method, r.URL.String())

	headerContentTtype := r.Header.Get("Content-Type")
	fmt.Printf("form :: Request headers:\n")
	for k, v := range r.Header {
		fmt.Printf("\t%s: %#v\n", k, v)
	}
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	r.ParseForm()

	fmt.Printf("\nrequestForm:\n")
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	for key, value := range r.Form {
		fmt.Printf("\t%s: %#v\n", key, value)
		m1[key] = value
	}
	fmt.Println("\nrequestPostForm:")
	for key, value := range r.PostForm {
		fmt.Printf("\t%s: %#v\n", key, value)
		m2[key] = value
	}
	m3 := make(map[string]interface{})
	m3["requestForm"] = m1
	m3["requestPostForm"] = m2
	b, err := json.Marshal(m3)
	if err != nil {
		fmt.Printf("\nError while json marshal: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	return
}

func acceptForm2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n\n=========================================\n\nform2 :: Request headers:\n")
	for k, v := range r.Header {
		fmt.Printf("%s: %#v\n", k, v)
	}
	r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	m1 := make(map[string]interface{})

	fmt.Println("\nrequestMultipartFormValue:")
	for key, value := range r.MultipartForm.Value {
		fmt.Printf("%s: %#v\n", key, value)
		m1[key] = value
	}
	m3 := make(map[string]interface{})
	m3["requestForm"] = m1
	m3["requestPostForm"] = map[string]interface{}{}
	b, err := json.Marshal(m3)
	if err != nil {
		fmt.Printf("\nError while json marshal: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	return
}
