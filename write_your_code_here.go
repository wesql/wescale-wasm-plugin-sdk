package wescale_wasm_plugin_template

import (
	"errors"
	"fmt"
	"strconv"
	"wescale-wasm-plugin-template/internal"
	"wescale-wasm-plugin-template/internal/host_functions"
)

func RunBeforeExecution() error {
	// TODO: Write your code here
	hostfunction.GlobalLock()
	var moduleCount int
	countBytes, err := hostfunction.GetModuleValueByKey("moduleCount")
	if errors.Is(err, internal.ErrorStatusNotFound) {
		moduleCount = 0
		hostfunction.SetModuleValueByKey("moduleCount", []byte(strconv.Itoa(moduleCount)))
	}

	countBytes, _ = hostfunction.GetModuleValueByKey("moduleCount")
	moduleCount, _ = strconv.Atoi(string(countBytes))
	moduleCount++

	if moduleCount%2 == 0 {
		hostfunction.SetHostQuery("select * from guest.setquerytest;")
	} else {
		str, _ := hostfunction.GetHostQuery()
		hostfunction.SetHostQuery(str)
	}

	runtimeType, _ := hostfunction.GetRuntimeType()
	version, _ := hostfunction.GetAbiVersion()
	if runtimeType != "" {
		if version != "" {

		}
	}
	hostfunction.InfoLog(fmt.Sprintf("wasm guest runtime type:%v version:%v", runtimeType, version))

	hostfunction.SetModuleValueByKey("moduleCount", []byte(strconv.Itoa(moduleCount)))
	hostfunction.GlobalUnlock()
	return errors.New("error test foo bar")
}

func RunAfterExecution() error {
	// TODO: Write your code here
	//qr, err := hostfunction.GetQueryResult()
	//if err != nil {
	//	return err
	//}
	//qr.Rows = qr.Rows[:len(qr.Rows)-1]
	//qr.RowsAffected = qr.RowsAffected - 1
	//return hostfunction.SetQueryResult(qr)
	return nil
}
