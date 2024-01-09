package main

import (
	sql_queries_builder "doit-utils/utils"
	"fmt"
)

func main() {
	sql_queries_builder.AddQueries("check.sql")
	c1 := sql_queries_builder.GetQuery("check")
	fmt.Println(c1)
	c2 := sql_queries_builder.GetQuery("check1")
	fmt.Println(c2)
}
