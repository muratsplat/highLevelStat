// +build linux

package highlevelstat

import (
	"testing"
	"time"
)

func TestTakeSnapShot(t *testing.T) {
	// let's use our stat file
	pathProcStatOnLinux = "./testFiles/proc_stat1"

	var testSample snapShotsCPU

	var tsample snapShotsCPU = testSample.takeSnapShot()

	if tsample.cpu.user == uint64(0) {

		t.Error("Expected the value should not 0, got", tsample.cpu.user)
	}

	if tsample.cpu.system == uint64(0) {

		t.Error("Expected the value should not 0, got", tsample.cpu.system)
	}

	if tsample.cpu.idle == uint64(0) {

		t.Error("Expected the value should not 0, got", tsample.cpu.idle)
	}

	if tsample.cpu.nice == uint64(0) {

		t.Error("Expected the value should not 0, got", tsample.cpu.nice)
	}

	if tsample.cpu.iowait == uint64(0) {

		t.Error("Expected the value should not 0, got", tsample.cpu.iowait)
	}

	if tsample.cpu.sumOfall == uint64(0) {

		t.Error("Expected the value should not 0, got", tsample.cpu.sumOfall)

	}

	if tsample.cpu.sumOfUserNiceSystem == uint64(0) {

		t.Error("Expected the value should not 0, got", tsample.cpu.sumOfUserNiceSystem)

	}

}

func TestGetSnapShots(t *testing.T) {

	var tS snapShotsCPU
	// for quick test
	SetTimeOfRangeForCpuStat(time.Millisecond * 300)

	snaps := tS.getSnapShots()

	if len(snaps) != 2 {

		t.Error("Expected 2, got ", len(snaps))
	}

	for _, v := range snaps {

		if v.cpu.sumOfUserNiceSystem+v.cpu.sumOfall == uint64(0) {

			t.Error("It looks that samples is invalid!!!")

		}
	}

}

func TestGetCpuUsage(t *testing.T) {

	sS := NewCpuUsage()

	if sS.CpuUsage != float32(0) {

		t.Error("Expected 0, got ", sS.CpuUsage)

	}

}

func TestSetSampleTimeOfRange(t *testing.T) {

	SetTimeOfRangeForCpuStat(time.Millisecond * 200)

	if sampleTimeOfRange != time.Millisecond*200 {

		t.Error("Expected 600, got ", sampleTimeOfRange)

	}

}
