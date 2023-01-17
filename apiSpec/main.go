package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	// Read the OpenAPI specification file into memory
	data, err := ioutil.ReadFile("openapi.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	// Create a new loader
	loader := openapi3.NewLoader()

	// Parse the specification file
	document, err := loader.LoadFromData(data)
	if err != nil {
		log.Fatalln(err)
	}

	//validate the specification using the loader context
	if err := document.Validate(loader.Context); err != nil {
		log.Fatalln(err)
	}

	// Create a new template by parsing a file
	tmpl, err := template.ParseFiles("server.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	for _, pathItem := range document.Paths {
		fmt.Println("========================================\n")
		fmt.Printf("%v\n", pathItem)
		fmt.Println("%v\n", pathItem.ExtensionProps)
		fmt.Println("========================================\n")
		fmt.Printf("%v\n", string(pathItem.ExtensionProps.Extensions["x-arg-1"].(json.RawMessage)))
		// 		v := reflect.ValueOf(*pathItem)
		// 		values := make([]interface{}, v.NumField())

		// 		for i := 0; i < v.NumField(); i++ {
		// 			values[i] = v.Field(i).Interface()
		// 		}
		// fmt.Println(values)
		// fmt.Println(values[0])
		// fmt.Printf("%T\n", values[0])
	}
	// x := struct{Foo string; Bar int }{"foo", 2}

	// v := reflect.ValueOf(x)

	// values := make([]interface{}, v.NumField())

	// for i := 0; i < v.NumField(); i++ {
	//     values[i] = v.Field(i).Interface()
	// }

	// fmt.Println(values)
	// Open or Create the output file
	file, err := os.OpenFile("server.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, document); err != nil {
		log.Fatalln(err)
	}

}
