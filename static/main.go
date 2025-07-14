package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func indexfunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html");
	// Read and serve the index.html file
	http.ServeFile(w, r, "./static/index.html")
}


func hellofunc(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, Go is working!"})
}

func formfunc(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","text/html")
  http.ServeFile(w,r,"./static/form.html")
}

func main(){
	
	http.HandleFunc("/",indexfunc)
	http.HandleFunc("/hello",hellofunc)
	http.HandleFunc("/form",formfunc)

	fmt.Println("Server running at http://localhost:3000")
	
	http.ListenAndServe(":3000",nil)
}