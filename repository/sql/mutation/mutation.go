package mutation

import (
	_ "embed"
)

var (
	//go:embed scripts/editor/Create.sql
	CreateEditor string
)
