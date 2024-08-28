package main

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/dsnet/compress/bzip2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/pflag"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
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
	wasmFile              = ""
	wasmRuntime           = "wazero"
	wasmCompressAlgorithm = "bzip2"
)

// filter config
var createFilter = false
var (
	filterName                     = ""
	filterDesc                     = ""
	filterPriority                 = "1000"
	filterStatus                   = "ACTIVE"
	filterPlans                    = "Select,Insert,Update,Delete"
	filterFullyQualifiedTableNames = ""
	filterQueryRegex               = ``
	filterQueryTemplate            = ""
	filterRequestIpRegex           = ``
	filterUserRegex                = ``
	filterLeadingCommentRegex      = ``
	filterTrailingCommentRegex     = ``
	filterBindVarConds             = ""
	filterAction                   = "wasm_plugin"
)

const createFilterTemplate = `create filter if not exists %s (
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

func init() {
	pflag.StringVar(&command, "command", command, "the command of the script. default is install")

	pflag.StringVar(&mysqlUser, "mysql_user", mysqlUser, "the user of mysql")
	pflag.StringVar(&mysqlPassword, "mysql_password", mysqlPassword, "the password of mysql")
	pflag.StringVar(&mysqlHost, "mysql_host", mysqlHost, "the host of mysql")
	pflag.IntVar(&mysqlPort, "mysql_port", mysqlPort, "the port of mysql")
	pflag.StringVar(&mysqlDb, "mysql_db", mysqlDb, "the db of mysql")

	pflag.StringVar(&wasmFile, "wasm_file", wasmFile, "the wasmFile of wasm file")
	pflag.StringVar(&wasmRuntime, "wasm_runtime", wasmRuntime, "the wasm_runtime of wasm")
	pflag.StringVar(&wasmCompressAlgorithm, "wasm_compress_algorithm", wasmCompressAlgorithm, "the wasm_compress_algorithm of wasm")

	pflag.BoolVar(&createFilter, "skip_filter", false, "skip filter creation")
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

	filterName = generateFilterName(filterName, getWasmFileName())
}

func main() {
	switch command {
	case "install":
		InstallWasm()
		if createFilter {
			InstallFilter()
		}
	case "uninstall":
		UnInstall()
	}
}

func generateFilterSQL() string {
	return fmt.Sprintf(createFilterTemplate,
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
		fmt.Sprintf("wasm_binary_name=\"%v\"", getWasmFileName()),
	)
}

func getDB() *sql.DB {
	db, err := sql.Open("mysql", generateMysqlDsn())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connected")
	return db
}

func InstallWasm() {
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

	db := getDB()
	defer db.Close()

	insertIntoWasmBinary(db, getWasmFileName(), wasmRuntime, wasmCompressAlgorithm, wasmBytes, hash)
}

func InstallFilter() {
	db := getDB()
	defer db.Close()

	query := generateFilterSQL()
	fmt.Println(query)
	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func InstallCdc() {
	db := getDB()
	defer db.Close()

	//todo
	query := "insert into mysql.cdc_consumer ......"
	fmt.Println(query)
	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func UnInstall() {
	db, err := sql.Open("mysql", generateMysqlDsn())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dropFilterSQL := fmt.Sprintf("drop filter %s", filterName)
	deleteWasmSQL := fmt.Sprintf("delete from mysql.wasm_binary where name='%s'", getWasmFileNameFromFilter(db))

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
		} else {
			fmt.Printf("filter %s deleted\n", filterName)
		}
	}

	{
		r, err := db.Exec(deleteWasmSQL)
		if err != nil {
			panic(err.Error())
		}
		if affected, _ := r.RowsAffected(); affected == 0 {
			fmt.Printf("wasm %s not found\n", getWasmFileName())
		} else {
			fmt.Printf("wasm %s deleted\n", getWasmFileName())
		}
	}
}

func insertIntoWasmBinary(db *sql.DB, wasmName, wasmRuntime, wasmCompressAlgorithm string, wasmBytes []byte, hash string) {
	preparedStmt, err := db.Prepare(`insert ignore into mysql.wasm_binary(name,runtime,data,compress_algorithm,hash_before_compress) values (?,?,?,?,?)`)
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

func calcMd5String32(data []byte) string {
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}

func getWasmFilePath() string {
	// Get the directory path from wasmFile
	dirPath := filepath.Dir(wasmFile)
	return dirPath
}

func getWasmFileName() string {
	// Get the file name from wasmFile
	fileName := filepath.Base(wasmFile)
	return fileName
}

func getWasmFileNameFromFilter(db *sql.DB) string {
	type Row struct {
		Id                       int
		CreateTimestamp          string
		UpdateTimestamp          string
		Name                     string
		Description              string
		Priority                 int
		Status                   string
		Plans                    string
		FullyQualifiedTableNames string
		QueryRegex               string
		QueryTemplate            string
		RequestIpRegex           string
		UserRegex                string
		LeadingCommentRegex      string
		TrailingCommentRegex     string
		BindVarConds             string
		Action                   string
		ActionArgs               string
	}

	query := fmt.Sprintf("show filter %s", filterName)
	r, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer r.Close()

	var row Row
	if r.Next() {
		if err := r.Scan(&row.Id, &row.CreateTimestamp, &row.UpdateTimestamp, &row.Name, &row.Description,
			&row.Priority, &row.Status, &row.Plans, &row.FullyQualifiedTableNames, &row.QueryRegex, &row.QueryTemplate,
			&row.RequestIpRegex, &row.UserRegex, &row.LeadingCommentRegex, &row.TrailingCommentRegex, &row.BindVarConds,
			&row.Action, &row.ActionArgs); err != nil {
			panic(err.Error())
		}
	} else {
		panic("filter %s not found when uninstalling wasm plugin\n")
	}

	return getWasmFileNameFromActionArgs(row.ActionArgs)
}

func getWasmFileNameFromActionArgs(actionArgs string) string {
	type WasmPluginActionArgs struct {
		WasmBinaryName string `toml:"wasm_binary_name"`
	}

	var result strings.Builder
	insideQuotes := false

	for _, char := range actionArgs {
		if char == '"' {
			insideQuotes = !insideQuotes
		}
		if char == ';' && !insideQuotes {
			result.WriteString("\n")
		} else {
			result.WriteRune(char)
		}
	}
	userInputTOML := result.String()

	w := &WasmPluginActionArgs{}

	err := toml.Unmarshal([]byte(userInputTOML), w)
	if err != nil {
		panic(fmt.Sprintf("error when parsing wasm plugin action args: %v", err))
	}
	if w.WasmBinaryName == "" {
		panic("wasm binary name is empty")
	}

	wasmFile = w.WasmBinaryName

	return w.WasmBinaryName
}

func generateFilterName(filterName, wasmFileName string) string {
	if wasmFileName == "" {
		return filterName
	}
	if filterName == "" {
		filterName = strings.ReplaceAll(wasmFileName, ".", "_") + "_filter"
	}
	return filterName
}
