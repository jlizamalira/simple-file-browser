package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// File struct
type File struct {
	Name      string
	URL       string
	Directory bool
}

// Folder struct
type Folder struct {
	Files []File
}

// AddItem - add files to folder
func (folder *Folder) AddItem(file File) {
	folder.Files = append(folder.Files, file)
}

// PageData - page struct
type PageData struct {
	PageTitle string
	Files     []File
	PageBody  string
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/directory", handler2)

	http.ListenAndServe(":80", nil)

}

func getFiles(url string) Folder {

	var url2 string
	var dir bool

	box := Folder{}
	files, err := ioutil.ReadDir(url)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		if f.IsDir() {
			url2 = "/directory?key=" + url + "/" + f.Name()
			dir = true
		} else {
			url2 = f.Name()
			dir = false
		}

		item := File{Name: f.Name(), URL: url2, Directory: dir}

		box.AddItem(item)

	}
	return box

}

func handler(w http.ResponseWriter, r *http.Request) {

	carpeta := getFiles("/Users")
	tmpl := template.Must(template.ParseFiles("pag.html"))

	data := PageData{
		PageTitle: "My Documents:",

		Files:    carpeta.Files,
		PageBody: "/Users",
	}
	tmpl.Execute(w, data)

}

func handler2(w http.ResponseWriter, r *http.Request) {

	keys := r.URL.Query()["key"]
	var dir string

	dir = keys[0]

	carpeta := getFiles(dir)

	tmpl := template.Must(template.ParseFiles("pag.html"))

	data := PageData{
		PageTitle: "My Documents:",

		Files:    carpeta.Files,
		PageBody: dir,
	}
	tmpl.Execute(w, data)

}
