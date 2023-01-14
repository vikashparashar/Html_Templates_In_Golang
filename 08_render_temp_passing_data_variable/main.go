// creating a new html file store temp value in that file

package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	Channel   string
}

var u1 User = User{
	"vikash",
	"parashar",
	29,
	"#Code",
}

func main() {
	// parsing template file index.html
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	// executing template that we already parsed above
	if err = temp.Execute(os.Stdout, u1); err != nil {
		log.Fatalln("failed to execute templates")
	}

}
