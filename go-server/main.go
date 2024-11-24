package main

import (
	"fmt"
	"log"
	"net/http"

)
// here are three routs 1. root rout (/) , 2. hello rout (/ hello) and third is form fout (/form)
func formHandler(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm();
	err != nil {
		fmt.Fprintf(w, "ParseFoen() err: %v" , err)
		return
	}
	fmt.Fprintf(w, "POST request sucessful")
	name:= r.FormValue("name")
	adress:= r.FormValue("adress")
	fmt.Fprintf(w, "name = %s\n" , name)
	fmt.Fprintf(w, "adress = %s\n" , adress)

}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found" , http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method is not found" , http.StatusNotFound)
		return
		
	}
	fmt.Fprintf(w , "hello!") 
}
func main() {
  fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/" , fileServer)// ye sirf /handle hai tabhi  HandleFuck nhi use hua hai
	http.HandleFunc("/form" , formHandler)
	http.HandleFunc("/hello" , helloHandler)

	fmt.Printf("starting server at port:8080\n")
	if err := http.ListenAndServe(":8080" , nil);// hart of softwere
	 err != nil {
		log.Fatal(err)
	}
}