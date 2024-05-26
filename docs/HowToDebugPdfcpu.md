# 源代码

```
git clone https://github.com/pdfcpu/pdfcpu.git

```

# 如何 build 文件

参考：http://events.jianshu.io/p/aaa07f767d8e

```
cd cmd/pdfcpu

#linux
go build

# window64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

生成pdfcpu.exe ， copy至window目录下，进行测试，命令如下：
 .\pdfcpu.exe info .\in2.pdf

```

# 如何调试文件

/pkg/pdfcpu/gf 放测试 pdf 文件
/pkg/pdfcpu/api/test/pdf_view_fix_test.go 写单元测试

```
package test

import (
	// "fmt"
	// "io"
	// "os"
	"path/filepath"
	// "strings"
	"testing"
	// "time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	// "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

// go test -timeout 3600s -v -run ^Test_Info$ github.com/pdfcpu/pdfcpu/pkg/api/test
func Test_Info(t *testing.T) {
	inDir = filepath.Join("..", "..", "gf")
	msg := "TestInfo"
	inFile := filepath.Join(inDir, "in4.pdf")

	if _, err := api.InfoFile(inFile, nil, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}
}

```

# 修改记录
1， 20240526 fix Unescape: illegal escape sequence \P detected: <Created By HP MFP (\PDF - 300X300 dpi)>