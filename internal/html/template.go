package html

import (
	"fmt"
	"html/template"
)

func GetTemplate(name string) (*template.Template, error) {
	file := fmt.Sprintf("web/templates/%s.html", name)
	return template.ParseFiles("web/layout/layout.html", file)
}

func GetComponent(name string) (*template.Template, error) {
	file := fmt.Sprintf("web/components/%s.html", name)
	return template.ParseFiles(file)
}
