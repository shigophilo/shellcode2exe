# shellcode2exe

Cobalt Strike Raw格式

+ 将shellcode编译成exe, 并与另一个文件"捆绑"

```
  -eo string
        真实马输出的路径和文件名 (default "exe.exe")
  -i string
        payload.bin文件 (default "payload.bin")
  -o string
        最后输出的exe文件名 (default "out.exe")
  -w string
        捆绑的word.docx (default "新建 DOCX 文档.docx")
  -wn string
        打开的word文件名 (default "test.docx")
```

运行 -o 生成的exe文件, 会同时打开"捆绑的" exe文件和 .docx文件