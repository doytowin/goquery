package web

import (
	. "github.com/doytowin/goquery/core"
	"regexp"
)

type RestAPI[E any, Q GoQuery] interface {
	Page(q Q) (PageList[E], error)
	Get(id any) (*E, error)
}

type Service[C any, E any, Q GoQuery] struct {
	c            C
	dataAccess   DataAccess[C, E]
	createQuery  func() Q
	createEntity func() E
	idRgx        *regexp.Regexp
}

func (s *Service[C, E, Q]) Page(q Q) (PageList[E], error) {
	return s.dataAccess.Page(s.c, q)
}

func (s *Service[C, E, Q]) Get(id any) (*E, error) {
	return s.dataAccess.Get(s.c, id)
}

func BuildService[C any, E any, Q GoQuery](
	prefix string, c C,
	dataAccess DataAccess[C, E],
	createEntity func() E,
	createQuery func() Q,
) *Service[C, E, Q] {
	return &Service[C, E, Q]{
		c:            c,
		dataAccess:   dataAccess,
		createQuery:  createQuery,
		createEntity: createEntity,
		idRgx:        regexp.MustCompile(prefix + `(\d+)$`),
	}
}