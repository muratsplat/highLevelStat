// highlevelstat_test
package highlevelstat

import (
	"testing"
)

func TestConvertStringToUint64(t *testing.T) {

	v1 := convertStringToUint64("")

	if v1 != 0 {

		t.Error("Expected 0, got ", v1)
	}

	v2 := convertStringToUint64("0")

	if v2 != 0 {

		t.Error("Expected 0, got ", v2)

	}

	v3 := convertStringToUint64("0,1")

	if v3 != 0 {

		t.Error("Expected 0, got ", v3)
	}

}

func TestIsSupported(t *testing.T) {

	var environment env

	if environment.IsSupported() == false {

		t.Error("Expected true, got", environment.support)
	}

}
