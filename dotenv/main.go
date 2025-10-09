/*
🧩 1. Declaring (Setting) Environment Variables

Environment variables are not declared inside Go code —
they are declared in your system (terminal or OS).

🔹 On Linux or macOS:
export DATABASE_URL=postgres://postgres:1234@localhost:5432/movies

🔹 On Windows (PowerShell):
$env:DATABASE_URL = "postgres://postgres:1234@localhost:5432/movies"

This variable now exists in your environment while the terminal is open.
When you run your Go program in the same terminal, Go can access it.

----------------------------------------------------------------------
to use a package to read dotenv file:

	package: 	  go get github.com/joho/godotenv
*/
package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	 
	//load dotenv

	err:= godotenv.Load("./dotenv/.env");

	//handling error
	if err!=nil{
		fmt.Println(err);
	}
	//get the environment variable
	port:= os.Getenv("PORT")

	fmt.Println(port)
}