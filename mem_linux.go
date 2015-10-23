// +build linux

// the statics of memory on linux
package highlevelstat

import (
	"bufio"
	"log"
	"os"
	str "strings"
)

// memRaw data struct to create MemInfo struct
type memRaw struct {
	MemTotal uint64
	MemFree  uint64
	Buffers  uint64
	Cached   uint64
}

// readable struct for humans
type MemInfo struct {
	PercentOfUsedMem float32

	PercentOfUsedMemForHuman float32
	PercentOfCachedMem       float32
	PercentOfBuffersedMem    float32
}

// pick up struct for file that "/proc/meminfo"
type keyAndValueInMemInfo struct {
	key   string
	value uint64
}

// simple pick uper for key and value in "/proc/meminfo" file
func (ky *keyAndValueInMemInfo) pickupKeyAndValueInmemInfo(s string) *keyAndValueInMemInfo {

	ky.key = str.Fields(s)[0]

	ky.value = convertStringToUint64(str.Fields(s)[1])

	return ky

}

// the path of Gnu/linux kernel file
// this value is changeble for testing
var procMemInfoPath string = "/proc/meminfo"

// to take snapshot of the status of memory
func (m memRaw) takeSnapShot() memRaw {

	file, err := os.Open(procMemInfoPath)

	if err != nil {

		log.Fatalln(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	var ky keyAndValueInMemInfo

	for scanner.Scan() {

		switch ky.pickupKeyAndValueInmemInfo(scanner.Text()).key {

		case "MemTotal:":
			m.MemTotal = ky.value

		case "MemFree:":
			m.MemFree = ky.value

		case "Buffers:":
			m.Buffers = ky.value

		case "Cached:":
			m.Cached = ky.value

		default:
			break

		}

	}

	return m

}

// to calculate percents of used mem, readable used mem,
// cached mem and buffered mem
func calculateMemInfo(m memRaw) MemInfo {

	return MemInfo{

		PercentOfUsedMem: (float32(m.MemTotal-m.MemFree) / float32(m.MemTotal)) * 100.0,

		PercentOfUsedMemForHuman: (float32((m.MemTotal-m.MemFree)-(m.Cached+m.Buffers)) / float32(m.MemTotal)) * 100.0,

		PercentOfCachedMem: (float32(m.Cached) / float32(m.MemTotal)) * 100.0,

		PercentOfBuffersedMem: (float32(m.Buffers) / float32(m.MemTotal)) * 100.0,
	}

}

// to get proccesed MemInfo struct
func GetMemInfo() MemInfo {

	var mem memRaw

	raw := mem.takeSnapShot()

	memInfo := calculateMemInfo(raw)

	return memInfo
}

// to get used mem which is without caches ana buffers
// For mostly people it may be suitable
func (m MemInfo) GetUsedMemForHuman() float32 {

	return m.PercentOfUsedMemForHuman
}
