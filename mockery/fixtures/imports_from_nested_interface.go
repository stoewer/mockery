package test

import (
	"github.com/stoewer/mockery/mockery/fixtures/http"
)

type HasConflictingNestedImports interface {
	RequesterNS
	Z() http.MyStruct
}
