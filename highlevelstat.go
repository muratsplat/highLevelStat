// The package simply generates the information of system status such as
// the percent of cpu(s) usage, mem, network, disk and something like that..
//
// The package is experimental!!!

// Copyright (C) 2014  Murat ÖDÜNÇ
// See LICENSES.md file to know details the license

package highlavelstat

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strconv"
	str "strings"
	"time"
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

// This struct will include the sample of all CPU(s)
type sampleCPUS struct {
	allCpu []sampleCpuStat

	// sum of total all values
	sumOfallCpu uint64

	// sum of user, nice  and system
	sumOfUserNiceSystemAllCpu uint64
}

type snapShotsCPU struct {
	cpus sampleCPUS
}

// System Status struct is readable for human
type SystemStatus struct {
	// all cpu usage
	CpuUsage float32
}

// The package's values
var (
	sampleRangeOfTime   float64 = 300 // it is converted Type of Milisecond
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

	if err == nil {

		return number

	}

	return number
}

// to take snapshot that state sample of CPU(s)
// between in first time point and second time point.
func (sample sampleCPUS) takeSnapShot() sampleCPUS {

	var e env

	if e.IsSupported() == false {

		log.Fatalln("Your Os is not supported!")

		os.Exit(1)
	}

	file, err := os.Open(pathProcStatOnLinux)
	if err != nil {

		log.Fatalln("It looks that the file was not existed: ", pathProcStatOnLinux)

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	cpuStat := make([]sampleCpuStat, e.numberOfCpus)

	for i := 0; scanner.Scan(); i++ {

		// user: normal processes executing in user mode
		cpuStat[i].user = convertStringToUint64(str.Fields(scanner.Text())[1])

		// nice: niced processes executing in user mode
		cpuStat[i].nice = convertStringToUint64(str.Fields(scanner.Text())[2])

		// system: processes executing in kernel mode
		cpuStat[i].system = convertStringToUint64(str.Fields(scanner.Text())[3])

		// idle: twiddling thumbs
		cpuStat[i].idle = convertStringToUint64(str.Fields(scanner.Text())[4])

		// iowait: waiting for I/O to complete
		cpuStat[i].iowait = convertStringToUint64(str.Fields(scanner.Text())[5])

		//irq: servicing
		cpuStat[i].irq = convertStringToUint64(str.Fields(scanner.Text())[6])

		//softirq: servicing softirqs
		cpuStat[i].softirq = convertStringToUint64(str.Fields(scanner.Text())[7])

		// calculate things that have to
		cpuStat[i].calculateToAll()

		if i == e.numberOfCpus-1 {

			break // we need only cpu data rather than others
		}

	}

	sample.allCpu = cpuStat

	return sample

}

// To calculate the plus of single cpu values
func (cpu *sampleCpuStat) calculateToAll() {

	cpu.sumOfall = cpu.idle + cpu.iowait + cpu.irq + cpu.nice + cpu.softirq + cpu.system + cpu.user

	cpu.sumOfUserNiceSystem = cpu.nice + cpu.system + cpu.user

}

// To calculate the plus of all cpus values
func (s *sampleCPUS) sumTotal() {

	for i := 0; i < len(s.allCpu); {

		s.sumOfallCpu = s.allCpu[i].sumOfall
		s.sumOfUserNiceSystemAllCpu = s.allCpu[i].sumOfUserNiceSystem
		i++
	}

}

// To get the percent of CPU(s) usage
func (s *SystemStatus) GetCpuUsage() *SystemStatus {

	var snaps snapShotsCPU

	snapshots := snaps.getSnapShots()

	workOverPeriod := float32(snapshots[1].cpus.sumOfUserNiceSystemAllCpu - snapshots[0].cpus.sumOfUserNiceSystemAllCpu)
	totalOverPeriod := float32(snapshots[1].cpus.sumOfallCpu - snapshots[0].cpus.sumOfallCpu)

	s.CpuUsage = float32((workOverPeriod / totalOverPeriod) * 100.00)

	return s

}

func (s *snapShotsCPU) getSnapShots() []snapShotsCPU {

	var samples sampleCPUS

	snapShots := make([]snapShotsCPU, 2)

	for i := 0; i < len(snapShots); {

		snapShots[i].cpus = samples.takeSnapShot()
		snapShots[i].cpus.sumTotal()

		i++

		time.Sleep(time.Millisecond * time.Duration(sampleRangeOfTime))

	}

	return snapShots
}

// Examle use

//func main() {

//	go func() {

//		var status SystemStatus

//		for {

//			fmt.Printf("Cpu(s): %.2f%%\n", status.getCpuUsage().CpuUsage)
//		}

//	}()

//	var input string

//	fmt.Scanln(&input)

//}
