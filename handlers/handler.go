package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"ascii-art-export/functions"
)

var tp1 *template.Template

type Table struct {
	Error string
}

func Download_Handler(w http.ResponseWriter, r *http.Request) {
	input := r.PostFormValue("inputValue")
	Banner := r.PostFormValue("bannerValue")
	output := functions.Web_ascii_art(input, Banner)

	// Set the headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=Ascii-art.txt")
	w.Header().Set("Content-Type: text/html", r.Header.Get("Content-Type"))

	// Read your file content and write it to the response writer
	// Example:
	fileContent := []byte(output)
	w.Write(fileContent)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found!", http.StatusMethodNotAllowed)
		return
	}
	tp1, _ = template.ParseFiles("templates/Homepage.html")
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	tp1.Execute(w, nil)
}

func OutputHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Output string
		Input  string
		Banner string
	}

	tp2, err := template.ParseFiles("templates/Outputpage.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	input := r.PostFormValue("input")
	if len(input) > 11076 {
		w.WriteHeader(http.StatusBadRequest)
		tp1.ExecuteTemplate(w, "Homepage.html", nil)
		return
	}

	if input == "" {
		data := Table{
			Error: "The input area is empty",
		}
		tp1.ExecuteTemplate(w, "Homepage.html", data)
		return
	}

	Banner := r.PostFormValue("template")

	if !functions.Check_Banner(Banner) {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Page Not Found!", 404)
		return
	}

	executeoutput := Output{
		Output: "\n" + functions.Web_ascii_art(input, Banner),
		Input:  input,
		Banner: Banner,
	}

	err = tp2.ExecuteTemplate(w, "Outputpage.html", executeoutput)
	if err != nil {
		fmt.Println(err)
		return
	}
}
