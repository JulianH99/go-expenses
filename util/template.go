package util

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

type TemplateOptions struct {
	TemplateName string
	LayoutName   string
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

func RenderTemplate(options TemplateOptions, data any, w http.ResponseWriter) error {

	templatePath := path.Join("views", options.TemplateName+".gohtml")
	var layoutPath string
	if options.LayoutName == "" {
		layoutPath = path.Join("views", "layouts", "default.gohtml")
	} else {
		layoutPath = path.Join("view", "layouts", options.LayoutName+".gohtml")
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
