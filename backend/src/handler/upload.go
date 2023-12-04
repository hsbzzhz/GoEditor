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

		ret := map[string]interface{}{}

		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		var codebody CodeBody
		json.Unmarshal(body, &codebody)

		// 创建文件
		err := util.SaveStrToFile("template.go", codebody.Code)
		if err != nil {
			print(err)
		}
		// 执行文件，获取结果
		workDir, _ := os.Getwd()
		filePath := filepath.Join(workDir, "src", "var", "template.go")
		print("run file dir:" + filePath)
		res := util.CmdAndRunFile(filePath)
		// 对解析得到对body进行操作
		ret["code"] = res.Code
		ret["result"] = res.Res

		str, _ := json.Marshal(ret)
		w.Write(str)
	}
}
