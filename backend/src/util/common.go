package util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type ExecRes struct {
	Code    int    `json:"code"`    // 0 : success; 1 : fail
	Message string `json:"message"` // success|fail
	Data    string `json:"data"`    // output of execution
}

func SaveStrToFile(fileName, fileBody string) error {
	workDir, _ := os.Getwd()
	fileRelPath := filepath.Join(workDir, "src", "var", fileName)
	print("save file dir:" + fileRelPath)
	f, err := os.Create(fileRelPath) //create the code file locally
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fileBody)
	f.Sync()

	return err
}

func CmdAndRunFile(filePath string) ExecRes {
	cmd := exec.Command("go", "run", filePath)
	var ret ExecRes
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error out:\n%s\n", string(out))
		ret.Code = 1
		ret.Message = "fail"
		ret.Data = string(out)

		return ret
	}
	ret.Code = 0
	ret.Message = "success"
	ret.Data = string(out)
	fmt.Printf("combined out:\n%s\n", string(out))
	return ret
}
