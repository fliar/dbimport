package main
import(
	"os"
	"fmt"
	"path/filepath"
)

type Walker struct {
	RootPath string
}

func (w Walker) walk(path string, f os.FileInfo, err error) error {
	fmt.Println("walking ", path)
	return err
}

func (w Walker) MakeWalker() func(path string, f os.FileInfo, err error) error {
	f := func(path string, f os.FileInfo, err error) error {
		return w.walk(path,f,err)
	}
	return f
}

func (w Walker) Run() {
	err := filepath.Walk(w.RootPath, w.MakeWalker())
	if err != nil {
		panic(err.Error())
	}
}
