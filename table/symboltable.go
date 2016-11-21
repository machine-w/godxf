package table

import (
	"github.com/ma2ma/godxf/format"
	"github.com/ma2ma/godxf/handle"
)

// SymbolTable is interface for AcDbSymbolTableRecord.
type SymbolTable interface {
	IsSymbolTable() bool
	Format(format.Formatter)
	Handle() int
	SetHandle(*int)
	SetOwner(handle.Handler)
	Name() string
}
