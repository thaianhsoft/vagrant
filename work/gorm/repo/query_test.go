package repo

import (
	"fmt"
	"testing"
)

func TestQueryBuilder(t *testing.T) {
	q := NewQueryBuilder()
	q.Select(C("Id", "users")).
		Select(C("Age", "users").As("PetRel$Age")).
		Where(P("Id", "users", Equal, 3)).
		Where(P("Age", "users", LessThanEqual, 15), PrefixAnd).
		Where(P("Name", "users", LikeString, "%ThaiAnh%"), PrefixOr)
	query, args := q.ToQuery()
	fmt.Println(query, args)
}
