package util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type ExecRes struct {
	Code string // 运行结果 succ or fail
	Res  string // 输出结果
}

func SaveStrToFile(fileName, fileBody string) error {
	workDir, _ := os.Getwd()
	fileRelPath := filepath.Join(workDir, "src", "var", "template.go")

	f, err := os.Create(fileRelPath) //创建文件
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
		ret.Code = "fail"
		ret.Res = string(out)

		return ret
	}
	ret.Code = "succ"
	ret.Res = string(out)
	fmt.Printf("combined out:\n%s\n", string(out))
	return ret
}
