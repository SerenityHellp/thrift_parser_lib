package main

import (
	"path/filepath"
	"github.com/SerenityHellp/thrift_parse_lib/parser"
	"fmt"
)


func main()  {
	parser := parser.Parser{}
	file_name :=  "cassandra.thrift"
	//thrift文件；绝对路径；错误信息
	thrift,abs_path, err := parser.ParseFile(filepath.Join("testfiles", file_name))

	fmt.Println(thrift)
	fmt.Println(abs_path)
	fmt.Println(err)
}
