package main

import (
	_ "embed"
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"

	"github.com/corona10/goimagehash"
)

//go:embed index.html
var indexHTML []byte

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/phash", handlePhash)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(indexHTML)
}

func handlePhash(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	const maxSize = 10 << 20 // 10MB
	r.Body = http.MaxBytesReader(w, r.Body, maxSize)
	if err := r.ParseMultipartForm(maxSize); err != nil {
		writeJSONError(w, "failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		writeJSONError(w, "missing or invalid form field 'image': "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		writeJSONError(w, "failed to decode image: "+err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := goimagehash.PerceptionHash(img)
	if err != nil {
		writeJSONError(w, "failed to compute phash: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"phash": hash.ToString(),
		"hash":  hash.GetHash(),
	})
}

func writeJSONError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
