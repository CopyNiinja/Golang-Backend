package main

import (
	"fmt"
	"net/http"
)

func main(){
     //separate goroutine for each user!
	handler:=func(w http.ResponseWriter ,r *http.Request){ 
		//An http.ResponseWriter which is where you write your text/html response to.
		//An http.Request which contains all information about this HTTP request including things like the URL or header fields.
		fmt.Fprintf(w,"Hello world!")
	}
     //  for home url
	http.HandleFunc("/",handler);

	//port 
	http.ListenAndServe(":80",nil);



}