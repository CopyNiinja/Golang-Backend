package main

import "fmt"

func main() {
	name := "Faiyaz"

	template := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    <h1>Hi ` + name + ` </h1>
  </body>
</html>
`
	fmt.Println(template)

	//run code with this line: go run templates/html-templates.go > index.html
	//  > → redirects the program’s standard output (stdout) to a file.
  
}
