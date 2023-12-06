package handler

import (
	"editorDemo/src/util"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type CodeBody struct {
	Code string `json:"code"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-requested-with")
		return
	}
	if r.Method == "POST" {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("content-type", "application/json")

		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		var codebody CodeBody
		json.Unmarshal(body, &codebody)

		// receive the code and save as local file
		util.SaveStrToFile("template.go", codebody.Code)

		// run file and get the results
		workDir, _ := os.Getwd()
		filePath := filepath.Join(workDir, "src", "var", "template.go")
		print("run file dir:" + filePath)
		res := util.CmdAndRunFile(filePath)

		str, _ := json.Marshal(res)
		w.Write(str)
	}
}
