package assert

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected interface{}, val interface{}) {
	equalf(t, expected, val, "[assert.Equal]expected: %+v, get: %+v", expected, val)
}

func Equalf(t *testing.T, expected interface{}, val interface{}, format string, args... interface{}) {
	equalf(t, expected, val, format, args...)
}

func equalf(t *testing.T, expected interface{}, val interface{}, format string, args... interface{}) {
	t.Helper()
	// check type
	typeExpected := reflect.TypeOf(expected)
	typeVal := reflect.TypeOf(val)
	if typeExpected != typeVal {
		t.Errorf("assert type not equal! expected type: %s, get type: %s", typeExpected, typeVal)
		return
	}
	// check equal
	if expected != val {
		t.Errorf(format, args...)
		return
	}
}
