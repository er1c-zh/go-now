package assert

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected interface{}, val interface{}) {
	t.Helper()
	equalf(t, expected, val, "[assert.Equal]expected: %+v, get: %+v", expected, val)
}

func Equalf(t *testing.T, expected interface{}, val interface{}, format string, args ...interface{}) {
	t.Helper()
	equalf(t, expected, val, format, args...)
}

func NotEqual(t *testing.T, expected interface{}, val interface{}) {
	t.Helper()
	notEqualf(t, expected, val, "[assert.NotEqual] expect %+v != %+v, but equal", val, expected)
}

func NotEqualf(t *testing.T, expected interface{}, val interface{}, format string, args ...interface{}) {
	t.Helper()
	notEqualf(t, expected, val, format, args...)
}

func equalf(t *testing.T, expected interface{}, val interface{}, format string, args ...interface{}) {
	t.Helper()
	// check type
	if isEqual(expected, val) == false {
		t.Errorf(format, args...)
		return
	}
}

func notEqualf(t *testing.T, expected interface{}, val interface{}, format string, args ...interface{}) {
	t.Helper()
	// check equal
	if isEqual(expected, val) == true {
		t.Errorf(format, args...)
		return
	}
}

/////////////////////////////////////////
// utils                               //
/////////////////////////////////////////

/*
compare any type args
*/
func isEqual(v1 interface{}, v2 interface{}) bool {
	return reflect.DeepEqual(v1, v2)
}
