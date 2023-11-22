package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

type TemplateOptions struct {
	templateName string
	layoutName   string
}

func getComponentPaths() []string {
	components, error := os.ReadDir(path.Join("views", "component"))
	componentPaths := []string{}

	if error != nil {
		log.Fatal(error)
	}

	for _, compoPath := range components {
		componentPaths = append(componentPaths, path.Join("views", "component", compoPath.Name()))
	}

	return componentPaths
}

func renderTemplate(options TemplateOptions, data any, w http.ResponseWriter) error {

	templatePath := path.Join("views", options.templateName+".gohtml")
	var layoutPath string
	if options.layoutName == "" {
		layoutPath = path.Join("views", "layouts", "default.gohtml")
	} else {
		layoutPath = path.Join("view", "layouts", options.layoutName+".gohtml")
	}

	components := getComponentPaths()

	w.Header().Add("Content-Type", "text/html")

	templateFiles := []string{layoutPath, templatePath}

	templateFiles = append(templateFiles, components...)

	t, err := template.ParseFiles(templateFiles...)

	if err != nil {
		fmt.Println("Error occurred")
		fmt.Println(err)
		return err
	}

	return t.Execute(w, data)

}
