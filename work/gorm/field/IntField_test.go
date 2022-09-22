package field

import (
	"fmt"
	"testing"
)

func TestIntIField(t *testing.T) {
	f := (&IntIField{}).AI()
	fmt.Println(f.GetSqlType())
}
