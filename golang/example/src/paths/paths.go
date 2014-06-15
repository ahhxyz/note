import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
)


	file, _ := os.Getwd() //可执行文件所在的目录
	fmt.Fprintf(c.Ctx.Res, "<h1 style=\"color\" >current path:", file, "</h1>")

	file, _ = exec.LookPath(os.Args[0])
	fmt.Fprintf(c.Ctx.Res, "<h1 style=\"color\" >exec path:", file, "</h1>")

	dir, _ := path.Split(file)
	fmt.Fprintf(c.Ctx.Res, "<h1 style=\"color\" >exec folder relative path:", dir, "</h1>")

	os.Chdir(dir)
	wd, _ := os.Getwd()
	fmt.Fprintf(c.Ctx.Res, "<h1 style=\"color\" >exec folder absolute path:", wd, "</h1>")
