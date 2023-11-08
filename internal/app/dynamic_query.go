package app

import (
	"fmt"
	"strings"
)

type DynamicQuery struct {
	query                   *strings.Builder
	currentPlaceholderIndex int
	args                    []interface{}
}

func NewDynamicQuery(query string) *DynamicQuery {
	q := &strings.Builder{}
	q.WriteString(query)
	q.WriteString(" WHERE 1=1")

	dq := &DynamicQuery{query: q, args: make([]interface{}, 0)}

	return dq
}

func (dq *DynamicQuery) Where(column string, operator string, value interface{}) *DynamicQuery {
	placeholder := dq.AddArgument(value)
	dq.query.WriteString(fmt.Sprintf(" AND %s %s %s", column, operator, placeholder))
	return dq
}

func (dq *DynamicQuery) Contains(column string, values []string) *DynamicQuery {
	placeholders := make([]string, len(values))
	for i, value := range values {
		placeholder := dq.AddArgument(value)
		placeholders[i] = placeholder
	}

	dq.query.WriteString(fmt.Sprintf(" AND %s IN (%s)", column, strings.Join(placeholders, ",")))
	return dq
}

func (dq *DynamicQuery) Order(value interface{}) *DynamicQuery {
	dq.query.WriteString(fmt.Sprintf(" ORDER BY %s", value))
	return dq
}

func (dq *DynamicQuery) Limit(value int) *DynamicQuery {
	dq.query.WriteString(fmt.Sprintf(" LIMIT %d", value))
	return dq
}

func (dq *DynamicQuery) Offset(value int) *DynamicQuery {
	dq.query.WriteString(fmt.Sprintf(" OFFSET %d", value))
	return dq
}

func (dq *DynamicQuery) AddArgument(value interface{}) string {
	dq.currentPlaceholderIndex++
	dq.args = append(dq.args, value)
	return fmt.Sprintf("$%d", dq.currentPlaceholderIndex)
}
