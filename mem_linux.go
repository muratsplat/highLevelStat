// +build linux

// the statics of memory on linux
package highlevelstat

import (
	"bufio"
	"log"
	"os"

	//"runtime"
	//"strconv"
	str "strings"
	//"time"
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

var procMemInfo string = "/proc/meminfo"

func (m memRaw) takeSnapShot() memRaw {

	file, err := os.Open(procMemInfo)

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

func calculateMemInfo(m memRaw) *MemInfo {

	return &MemInfo{

		PercentOfUsedMem: (float32(m.MemTotal-m.MemFree) / float32(m.MemTotal)) * 100.0,

		PercentOfUsedMemForHuman: (float32((m.MemTotal-m.MemFree)-(m.Cached+m.Buffers)) / float32(m.MemTotal)) * 100.0,

		PercentOfCachedMem: (float32(m.Cached) / float32(m.MemTotal)) * 100.0,

		PercentOfBuffersedMem: (float32(m.Buffers) / float32(m.MemTotal)) * 100.0,
	}

}

func (s SystemStatus) GetMemInfo() *MemInfo {

	var mem memRaw

	raw := mem.takeSnapShot()

	memInfo := calculateMemInfo(raw)

	return memInfo
}
