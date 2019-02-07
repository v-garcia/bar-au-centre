package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const LISTENING_PORT = 8086

// All fiels will be loaded in memory (for now)
var inMemoryFiles map[string]ResponseFile

// The in memory representation of a file
type ResponseFile struct {
	mime    string
	path    string
	content []byte
}

// Scan for every files in the SERVING_PATH
func scanForFiles(basePath string) map[string]ResponseFile {
	var files = make(map[string]ResponseFile)

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {

		// Don't consider folders
		if info.IsDir() {
			return nil
		}

		// Read the whole file
		content, errReadFile := ioutil.ReadFile(path)

		if errReadFile != nil {
			fmt.Println("errReadFile")
			return errReadFile
		}

		// Get MIME content type
		mime, errMime := FileContentType(path)

		if errMime != nil {
			fmt.Println("errRel")
			return errMime
		}

		// Make path relative to SERVING_PATH
		relPath, errRel := filepath.Rel(basePath, path)

		if errRel != nil {
			fmt.Println("errRel")
			return errRel
		}

		relPath = "/" + strings.ToLower(relPath)
		fmt.Printf("%s loaded in memory \n", relPath)

		// Fill ResponseFile struct and save it in the map structure
		fileStruct := ResponseFile{
			mime:    mime,
			path:    relPath,
			content: content,
		}

		files[fileStruct.path] = fileStruct

		return nil
	})

	panicOnError(err)

	return files
}

// Custom http query handler
// Always returns index.html on unknow route
// Letting the browser handling routes
func serveFileHandler(w http.ResponseWriter, r *http.Request) {
	// Only respond on get query
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("This server only support 'GET' http verb"))
		return
	}

	// Handle liveness probe
	if r.URL.String() == "/healthz" {
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
		return
	}

	fmt.Printf("wanted url: '%s'\n", r.URL.String())

	// Unescape query
	wantedUrl, errUnescaping := url.QueryUnescape(r.URL.String())
	panicOnError(errUnescaping)

	wantedUrl = strings.ToLower(wantedUrl)

	fileToServe, keyFound := inMemoryFiles[wantedUrl]

	// On not found, return index.html
	if !keyFound {
		fileToServe = inMemoryFiles["/index.html"]
	}

	fmt.Printf("served file: '%s'\n", fileToServe.path)

	serveFile(fileToServe, w)
}

func serveFile(file ResponseFile, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", file.mime)
	w.Write(file.content)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Load the simple HTTP Server")

	// Get serving port
	var listeningPort = getenv("LISTENING_PORT", "8080")
	_, errParsePort := strconv.Atoi(listeningPort)
	panicOnError(errParsePort)
	fmt.Printf("Listening port: %s\n", listeningPort)

	// Get path to serve
	var servingPath = getenv("SERVING_PATH", "/www")
	fmt.Printf("Serving path: '%s'\n", servingPath)
	_, errDirectory := os.Stat(servingPath)
	panicOnError(errDirectory)

	// Load all files
	inMemoryFiles = scanForFiles(servingPath)
	fmt.Printf("%d files loaded\n", len(inMemoryFiles))

	// Listen for queries
	err := http.ListenAndServe(":"+listeningPort, http.HandlerFunc(serveFileHandler))
	panicOnError(err)

}
