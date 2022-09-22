package repo

import (
	"fmt"
)

type PrefixWherePredicate uint8
const (
	PrefixAnd PrefixWherePredicate = iota + 1
	PrefixOr
)
var mapPrefixWherePredicate = map[PrefixWherePredicate]string{
	PrefixAnd: "AND",
	PrefixOr: "OR",
}
type Stmt interface {
	Parse() (query string, args []interface{})
}
type Operator uint8
const (
	Equal Operator = iota
	LessThan
	LessThanEqual
	GreaterThan
	GreaterThanEqual
	LikeString
)

var mapOperator = [...]string{"=", "<", "<=", ">", ">=", "LIKE"}

func (o *Operator) ToStmt() string {
	return mapOperator[*o]
}

type selectCol struct {
	col string
	tb string
	as string
}

func C(col, tb string) *selectCol {
	return &selectCol{
		col: col,
		tb: tb,
	}
}


func (s *selectCol) As(reName string) *selectCol{
	s.as = reName
	return s
}
type selectStmt struct {
	selectCols []*selectCol
}


func (s *selectStmt) Parse() (query string, args []interface{}) {
	q := fmt.Sprintf("SELECT ")
	if len(s.selectCols) == 0 {
		return "", nil
	}
	for i, sc := range s.selectCols {
		q += fmt.Sprintf("`%v`.`%v`", sc.tb, sc.col)
		if sc.as != "" {
			q += fmt.Sprintf(" AS %v", sc.as)
		}
		if i < len(s.selectCols) - 1 {
			q += ", "
		}
	}
	return q, nil
}

type predicateCol struct {
	col string
	tb string
	operator Operator
	args interface{}
	prefixPredicate PrefixWherePredicate
}
func P(col, tb string, operator Operator, args interface{}) *predicateCol {
	return &predicateCol{
		col:      col,
		tb:       tb,
		operator: operator,
		args:     args,
	}
}

type predicateStmt struct {
	predicateCols []*predicateCol
}



func (p *predicateStmt) Parse() (query string, args []interface{}) {
	query = fmt.Sprintf("WHERE ")
	if len(p.predicateCols) == 0 {
		return "", nil
	}
	for _, pc := range p.predicateCols {
		if pc.prefixPredicate >= PrefixAnd && pc.prefixPredicate <= PrefixOr {
			query += " " + mapPrefixWherePredicate[pc.prefixPredicate] + " "
			fmt.Println("have prefix And")
		}
		query += fmt.Sprintf("`%v`.`%v` %v ?", pc.tb, pc.col, pc.operator.ToStmt())
		args = append(args, pc.args)
	}
	return
}


func (q *QueryBuilder) Where(pCol *predicateCol, prefixAndOr ...PrefixWherePredicate) *QueryBuilder{
	if len(q.stmt) > 1 {
		if len(prefixAndOr) > 0 {
			pCol.prefixPredicate = prefixAndOr[0]
		}
		q.stmt[1].(*predicateStmt).predicateCols = append(q.stmt[1].(*predicateStmt).predicateCols, pCol)
	}
	return q
}

type QueryBuilder struct {
	stmt []Stmt
	args []interface{}
}

func (q *QueryBuilder) Select(selectCol *selectCol) *QueryBuilder {
	if len(q.stmt) > 0 {
		q.stmt[0].(*selectStmt).selectCols = append(q.stmt[0].(*selectStmt).selectCols, selectCol)
	}
	return q
}

func (q *QueryBuilder) ToQuery() (string, []interface{}) {
	v := ""
	for i, stmt := range q.stmt {
		query, args := stmt.Parse()
		if args != nil {
			q.args = append(q.args, args...)
		}
		if i < len(q.stmt) - 1 {
			query += " "
		}
		v += query
	}
	return v, q.args
}
func NewQueryBuilder() *QueryBuilder {
	q := &QueryBuilder{
		stmt: make([]Stmt, 0),
	}
	q.stmt = append(q.stmt, &selectStmt{selectCols: []*selectCol{}})
	q.stmt = append(q.stmt, &predicateStmt{predicateCols: []*predicateCol{}})
	fmt.Println(len(q.stmt))
	return q
}



