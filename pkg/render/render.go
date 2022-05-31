package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	fmt.Println("rendering ", tmpl)
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing templates", err)
		return
	}
}
