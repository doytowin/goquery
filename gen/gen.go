/*
 * The Clear BSD License
 *
 * Copyright (c) 2025, DoytoWin, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	// 定义命令行参数
	generatorType := flag.String("type", "./sql", "Generated type.")
	inputFile := flag.String("f", "./user.go", "The Go file containing the query definition")
	outputFile := flag.String("o", "./query_builder.go", "The Go file to output the query builder")

	// 解析命令行参数
	flag.Parse()

	var generator Generator
	switch *generatorType {
	case "sql":
		generator = NewSqlGenerator()
	case "mongodb":
		generator = NewMongoGenerator()
	}

	// 调用代码生成方法
	err := GenerateQueryBuilder(generator, *inputFile, *outputFile)
	if err != nil {
		log.Fatalf("Error generating query builder: %v", err)
	}

	fmt.Println("Query builder generated successfully!")
}
