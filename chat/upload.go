package chat

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

type uploaderHandler struct {
	path string
}

// UploaderHandler hadnles upload of images process
func UploaderHandler(path string) http.Handler {
	return &uploaderHandler{
		path: path,
	}
}

func (uh *uploaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("uploaderHandler:ServeHTTP ")

	userID := r.FormValue("user_id")
	file, header, err := r.FormFile("avatar_file")
	if err != nil {
		log.Printf("FormFile: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("ReadAll: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := strings.Join([]string{uh.path, userID + path.Ext(header.Filename)}, string(filepath.Separator))
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		log.Printf("WriteFile: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = io.WriteString(w, "Successful"); err != nil {
		log.Printf("WriteString: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
