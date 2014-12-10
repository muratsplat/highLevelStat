// +build linux

// tests for the statics of memory on linux
package highlevelstat

import "testing"

func TestpickupKeyAndValueInmemInfo(t *testing.T) {

	var tT keyAndValueInMemInfo

	if tT.pickupKeyAndValueInmemInfo("MemTotal:        3759720 kB").key != "MemTotal:" {

		t.Error("Expected `MemTotal:`, got ", tT.pickupKeyAndValueInmemInfo("MemTotal:        3759720 kB").key)
	}

	if tT.pickupKeyAndValueInmemInfo("MemTotal:        999 kB").value == 3759720 {

		t.Error("Expected the value  should  be `999`, got ", 3759720)
	}

}

// a helper for changing the proc file path
func setProcMemInfoPath() {

	// changing the path of file for testing
	procMemInfoPath = "./testFiles/proc_meminfo1"
}

func TestTakeSnapShotOnMemInfo(t *testing.T) {

	var tT memRaw

	setProcMemInfoPath()

	if tT.takeSnapShot().MemTotal == 0 {

		t.Error("Expected the value is not zero(0), got : ", 0)
	}

	if tT.takeSnapShot().Cached == 0 {

		t.Error("Expected the value is not zero(0), got : ", 0)
	}

	if tT.takeSnapShot().Buffers == 0 {

		t.Error("Expected the value is not zero(0), got : ", 0)
	}

	if tT.takeSnapShot().MemFree == 0 {

		t.Error("Expected the value is not zero(0), got : ", 0)
	}
}

func TestCalculateAllMemInfo(t *testing.T) {

	var tT memRaw

	tT.MemTotal = 100

	tT.MemFree = 50

	v1 := calculateMemInfo(tT)

	if v1.PercentOfUsedMem != 50 {

		t.Error("Expected 50.00, got ", v1.PercentOfUsedMem)

	}

	tT.MemTotal = 1000

	tT.Buffers = 5

	v2 := calculateMemInfo(tT)

	if v2.PercentOfBuffersedMem != 0.5 {

		t.Error("Expected 0.5, got ", v2.PercentOfBuffersedMem)

	}

	tT.MemTotal = 1000

	tT.Cached = 5

	v3 := calculateMemInfo(tT)

	if v3.PercentOfCachedMem != 0.5 {

		t.Error("Expected 0.5, got ", v3.PercentOfBuffersedMem)

	}

	tT.MemTotal = 1000

	tT.Buffers = 5

	v4 := calculateMemInfo(tT)

	if v4.PercentOfBuffersedMem != 0.5 {

		t.Error("Expected 0.5, got ", v4.PercentOfBuffersedMem)

	}

}

func TestGetMemInfo(t *testing.T) {

	info := GetMemInfo()

	if info.PercentOfUsedMem == 0 {

		t.Error("Expected the value is not 0, got", info.PercentOfUsedMem)

	}

	if info.PercentOfUsedMemForHuman == 0 {

		t.Error("Expected the value is not 0, got", info.PercentOfUsedMemForHuman)
	}

	if info.PercentOfCachedMem == 0 {

		t.Error("Expected the value is not 0, got", info.PercentOfCachedMem)
	}

	if info.PercentOfBuffersedMem == 0 {

		t.Error("Expected the value is not 0, got", info.PercentOfBuffersedMem)
	}
}

func TestGetUsedMemForHuman(t *testing.T) {

	setProcMemInfoPath()

	info := GetMemInfo()

	if info.GetUsedMemForHuman() == 0.0 {

		t.Error("Expected the value is not 0, got ", info.GetUsedMemForHuman())

	}

	t.Log("percemt of used mem without caches and buffers is ", info.GetUsedMemForHuman())

}
