package example

import (
	"github.com/thaianhsoft/gorm/relations"
	"github.com/thaianhsoft/gorm/schema"
	"testing"
)

type Note struct {
	schema.Schema
}

func (n *Note) Relations() []relations.RelEdge {
	return []relations.RelEdge{
		relations.EdgeBack("NoteOf", &Student{}).
			OnRefBack("HasNotes").
			Unique(),
	}
}
type Student struct {
	schema.Schema
}

func (s *Student) Relations() []relations.RelEdge {
	return []relations.RelEdge{
		relations.EdgeTo("HasNotes", &Note{}),
	}
}

func TestRels(t *testing.T) {
	rm := relations.NewRelManager()
	(&Note{}).Relations()
	(&Student{}).Relations()
	t.Log(*rm.GetAllEdges())
	for relName, edge := range *rm.GetAllEdges() {
		t.Log(relName, edge)
	}
}