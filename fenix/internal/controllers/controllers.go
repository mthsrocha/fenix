package controllers

import (
	"html/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))