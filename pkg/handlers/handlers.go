package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jwmurray/udemy/pkg/render"
	"github.com/jwmurray/udemy/pkg/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

func CallingList(w http.ResponseWriter, r *http.Request) {
	filename := "callings.json"
	var err error

	buf, err := os.ReadFile(filename)
	string_buf := string(buf)

	utils.Log_on_error(err, "failure on file read")

	_, err = fmt.Fprintf(w, string_buf)

	utils.Log_on_error(err, "No message yet")
}

func WardList(w http.ResponseWriter, r *http.Request) {
	filename := "members.json"
	var err error

	buf, err := os.ReadFile(filename)
	string_buf := string(buf)

	utils.Log_on_error(err, "failure on file read")

	_, err = fmt.Fprintf(w, string_buf)

	utils.Log_on_error(err, "No message yet")
}
