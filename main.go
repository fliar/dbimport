package main
import(
	"fmt"
	"flag"
)

func main(){
	fmt.Println("mysql import")
	folder := flag.String("f", "./", "folder holding files to import")
	dbinfo := flag.String("d", "", "MySQL db connection info")
	flag.Parse()
	*dbinfo = "root:juxienet@tcp(localhost:3306)/test"
	fmt.Println("input folder is ", *folder)
	fmt.Println("connecting to ", *dbinfo)
	
	dbcon := new(DbConnect)
	dbcon.Connect("root:juxienet@tcp(localhost:3306)/test")
	dbcon.AddFile("files", "aa/bb", "txt", "hihihihihi")
}
