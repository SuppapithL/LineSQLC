// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const deleteFile = `-- name: DeleteFile :exec
DELETE FROM line_01 WHERE file_name = $1
`

func (q *Queries) DeleteFile(ctx context.Context, fileName string) error {
	_, err := q.db.ExecContext(ctx, deleteFile, fileName)
	return err
}

const getFileURL = `-- name: GetFileURL :one
SELECT file_content FROM line_01 WHERE file_name = $1
`

func (q *Queries) GetFileURL(ctx context.Context, fileName string) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getFileURL, fileName)
	var file_content sql.NullString
	err := row.Scan(&file_content)
	return file_content, err
}

const insertFileMetadata = `-- name: InsertFileMetadata :exec
INSERT INTO line_01 (user_id, file_name, file_content, created_at, theme) 
VALUES ($1, $2, $3, $4, $5)
`

type InsertFileMetadataParams struct {
	UserID      string
	FileName    string
	FileContent sql.NullString
	CreatedAt   time.Time
	Theme       sql.NullString
}

func (q *Queries) InsertFileMetadata(ctx context.Context, arg InsertFileMetadataParams) error {
	_, err := q.db.ExecContext(ctx, insertFileMetadata,
		arg.UserID,
		arg.FileName,
		arg.FileContent,
		arg.CreatedAt,
		arg.Theme,
	)
	return err
}

const listAllCategories = `-- name: ListAllCategories :many
SELECT DISTINCT theme FROM line_01
`

func (q *Queries) ListAllCategories(ctx context.Context) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, listAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var theme sql.NullString
		if err := rows.Scan(&theme); err != nil {
			return nil, err
		}
		items = append(items, theme)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFilesInCategory = `-- name: ListFilesInCategory :many
SELECT file_name FROM line_01 WHERE theme = $1
`

func (q *Queries) ListFilesInCategory(ctx context.Context, theme sql.NullString) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listFilesInCategory, theme)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var file_name string
		if err := rows.Scan(&file_name); err != nil {
			return nil, err
		}
		items = append(items, file_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const renameFile = `-- name: RenameFile :exec
UPDATE line_01 SET file_name = $1 WHERE file_name = $2
`

type RenameFileParams struct {
	FileName   string
	FileName_2 string
}

func (q *Queries) RenameFile(ctx context.Context, arg RenameFileParams) error {
	_, err := q.db.ExecContext(ctx, renameFile, arg.FileName, arg.FileName_2)
	return err
}

const updateFileURL = `-- name: UpdateFileURL :exec
UPDATE line_01 SET file_content = $1 WHERE file_name = $2
`

type UpdateFileURLParams struct {
	FileContent sql.NullString
	FileName    string
}

func (q *Queries) UpdateFileURL(ctx context.Context, arg UpdateFileURLParams) error {
	_, err := q.db.ExecContext(ctx, updateFileURL, arg.FileContent, arg.FileName)
	return err
}
