// Copyright (C) 2014  Murat ÖDÜNÇ
// murat.asya@gmail.com, http://muratodunc.wordpress.com
// See LICENSES.md file to know details the license

// CPU(s) State
package highlevelstat

import (
	"bufio"
	"log"
	"os"
	str "strings"
	"time"
)

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

	// fixes issue that the cases of workOverPeriod  equals to
	// totalOverPeriod
	if workOverPeriod == totalOverPeriod {

		s.CpuUsage = float32(0)

		return s
	}

	s.CpuUsage = float32((workOverPeriod / totalOverPeriod) * 100.00)

	return s

}

// To get two snaphots of cpu(s) state
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
