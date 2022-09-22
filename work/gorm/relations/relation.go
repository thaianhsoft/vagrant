package relations

import (
	"github.com/thaianhsoft/gorm/schema"
	"reflect"
	"runtime"
	"sync"
)

var RelManager *_RelManager = nil

type _RelManager struct {
	locker *sync.Mutex
	relEdges map[string]RelEdge
}

func (r *_RelManager) GetRel(relName string) RelEdge {
	r.locker.Lock()
	defer r.locker.Unlock()
	if relEdge, ok := r.relEdges[relName]; ok {
		return relEdge
	}
	return nil
}

func (r *_RelManager) AddRelEdge(relName string, rel RelEdge) {
	r.locker.Lock()
	defer r.locker.Unlock()
	if _, ok := r.relEdges[relName]; !ok {
		r.relEdges[relName] = rel
	}
}

func (r *_RelManager) GetAllEdges() *map[string]RelEdge {
	return &r.relEdges
}

type RelEdge interface {
	Unique() RelEdge
	OnRefBack(refName string) RelEdge
}

type rel struct {
	unique bool
	isRef bool
	relName string
	tb string
	relRefName string
}

func (r *rel) Unique() RelEdge {
	r.unique = true
	return r
}

func (r *rel) OnRefBack(refName string) RelEdge {
	defer func() {
		go func() {
			for RelManager.GetRel(refName) == nil {
				runtime.Gosched()
			}
			refRel := RelManager.GetRel(refName)
			r.relRefName = refRel.(*rel).relName
			refRel.(*rel).relRefName = r.relName

		}()
	}()
	return r
}

func NewRelManager() *_RelManager {
	if RelManager == nil {
		RelManager = &_RelManager{
			locker:   &sync.Mutex{},
			relEdges: make(map[string]RelEdge),
		}
	}
	return RelManager
}

func EdgeTo(relName string, relToTable schema.Schema) RelEdge {
	r := &rel{
		unique:     false,
		isRef:      false,
		relName:    relName,
		tb:         reflect.Indirect(reflect.ValueOf(relToTable)).Type().Name(),
		relRefName: "",
	}
	RelManager.AddRelEdge(relName, r)
	return r
}

func EdgeBack(relName string, relToTable schema.Schema) RelEdge {
	r := &rel{
		unique:     false,
		isRef:      true,
		relName:    "",
		tb:         reflect.Indirect(reflect.ValueOf(relToTable)).Type().Name(),
		relRefName: "",
	}
	RelManager.AddRelEdge(relName, r)
	return r
}