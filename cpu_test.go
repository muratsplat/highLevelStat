// Copyright (C) 2014  Murat ÖDÜNÇ
// murat.asya@gmail.com, http://muratodunc.wordpress.com
// See LICENSES.md file to know details the license

// Tests for CPU(s) State
package highlevelstat

import (
	"testing"
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

	var sS SystemStatus

	if sS.GetCpuUsage().CpuUsage != float32(0) {

		t.Error("Expected 0, got ", sS.GetCpuUsage().CpuUsage)

	}

}
