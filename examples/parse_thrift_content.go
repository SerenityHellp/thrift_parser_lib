package main

import (
	"github.com/SerenityHellp/thrift_parser_lib/parser"
	"strings"
	"fmt"
)


func parse (){
	parser := parser.Parser{}
	contents := `
		include "other.thrift"

		namespace go somepkg
		namespace python some.module123
		namespace python.py-twisted another

		const map<string,string> M1 = {"hello": "world", "goodnight": "moon"}
		const string S1 = "foo\"\tbar"
		const string S2 = 'foo\'\tbar'
		const list<i64> L = [1, 2, 3];

		union myUnion
		{
			1: double dbl = 1.1;
			2: string str = "2";
			3: i32 int32 = 3;
			4: i64 int64
				= 5;
		}

		enum Operation
		{
			ADD = 1,
			SUBTRACT = 2
		}

		enum NoNewLineBeforeBrace {
			ADD = 1,
			SUBTRACT = 2
		}

		service ServiceNAME extends SomeBase
		{
			# authenticate method
			// comment2
			/* some other
			   comments */
			string login(1:string password) throws (1:AuthenticationException authex),
			oneway void explode();
			blah something()
		}

		struct SomeStruct {
			1: double dbl = 1.2,
			2: optional string abc
		}

		struct NewLineBeforeBrace
		{
			1: double dbl = 1.2,
			2: optional string abc
		}`
	reader := strings.NewReader(contents)
	thrift, err := parser.Parse(reader)

	fmt.Println(thrift)
	fmt.Println(err)
}

func main()  {
	parse()
}