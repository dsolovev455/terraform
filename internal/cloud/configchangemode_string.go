// Code generated by "stringer -type ConfigChangeMode"; DO NOT EDIT.

package cloud

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ConfigMigrationIn-8600]
	_ = x[ConfigMigrationOut-8598]
	_ = x[ConfigChangeInPlace-8635]
	_ = x[ConfigChangeIrrelevant-129335]
}

const (
	_ConfigChangeMode_name_0 = "ConfigMigrationOut"
	_ConfigChangeMode_name_1 = "ConfigMigrationIn"
	_ConfigChangeMode_name_2 = "ConfigChangeInPlace"
	_ConfigChangeMode_name_3 = "ConfigChangeIrrelevant"
)

func (i ConfigChangeMode) String() string {
	switch {
	case i == 8598:
		return _ConfigChangeMode_name_0
	case i == 8600:
		return _ConfigChangeMode_name_1
	case i == 8635:
		return _ConfigChangeMode_name_2
	case i == 129335:
		return _ConfigChangeMode_name_3
	default:
		return "ConfigChangeMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}