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
)

// Environment Struct
type env struct {
	nameOs       string
	support      bool
	numberOfCpus int
	err          error
}

//  This struct  is for each one of all cpus
// referance : http://www.linuxhowtos.org/System/procstat.htm
type sampleCpuStat struct {

	// processes executing is user mode
	//such as Firefox, Mplayer...
	user uint64

	nice uint64

	// processes executing is system mode
	//such as kernel processes
	system uint64

	//idle: twiddling thumbs
	idle uint64

	//iowait: waiting for I/O to complete
	iowait uint64

	//irq: servicing interrupts
	irq uint64

	//softirq: servicing softirqs
	softirq uint64

	// sum of total all values
	sumOfall uint64

	// sum of user, nice  and system
	sumOfUserNiceSystem uint64
}

type snapShotsCPU struct {
	cpu sampleCpuStat
}

// System Status struct is readable for human
type SystemStatus struct {
	// all cpu usage
	CpuUsage float32

	MemInfo
}

// The package's values
var (
	sampleTimeOfRange   float64 = 300 // it  will convert to type of Milisecond
	pathProcStatOnLinux string  = "/proc/stat"
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

// to set the time of range for the sample of Cpu Stat.
// value type  is millisecond. For 1(one) second
// 1000 millisecond
func SetTimeOfRangeForCpuStat(t int) {

	sampleTimeOfRange = float64(t)
}
