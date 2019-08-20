package assert

import "testing"

func TestEqual(t *testing.T) {
	cases := []struct {
		IsEqual  bool
		Expected interface{}
		Val      interface{}
	}{
		{IsEqual: true, Expected: "same string", Val: "same string"},
		{IsEqual: false, Expected: "different string1", Val: "different string2"},
		{IsEqual: true, Expected: 1, Val: 1},
		{IsEqual: false, Expected: 1, Val: 2},
		{IsEqual: true, Expected: true, Val: true},
		{IsEqual: false, Expected: true, Val: false},
		{IsEqual: true, Expected: 1.2, Val: 1.2},
		{IsEqual: false, Expected: 1.2, Val: 1.3},

		{IsEqual: false, Expected: "different type", Val: 16},
		{IsEqual: false, Expected: "different type", Val: true},
		{IsEqual: false, Expected: false, Val: 16},
	}

	for _, _case := range cases {
		testEqual := &testing.T{}
		Equal(testEqual, _case.Expected, _case.Val)
		if testEqual.Failed() != !_case.IsEqual {
			t.Errorf("[TestEqual]%+v", _case)
			t.Fail()
		}
		testEqualf := &testing.T{}
		Equalf(testEqualf, _case.Expected, _case.Val, "format")
		if testEqualf.Failed() != !_case.IsEqual {
			t.Errorf("[TestEqualf]%+v", _case)
			t.Fail()
		}
	}
}
