package assert

import "testing"

var (
	cases = []struct {
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

		// array and slice
		{IsEqual: true, Expected: []string{"1", "2", "3"}, Val: []string{"1", "2", "3"}},
		{IsEqual: false, Expected: []string{"1", "2"}, Val: []string{"1", "2", "3"}},
		{IsEqual: true, Expected: nil, Val: nil},
		{IsEqual: false, Expected: nil, Val: []string{"1", "2", "3"}},
		{IsEqual: false, Expected: []string{"2", "1"}, Val: []string{"1", "2"}},
		{IsEqual: false, Expected: []int{1}, Val: []int64{1}},

		// map
		{IsEqual: true, Expected: map[string]string{"1":"1", "2":"2"},
			Val: map[string]string{"1":"1", "2":"2"}},
		{IsEqual: false, Expected: map[string]string{"1":"1", "2":"2"},
			Val: map[string]string{"1":"2", "2":"2"}}, // val not equal
		{IsEqual: false, Expected: map[string]string{"1":"1", "2":"2"},
			Val: map[string]string{"1":"1", "3":"2"}}, // key not equal
		{IsEqual: false, Expected: map[string]string{"1":"1", "2":"2"},
			Val: map[string]string{"1":"1"}}, // len not equal
		{IsEqual: false, Expected: nil, Val: map[string]string{"1": "1"}},
		{IsEqual: false, Expected: map[string]int{"1": 1},
			Val: map[string]int64{"1": 1}}, // value type not equal

		{IsEqual: true, Expected: struct {
			field1 string
		}{"1"}, Val: struct {
			field1 string
		}{"1"}},
		{IsEqual: false, Expected: struct {
			field1 string
		}{"1"}, Val: struct {
			field1 string
		}{"2"}},
		{IsEqual: false, Expected: struct {
			field1 string
		}{"1"}, Val: 1},
	}
)
func TestEqual(t *testing.T) {
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

func TestNotEqual(t *testing.T) {
	for _, _case := range cases {
		testNotEqual := &testing.T{}
		NotEqual(testNotEqual, _case.Expected, _case.Val)
		if testNotEqual.Failed() != _case.IsEqual {
			t.Errorf("[TestNotEqual]%+v", _case)
			t.Fail()
		}
		testNotEqualf := &testing.T{}
		NotEqualf(testNotEqualf, _case.Expected, _case.Val, "format")
		if testNotEqualf.Failed() != _case.IsEqual {
			t.Errorf("[TestNotEqualf]%+v", _case)
			t.Fail()
		}
	}
}
