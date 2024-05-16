package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

const path = "./wazero/myGuest.wasm"

var (
	wasmName = "test"
	runtime  = "wazero"

	filterName               = "wasm"
	desc                     = "wasm test"
	priority                 = "999"
	status                   = "ACTIVE"
	plans                    = ""
	fullyQualifiedTableNames = ""
	queryRegex               = ``
	queryTemplate            = ""
	requestIpRegex           = ``
	userRegex                = ``
	leadingCommentRegex      = ``
	trailingCommentRegex     = ``
	bindVarConds             = ""
	action                   = "wasm_plugin" // todo remove

)

func main() {
	wasmBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("error when reading wasm bytes: %v", err)
	}
	fmt.Printf("bytes num is %d\n", len(wasmBytes))
	fmt.Printf("last 5 bytes is %v %v %v %v %v\n", wasmBytes[len(wasmBytes)-5], wasmBytes[len(wasmBytes)-4], wasmBytes[len(wasmBytes)-3], wasmBytes[len(wasmBytes)-2], wasmBytes[len(wasmBytes)-1])

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:15306)/mysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic("database can't connect")
	}
	fmt.Println("database connected")

	// todo newborn22 5.14 blob or base64?
	binary := fmt.Sprintf("%v", wasmBytes)
	binary = binary[1 : len(binary)-1]
	insertWasmTemplate := `insert ignore into mysql.wasm_binary(name,runtime,data) values ('%s','%s','%s');`
	insertWasmSQL := fmt.Sprintf(insertWasmTemplate, wasmName, runtime, binary)
	_, err = db.Query(insertWasmSQL)
	if err != nil {
		panic(err.Error())
	}

	createFilterTemplate := `create filter if not exists %s (
	  desc='%s',
	  priority='%s',
	  status='%s'
		)
		with_pattern(
				plans='%s',
				fully_qualified_table_names='%s',
				query_regex='%s',
				query_template='%s',
				request_ip_regex='%s',
				user_regex='%s',
				leading_comment_regex='%s',
				trailing_comment_regex='%s',
				bind_var_conds='%s'
		)
		execute(
				action='%s',
				action_args='%s'
		);`

	actionArgs := fmt.Sprintf("wasm_binary_name=\"%v\"", wasmName)
	query := fmt.Sprintf(createFilterTemplate,
		filterName,
		desc,
		priority,
		status,
		plans,
		fullyQualifiedTableNames,
		queryRegex,
		queryTemplate,
		requestIpRegex,
		userRegex,
		leadingCommentRegex,
		trailingCommentRegex,
		bindVarConds,
		action,
		actionArgs)

	//	createFilterTemplate := ` INSERT IGNORE INTO mysql.wescale_plugin (name, description, priority, status, plans, fully_qualified_table_names, query_regex, query_template, request_ip_regex, user_regex, leading_comment_regex, trailing_comment_regex, bind_var_conds, action, action_args)
	//VALUES ('%s', '%s', %d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s');`
	//
	//	actionArgs := fmt.Sprintf("%v", wasmBytes)
	//	actionArgs = actionArgs[1 : len(actionArgs)-1]
	//	query := fmt.Sprintf(createFilterTemplate,
	//		filterName,
	//		desc,
	//		999,
	//		status,
	//		plans,
	//		fullyQualifiedTableNames,
	//		queryRegex,
	//		queryTemplate,
	//		requestIpRegex,
	//		userRegex,
	//		leadingCommentRegex,
	//		trailingCommentRegex,
	//		bindVarConds,
	//		action,
	//		actionArgs)
	//
	_, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
