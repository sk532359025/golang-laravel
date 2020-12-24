// Code generated by "stringer -type=Kind"; DO NOT EDIT.

package metric

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ValueRecorderKind-0]
	_ = x[ValueObserverKind-1]
	_ = x[CounterKind-2]
	_ = x[UpDownCounterKind-3]
	_ = x[SumObserverKind-4]
	_ = x[UpDownSumObserverKind-5]
}

const _Kind_name = "ValueRecorderKindValueObserverKindCounterKindUpDownCounterKindSumObserverKindUpDownSumObserverKind"

var _Kind_index = [...]uint8{0, 17, 34, 45, 62, 77, 98}

func (i Kind) String() string {
	if i < 0 || i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
