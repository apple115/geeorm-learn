package clause

import (
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/vars"
	"golang.org/x/text/unicode/rangetable"
)

type generator func(value ...interface{}) (string, []interface{})

var generators map[Type]generator

// init ...
func init() {
	generators = make(map[Type]generator)
	generators[INSERT] = _insert
	generators[VALUES] = _values
	generators[SELECT] = _select
	generators[LIMIT] = _limit
	generators[WHERE] = _where
	generators[ORDERBY] = _orderBy
	generators[UPDATE] =_update
	generators[DELECT] =_delete
	generators[COUNT] = _count
}

// genBindVars ...
func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ",")
}

// _insert INSET INOT $tableName ($fields)
func _insert(values ...interface{}) (string, []interface{}) {
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}

// _values ...
func _values(values ...interface{}) (string, []interface{}) {
	var bindStr string
	var sql strings.Builder
	var vars []interface{}
	sql.WriteString("VALUES")
	for i, value := range values {
		v := value.([]interface{})
		if bindStr == "" {
			bindStr = genBindVars(len(v))
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if i+1 != len(values) {
			sql.WriteString(",")
		}
		vars = append(vars, v...)
	}
	return sql.String(), vars
}

// _select ...
func _select(value ...interface{}) (string, []interface{}) {
	tabName := value[0]
	fields := strings.Join(value[1].([]string), ",")
	return fmt.Sprintf("SELECT %v FROM %s", fields, tabName), []interface{}{}
}

// _limit ...
func _limit(values ...interface{}) (string, []interface{}) {
	return "LIMIT ?", values
}

// _where ...
func _where(values ...interface{}) (string, []interface{}) {
	desc, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", desc), vars
}

// _orderBy ...
func _orderBy(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}

func _update(values...interface{})(string,[]interface{}){
	tableName := values[0]
	m := values[i].(map[string]interface{})
	var keys []string
	var vars []interface{}
	for k,v := range m{
		keys = append(keys, k+" = ? ")
		vars = append(vars, v)
	}
	return fmt.Sprint("UPDATE %s SET %s",tableName,strings.Join(keys,", ")),vars
}

func _delete(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("DELETE FROM %s", values[0]), []interface{}{}
}

func _count(values ...interface{}) (string, []interface{}) {
	return _select(values[0], []string{"count(*)"})
}
