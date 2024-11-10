package dialect

import (
	"reflect"
)

var dialectMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect ...
func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

// GetDialect ...
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
