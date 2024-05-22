package main

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/dsnet/compress/bzip2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/pflag"
	"io/ioutil"
	"log"
	"vitess.io/vitess/go/sqltypes"
	"vitess.io/vitess/go/vt/sqlparser"
)

var (
	command           = "install"
	wasmFile          = "./bin/myguest.wasm"
	wasmName          = "test"
	runtime           = "wazero"
	compressAlgorithm = "bzip2"

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

func CompressByBZip2(originalData []byte) []byte {
	var buf bytes.Buffer
	w, err := bzip2.NewWriter(&buf, &bzip2.WriterConfig{Level: bzip2.BestCompression})
	if err != nil {
		log.Fatal(err)
	}
	if _, err := w.Write(originalData); err != nil {
		log.Fatal(err)
	}
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func init() {
	pflag.StringVar(&command, "command", command, "the command of the script. default is install")
	pflag.StringVar(&wasmFile, "wasm_file", wasmFile, "the wasmFile of wasm file")
	pflag.Parse()
}

func main() {
	switch command {
	case "install":
		install()
	case "uninstall":
		uninstall()
	}
}

func calcMd5String32(data []byte) string {
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}

func install() {
	wasmBytes, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		log.Panicf("error when reading wasm bytes: %v", err)
	}
	fmt.Printf("before compress:\n")
	fmt.Printf("bytes num is %d\n", len(wasmBytes))
	fmt.Printf("last 5 bytes is %v %v %v %v %v\n", wasmBytes[len(wasmBytes)-5], wasmBytes[len(wasmBytes)-4], wasmBytes[len(wasmBytes)-3], wasmBytes[len(wasmBytes)-2], wasmBytes[len(wasmBytes)-1])

	hash := calcMd5String32(wasmBytes)
	fmt.Printf("hash is %v\n", hash)

	wasmBytes = CompressByBZip2(wasmBytes)
	fmt.Printf("after compress:\n")
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

	insertWasmTemplate := `insert ignore into mysql.wasm_binary(name,runtime,data,compress_algorithm,hash_before_compress) values (%a,%a,%a,%a,%a);`
	insertWasmSQL, err := sqlparser.ParseAndBind(insertWasmTemplate,
		sqltypes.StringBindVariable(wasmName),
		sqltypes.StringBindVariable(runtime),
		sqltypes.BytesBindVariable(wasmBytes),
		sqltypes.StringBindVariable(compressAlgorithm),
		sqltypes.StringBindVariable(hash))
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Query(insertWasmSQL)
	fmt.Printf("insert sql len %v\n", len(insertWasmSQL))
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

	_, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func uninstall() {
	dropFilterSQL := fmt.Sprintf("drop filter %s", filterName)
	deleteWasmSQL := fmt.Sprintf("delete from mysql.wasm_binary where name='%s'", wasmName)

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:15306)/mysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic("database can't connect")
	}
	fmt.Println("database connected")

	{
		r, err := db.Exec(dropFilterSQL)
		if err != nil {
			panic(err.Error())
		}
		if affected, _ := r.RowsAffected(); affected == 0 {
			fmt.Printf("filter %s not found\n", filterName)
		}
	}

	{
		r, err := db.Exec(deleteWasmSQL)
		if err != nil {
			panic(err.Error())
		}
		if affected, _ := r.RowsAffected(); affected == 0 {
			fmt.Printf("wasm %s not found\n", wasmName)
		}
	}
}
