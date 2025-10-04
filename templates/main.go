package main

import (
	"html/template"
	"net/http"
)

type Content struct {
	Title   string
	Heading string
	Message string
}
// //instead of parsing of each request ,parse first then wait for req
// var tmpl *template.Template;
// //runs before main
// func init(){
//  tmpl=template.Must(template.ParseFiles("index.gohtml"));
// }
func main() {
	// html/template package  ( Parse and execute)
    
	// home url
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		//Parse
        if tmpl,err:= template.ParseFiles("index.gohtml");err!=nil{
            //error
			http.Error(w,"template parse error: "+err.Error(),http.StatusInternalServerError)
			return
		}else{
            data:=Content{
				Title: "This is a title",
				Heading: "This is a heading",
				Message: "This is a message",
			}
			//Execute
			if err:=tmpl.Execute(w,data);err!=nil{
				http.Error(w,"template execution error:"+err.Error(),http.StatusInternalServerError)
				return
			}

		}
		
	})
    
	//server port and listening
	http.ListenAndServe(":80",nil);

}