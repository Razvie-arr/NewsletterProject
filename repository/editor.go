package repository

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"newsletterProject/pkg/id"
	dbmodel "newsletterProject/repository/sql/model"
	"newsletterProject/repository/sql/query"
	"newsletterProject/service/model"
)

type EditorRepository struct {
	pool *pgxpool.Pool
}

func NewEditorRepository(pool *pgxpool.Pool) *EditorRepository {
	return &EditorRepository{
		pool: pool,
	}
}

func (r *EditorRepository) ReadEditor(ctx context.Context, editorId id.ID) (*model.Editor, error) {
	var editor dbmodel.Editor
	if err := pgxscan.Get(
		ctx,
		r.pool,
		&editor,
		query.ReadEditor,
		pgx.NamedArgs{
			"id": editor,
		},
	); err != nil {
		return nil, err
	}
	return &model.Editor{}, nil
}
