package main

import (
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

// Determine a file's content type, like: text/plain; charset=utf-8
// Inspired by go's serveContent() in net/http package.
// More information: src/pkg/net/http/fs.go
func FileContentType(filename string) (string, error) {
	ctype := mime.TypeByExtension(filepath.Ext(filename))
	if ctype == "" {

		file, err := os.Open(filename)
		if err != nil {
			return "", err
		}
		defer file.Close()

		ctype, err = ContentType(file)
		if err != nil {
			return "", err
		}

	}

	return ctype, nil
}

// Determine a data stream's content type, like: text/plain; charset=utf-8
// Inspired by go's serveContent() in net/http package.
// More information: src/pkg/net/http/fs.go
func ContentType(reader io.ReadSeeker) (string, error) {
	data := make([]byte, 512)

	_, err := reader.Seek(0, os.SEEK_SET)
	if err != nil {
		return "", err
	}

	n, err := reader.Read(data)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(data[:n]), nil
}
