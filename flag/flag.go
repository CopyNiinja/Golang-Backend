// The flag package in Go is used for command-line arguments parsing. It helps you read values like integers, strings, booleans, etc. from the command line when running your program.
// it is actually pointer. it holds the address
// you must parse it for cmd
// If you skip arguments, defaults are used
package main

import (
	"flag"
	"fmt"
)

func main() {

	name:= flag.String("name","guest","your name");
    
	// flag.Parse();   //flag.Parse() can only be called once.

	//fmt.Println(*name); // faiyaz   go run main.go -name=faiyaz
	//fmt.Println(*name); // guest   go run main.go (without -name or --name flag)


	//Instead of pointer return, you can bind directly to a variable:

	var port int;

	flag.IntVar(&port,"port",8080,"port number");

	flag.Parse();
    fmt.Println(*name)
	fmt.Printf("Port:%d",port); //if no argument: port:8080 
	
    //go run flag.go --name=Faiyaz -port=1234

	//Faiyaz
    //Port:1234


}