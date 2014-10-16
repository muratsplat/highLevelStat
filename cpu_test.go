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

	var testSample sampleCPUS

	var tsample sampleCPUS = testSample.takeSnapShot()

	if len(tsample.allCpu) != 4 {

		t.Error("Expected 4, got", len(tsample.allCpu))
	}

	if tsample.allCpu[0].user == uint64(0) {

		t.Error("Expected the value is not 0, got", tsample.sumOfUserNiceSystemAllCpu)
	}

	for _, v := range tsample.allCpu {

		if v.user == uint64(0) {

			t.Error("Expected the value should not 0, got", v.user)
		}

		if v.system == uint64(0) {

			t.Error("Expected the value should not 0, got", v.system)
		}

		if v.idle == uint64(0) {

			t.Error("Expected the value should not 0, got", v.idle)
		}

		if v.nice == uint64(0) {

			t.Error("Expected the value should not 0, got", v.nice)
		}

		if v.iowait == uint64(0) {

			t.Error("Expected the value should not 0, got", v.iowait)
		}

		if v.sumOfall == uint64(0) {

			t.Error("Expected the value should not 0, got", v.sumOfall)

		}

		if v.sumOfUserNiceSystem == uint64(0) {

			t.Error("Expected the value should not 0, got", v.sumOfUserNiceSystem)

		}

	}
}

func TestGetSnapShots(t *testing.T) {

	var tS snapShotsCPU

	// let's use our stat file
	pathProcStatOnLinux = "./testFiles/proc_stat1"

	snaps := tS.getSnapShots()

	if len(snaps) != 2 {

		t.Error("Expected 2, got ", len(snaps))
	}

	for _, v := range snaps {

		if len(v.cpus.allCpu) != 4 || v.cpus.sumOfallCpu+v.cpus.sumOfUserNiceSystemAllCpu == uint64(0) {

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
