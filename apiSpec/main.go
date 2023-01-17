package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetExtension(i interface{}) string {
	var v string
	raw, ok := i.(json.RawMessage)
	if !ok {
		log.Fatalln("expected json.RawMessage, received something else")
	}
	json.Unmarshal(raw, &v)
	return v
}

var funcMap template.FuncMap

// var fileStructure map[string]interface{}
var root Dir

type Dir struct {
	Name      string
	ChildDirs map[string]Dir
}

func init() {
	funcMap = make(template.FuncMap, 0)
	funcMap["GetExtension"] = GetExtension

	root := map[string]interface{}{
		"server": map[string]interface{}{
			"ui": map[string]interface{}{
				"controller"
			}
		}
	}
}

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
	// tmpl := template.New("controller").Funcs(funcMap)
	// // Create a new template by parsing a file
	// parsedTmpl, err := tmpl.Parse("controller.tmpl")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(parsedTmpl)

	parsedTmpl, err := template.New("controller.tmpl").Funcs(funcMap).ParseFiles("controller.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	groupedOperations := make(map[string][]*openapi3.Operation, 0)
	for _, pathItem := range document.Paths {
		operations := pathItem.Operations()
		for key, operation := range operations {
			fmt.Println("=========KEY")
			fmt.Println(key)
			fmt.Println("=========KEY")
			tag := operation.Tags[0]
			if _, ok := groupedOperations[tag]; !ok {
				groupedOperations[tag] = []*openapi3.Operation{operation}
			} else {
				groupedOperations[tag] = append(groupedOperations[tag], operation)
			}
		}
		// fmt.Println("========================================")
		// fmt.Println(len(groupedOperations["photo"]))
		// fmt.Println(groupedOperations["photo"][0])
		// fmt.Println(groupedOperations["photo"][1])
		// fmt.Println("========================================\n")
	}

	caser := cases.Title(language.Und)

	for tag, op := range groupedOperations {
		fmt.Println(tag)
		fmt.Println(len(op))

		file, err := os.OpenFile(fmt.Sprintf("%sController.go", tag), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		fmt.Println(caser.String(tag))
		if err := parsedTmpl.Execute(file, struct {
			Name       string
			Operations []*openapi3.Operation
		}{
			Name:       caser.String(tag),
			Operations: op,
		}); err != nil {
			log.Fatalln(err)
		}
	}

	// fmt.Println(values)
	// Open or Create the output file
	// file, err := os.OpenFile("server.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer file.Close()

	// if err := tmpl.Execute(file, document); err != nil {
	// 	log.Fatalln(err)
	// }

}
