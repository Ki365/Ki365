package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func uploadProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading file...")
	fmt.Println(r.Header)
	fmt.Println(r.Body)
	fmt.Println(r.Host)
	fmt.Println(r.Method)
	fmt.Println(r.Response)
	fmt.Println(r.Trailer)

	// Fixed 167.7MB limit on file upload
	r.ParseMultipartForm(10 << 2)

	file, header, err := r.FormFile("project-archive")

	if err != nil {
		fmt.Println("Error retrieving the uploaded file, canceling upload...")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded filename: %+v\n", header.Filename)
	fmt.Printf("Filesize: %+v\n", header.Size)
	fmt.Printf("MIME: header: %+v\n", header.Header)

	dst, err := os.Create(filepath.Join("./", header.Filename))
	if err != nil {
		fmt.Println(err)
	}
	defer dst.Close()

	fileBytes, err := io.Copy(dst, file)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fileBytes)
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)

	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	port := flag.String("p", "8100", "port to serve directory on")
	directory := flag.String("d", ".", "directory of files to serve")
	flag.Parse()

	router := mux.NewRouter()

	spa := spaHandler{staticPath: *directory, indexPath: "index.html"}
	router.HandleFunc("/api/upload/project", uploadProject)
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + *port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(srv.ListenAndServe())
}
