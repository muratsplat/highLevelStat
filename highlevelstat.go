// Copyright (C) 2014  Murat ÖDÜNÇ
// murat.asya@gmail.com, http://muratodunc.wordpress.com
// See LICENSES.md file to know details the license

// The package simply generates the information of system status such as
// the percent of cpu(s) usage, mem, network, disk and something like that..
//
// The package is experimental!!!
//
// Examle use:
//
//	package main
//
//	import (
//		"fmt"
//		"github.com/muratsplat/highLevelStat"
//	)
//
//	func main() {
//
//	go func() {
//
//		for {
//
//			var test highlevelstat.SystemStatus
//			// getting information the structer of memory
//			var memInfo *highlevelstat.MemInfo = highlevelstat.GetMemInfo()
//
//			fmt.Printf("Cpu(s): %.f%% UsedMem: %.f%%\n", test.GetCpuUsage().CpuUsage, memInfo.GetUsedMemForHuman())
//
//		}
//
//	}()
//
//	var in string
//
//	fmt.Scanln(&in)
//}
package highlevelstat

import (
	"runtime"
	"strconv"
	"time"
)

// Environment Struct
type env struct {
	nameOs       string
	support      bool
	numberOfCpus int
	err          error
}

// System Status struct is readable for human
type SystemStatus struct {
	// all cpu usage
	CpuUsage float32

	MemInfo
}

// The package's values
var (
	sampleTimeOfRange   time.Duration = time.Duration(1000 * time.Millisecond)
	pathProcStatOnLinux string        = "/proc/stat"
)

// to detect environment. Unit now only Gnu/linux Os is supported.
// it can be added other os such as MacOSX, Free-Open BSD Unix maybe MS Windows.
func (env *env) detectEnv() *env {

	env.nameOs = runtime.GOOS

	switch true {
	case "linux" == env.nameOs:
		env.support = true
	default:
		env.support = false
	}

	env.numberOfCpus = runtime.NumCPU()

	return env

}

// to ckeck that the os is supperted
func (env *env) IsSupported() (support bool) {

	return env.detectEnv().support
}

// to convert string struct to unit64 struct
func convertStringToUint64(s string) uint64 {

	number, err := strconv.ParseUint(s, 0, 64)

	if err != nil {

		return uint64(0)

	}

	return number
}
