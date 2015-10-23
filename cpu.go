// +build linux

package highlevelstat

import (
	"bufio"
	"log"
	"os"
	str "strings"
	"time"
)

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

// to take snapshot that state CPU
// between in first time point and second time point.
func takeSnapShot() snapShotsCPU {

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

	if scanner == nil {

		log.Println("Retuned NewScanner Pointer is not valid")

	}

	var sample snapShotsCPU

	for i := 0; scanner.Scan(); i++ {

		// user: normal processes executing in user mode
		sample.cpu.user = convertStringToUint64(str.Fields(scanner.Text())[1])

		// nice: niced processes executing in user mode
		sample.cpu.nice = convertStringToUint64(str.Fields(scanner.Text())[2])

		// system: processes executing in kernel mode
		sample.cpu.system = convertStringToUint64(str.Fields(scanner.Text())[3])

		// idle: twiddling thumbs
		sample.cpu.idle = convertStringToUint64(str.Fields(scanner.Text())[4])

		// iowait: waiting for I/O to complete
		sample.cpu.iowait = convertStringToUint64(str.Fields(scanner.Text())[5])

		//irq: servicing
		sample.cpu.irq = convertStringToUint64(str.Fields(scanner.Text())[6])

		//softirq: servicing softirqs
		sample.cpu.softirq = convertStringToUint64(str.Fields(scanner.Text())[7])

		// calculate things that have to
		sample.cpu.calculateToAll()

		break // we need only cpu data rather than multicores

	}

	return sample

}

// To get two snaphots of cpu(s) state
func getSnapShotsOfCpu() []snapShotsCPU {

	snapShots := make([]snapShotsCPU, 2)

	for i := 0; i < len(snapShots); {

		snapShots[i] = takeSnapShot()

		i++

		time.Sleep(time.Millisecond * time.Duration(sampleTimeOfRange))

	}

	return snapShots
}

// To calculate the plus of single cpu values
func (cpu *sampleCpuStat) calculateToAll() {

	cpu.sumOfall = cpu.idle + cpu.iowait + cpu.irq + cpu.nice + cpu.softirq + cpu.system + cpu.user

	cpu.sumOfUserNiceSystem = cpu.nice + cpu.system + cpu.user

}

// To get the percent of CPU(s) usage on linux
func (s SystemStatus) GetCpuUsage() SystemStatus {

	snapshots := getSnapShotsOfCpu()

	workOverPeriod := float32(snapshots[1].cpu.sumOfUserNiceSystem - snapshots[0].cpu.sumOfUserNiceSystem)
	totalOverPeriod := float32(snapshots[1].cpu.sumOfall - snapshots[0].cpu.sumOfall)

	// fixes issue that the cases of workOverPeriod  equals to
	// totalOverPeriod
	if workOverPeriod == totalOverPeriod {

		s.CpuUsage = float32(0)

		return s
	}

	s.CpuUsage = float32((workOverPeriod / totalOverPeriod) * 100.00)

	return s

}

// to set the time of range for the sample of Cpu Stat.
// value type  is millisecond. For 1(one) second
// 1000 millisecond
func SetTimeOfRangeForCpuStat(t int64) {

	sampleTimeOfRange = t
}
