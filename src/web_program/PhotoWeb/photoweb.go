package main

import (
	"net/http"
	// "fmt"
	"io"
	"log"
	"os"
	"html/template"
	"io/ioutil"
	_ "net/http/pprof"  
)

// import _ "net/http/pprof"  可以查看 系统运行状态
// 方法1， 通过web端访问： http://localhost:18082/debug/pprof/
// 方法2， 通过命令行，命令执行， go tool pprof分析性能， web端查看


const (
	UPLOAD_DIR = "./uploads"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath);!exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
    if err == nil {
    	return true
    }
	return os.IsExist(err)
}

func uploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		// io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
		// " enctype=\"multipart/form-data\">"+
		// "Choose an image to upload: <input name=\"image\" type=\"file\" />"+
		// "<input type=\"submit\" value=\"Upload\" />"+
		// "</form>")
		t, err := template.ParseFiles("upload.html")
		if err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return 
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
			http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
			http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),
			http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(),
		http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images 
	t, err := template.ParseFiles("list.html")
	if err != nil {
		http.Error(w, err.Error(),
		http.StatusInternalServerError)
		return
	}
	t.Execute(w, locals)
}

// func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) err error {
	// t, err = template.ParseFiles(tmpl + ".html")
	// if err != nil {
		// return err
	// }
	// err = t.Execute(w, locals)
	// return nil
// }

func main(){
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":18081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

