package main
import(
	"os"
	"fmt"
	"path/filepath"
	"io/ioutil"
)

type Walker struct {
	RootPath string
	db DbConnect
}

func (w Walker) walk(path string, f os.FileInfo, err error) error {
	fmt.Println("walking ", path)
	slashPath := filepath.ToSlash(path)
	fmt.Println("converted path: ", slashPath)
	if f.IsDir() {return err}
	fmt.Println("it is a file")
	fmt.Println(path, "has an ext called ", filepath.Ext(path))
	buf, e := ioutil.ReadFile(path)
	if e != nil {
		panic(e.Error())
	}
	content := string(buf)
	relpath, e := filepath.Rel(w.RootPath, path)
	if e != nil {
		fmt.Println("retrieving relative path error: ", e.Error())
		relpath = path
	}
	fmt.Println("relative path: ", relpath)
	w.db.ReplaceFile("files", filepath.ToSlash(relpath), filepath.Ext(path), content)
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
