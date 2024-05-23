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
)

var (
	command = "install"
)

// mysql config
var (
	mysqlHost     = "127.0.0.1"
	mysqlPort     = 15306
	mysqlUser     = "root"
	mysqlPassword = ""
	mysqlDb       = "mysql"
)

// wasm config
var (
	wasmFile              = "./bin/myguest.wasm"
	wasmName              = "test"
	wasmRuntime           = "wazero"
	wasmCompressAlgorithm = "bzip2"
)

// filter config
var (
	filterName                     = "wasm"
	filterDesc                     = "wasm test"
	filterPriority                 = "999"
	filterStatus                   = "ACTIVE"
	filterPlans                    = ""
	filterFullyQualifiedTableNames = ""
	filterQueryRegex               = ``
	filterQueryTemplate            = ""
	filterRequestIpRegex           = ``
	filterUserRegex                = ``
	filterLeadingCommentRegex      = ``
	filterTrailingCommentRegex     = ``
	filterBindVarConds             = ""
	filterAction                   = "wasm_plugin" // todo remove
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

	pflag.StringVar(&mysqlUser, "mysql_user", mysqlUser, "the user of mysql")
	pflag.StringVar(&mysqlPassword, "mysql_password", mysqlPassword, "the password of mysql")
	pflag.StringVar(&mysqlHost, "mysql_host", mysqlHost, "the host of mysql")
	pflag.IntVar(&mysqlPort, "mysql_port", mysqlPort, "the port of mysql")
	pflag.StringVar(&mysqlDb, "mysql_db", mysqlDb, "the db of mysql")

	pflag.StringVar(&wasmFile, "wasm_file", wasmFile, "the wasmFile of wasm file")
	pflag.StringVar(&wasmName, "wasm_name", wasmName, "the wasmName of wasm")
	pflag.StringVar(&wasmRuntime, "wasm_runtime", wasmRuntime, "the wasm_runtime of wasm")
	pflag.StringVar(&wasmCompressAlgorithm, "wasm_compress_algorithm", wasmCompressAlgorithm, "the wasm_compress_algorithm of wasm")

	pflag.StringVar(&filterName, "filter_name", filterName, "the filter_name of filter")
	pflag.StringVar(&filterDesc, "filter_desc", filterDesc, "the filter_desc of filter")
	pflag.StringVar(&filterPriority, "filter_priority", filterPriority, "the filter_priority of filter")
	pflag.StringVar(&filterStatus, "filter_status", filterStatus, "the filter_status of filter")
	pflag.StringVar(&filterPlans, "filter_plans", filterPlans, "the filter_plans of filter")
	pflag.StringVar(&filterFullyQualifiedTableNames, "filter_fully_qualified_table_names", filterFullyQualifiedTableNames, "the filter_fully_qualified_table_names of filter")
	pflag.StringVar(&filterQueryRegex, "filter_query_regex", filterQueryRegex, "the filter_query_regex of filter")
	pflag.StringVar(&filterQueryTemplate, "filter_query_template", filterQueryTemplate, "the filter_query_template of filter")
	pflag.StringVar(&filterRequestIpRegex, "filter_request_ip_regex", filterRequestIpRegex, "the filter_request_ip_regex of filter")
	pflag.StringVar(&filterUserRegex, "filter_user_regex", filterUserRegex, "the filter_user_regex of filter")
	pflag.StringVar(&filterLeadingCommentRegex, "filter_leading_comment_regex", filterLeadingCommentRegex, "the filter_leading_comment_regex of filter")
	pflag.StringVar(&filterTrailingCommentRegex, "filter_trailing_comment_regex", filterTrailingCommentRegex, "the filter_trailing_comment_regex of filter")
	pflag.StringVar(&filterBindVarConds, "filter_bind_var_conds", filterBindVarConds, "the filter_bind_var_conds of filter")
	pflag.StringVar(&filterAction, "filter_action", filterAction, "the filter_action of filter")

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

	db, err := sql.Open("mysql", generateMysqlDsn())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic("database can't connect")
	}
	fmt.Println("database connected")

	insertIntoWasmBinary(db, wasmName, wasmRuntime, wasmCompressAlgorithm, wasmBytes, hash)

	createFilterTemplate := `create filter if not exists %s (
	  filterDesc='%s',
	  filterPriority='%s',
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
		filterDesc,
		filterPriority,
		filterStatus,
		filterPlans,
		filterFullyQualifiedTableNames,
		filterQueryRegex,
		filterQueryTemplate,
		filterRequestIpRegex,
		filterUserRegex,
		filterLeadingCommentRegex,
		filterTrailingCommentRegex,
		filterBindVarConds,
		filterAction,
		actionArgs)

	_, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func uninstall() {
	dropFilterSQL := fmt.Sprintf("drop filter %s", filterName)
	deleteWasmSQL := fmt.Sprintf("delete from mysql.wasm_binary where name='%s'", wasmName)

	db, err := sql.Open("mysql", generateMysqlDsn())
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

func insertIntoWasmBinary(db *sql.DB, wasmName, wasmRuntime, wasmCompressAlgorithm string, wasmBytes []byte, hash string) {
	preparedStmt, err := db.Prepare(`insert ignore into mysql.wasm_binary(name,wasmRuntime,data,compress_algorithm,hash_before_compress) values (?,?,?,?,?)`)
	if err != nil {
		panic(err.Error())
	}
	defer preparedStmt.Close()
	r, err := preparedStmt.Exec(wasmName, wasmRuntime, wasmBytes, wasmCompressAlgorithm, hash)
	if err != nil {
		panic(err.Error())
	}
	if affected, _ := r.RowsAffected(); affected != 1 {
		panic("insert failed")
	}
}

func generateMysqlDsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDb)
	return dsn
}
