package field

import "fmt"


type IntIField struct {
	name string
	size int
	ops _Option
}

func (i IntIField) GetSqlType() string {
	v := fmt.Sprintf(`%v INT`, i.name)
	if !i.ops.CheckOp(_Nullable) {
		v += " NOT NULL"
	}
	if i.ops.CheckOp(_AI) {
		v += " AUTO INCREMENT"
		return v
	}
	if i.ops.CheckOp(_Unique)  {
		v += " UNIQUE"
	}
	return v
}

func (i *IntIField) DefaultSize(size int) IField {
	i.size = size
	return i
}

func (i *IntIField) AI() IField {
	i.ops.SetOp(_AI)
	return i
}

func (i *IntIField) Unique() IField {
	i.ops.SetOp(_Unique)
	return i
}

func (i *IntIField) Nullable() IField {
	i.ops.SetOp(_Nullable)
	return i
}


