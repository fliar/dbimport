package main
import(
	"fmt"
	"flag"
)

func main(){
	fmt.Println("mysql import")
	folder := flag.String(
		"f",
		"c:/temp/root",
		"folder holding files to import")
	dbinfo := flag.String(
		"d",
		"user:pass@tcp(192.168.1.114:3306)/test",
		"MySQL db connection info")
	flag.Parse()
	fmt.Println("input folder is ", *folder)
	fmt.Println("connecting to ", *dbinfo)
	
	dbcon := new(DbConnect)
	dbcon.Connect(*dbinfo)
	dbcon.AddFile("files", "aa", "txt", "abcdefg")
	walk := Walker{*folder}
	walk.Run()
}
