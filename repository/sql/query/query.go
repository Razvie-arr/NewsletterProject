package query

import (
	_ "embed"
)

var (
	//go:embed scripts/editor/Read.sql
	ReadEditor string
)
