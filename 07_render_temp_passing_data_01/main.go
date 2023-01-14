// creating a new html file store temp value in that file

package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	var myName = "Vikash Parashar"
	// parsing template file index.html
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	// executing template that we already parsed above
	if err = temp.Execute(os.Stdout, myName); err != nil {
		log.Fatalln("failed to execute templates")
	}

}
