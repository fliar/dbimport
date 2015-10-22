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
	fmt.Println("input folder is ", *folder)
	fmt.Println("connecting to ", *dbinfo)
}
