// Code generated by "stringer -type=Type"; DO NOT EDIT.

package label

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[INVALID-0]
	_ = x[BOOL-1]
	_ = x[INT32-2]
	_ = x[INT64-3]
	_ = x[UINT32-4]
	_ = x[UINT64-5]
	_ = x[FLOAT32-6]
	_ = x[FLOAT64-7]
	_ = x[STRING-8]
	_ = x[ARRAY-9]
}

const _Type_name = "INVALIDBOOLINT32INT64UINT32UINT64FLOAT32FLOAT64STRINGARRAY"

var _Type_index = [...]uint8{0, 7, 11, 16, 21, 27, 33, 40, 47, 53, 58}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
