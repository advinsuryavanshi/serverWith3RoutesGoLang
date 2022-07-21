package main

import  "fmt"
import	"log"
import  "net/http"
func formHandle(w http.ResponseWriter, r *http.Request){
	if  err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "PraseForm() err: %v", err)
		return 
	}
	fmt.Fprintf(w,"POST request succesful\n")
	name:= r.FormValue("name")
	address:= r.FormValue("address")
	fmt.Fprintf(w, "name = %s\n" ,Name)
	fmt.Fprintf(w,"address = %s\n" ,address)
}


func helloHandler(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/hello" {
		http.Error(w, "ERROR 404 NOT FOUND", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{

		http.Error(w, "Methord Not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}


func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandle)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("statring the server 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}