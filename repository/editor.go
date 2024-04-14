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
			"id": editorId,
		},
	); err != nil {
		return nil, err
	}
	return &model.Editor{
		ID:       editor.Id,
		Email:    editor.Email,
		Password: editor.Password,
	}, nil
}

func (r *EditorRepository) ReadEditorByEmail(ctx context.Context, email string) (*model.Editor, error) {
	var editor dbmodel.Editor
	if err := pgxscan.Get(
		ctx,
		r.pool,
		&editor,
		query.ReadEditorByEmail,
		pgx.NamedArgs{
			"email": email,
		},
	); err != nil {
		return nil, err
	}
	return &model.Editor{
		ID:       editor.Id,
		Email:    editor.Email,
		Password: editor.Password,
	}, nil
}
