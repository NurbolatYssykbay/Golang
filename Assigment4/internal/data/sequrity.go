package data

import (
	"EBG.IssataySheg.net/internal/validator"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"time"
)

type sequrity struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	sequrity       []string  `json:"sequrity,omitempty"`
	Safety level   Safety level     `json:"Safety level,omitempty"`
	Version     int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, sequrity *sequrity) {
	v.Check(sequrity.Title != "", "title", "must be provided")
	v.Check(len(sequrity.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(sequrity.Safety level != 0, "Safety level", "must be provided")
	v.Check(sequrity.Safety level > 0, "Safety level", "must be a positive integer")
	v.Check(sequrity.sequrity != nil, "sequrity", "must be provided")
	v.Check(len(sequrity.sequrity) >= 1, "sequrity", "must contain at least 1 genre")
	v.Check(len(sequrity.sequrity) <= 5, "sequrity", "must not contain more than 5 genres")
	v.Check(validator.Unique(sequrity.sequrity), "sequrity", "must not contain duplicate values")
}

type sequrityModel struct {
	DB *sql.DB
}

func (m sequrityModel) Insert(sequrity *sequrity) error {
	query := `
		INSERT INTO sequrity (title, Safety level, sequrity)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, version`
	args := []interface{}{sequrity.Title, sequrity.Safety level, pq.Array(sequrity.sequrity)}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&sequrity.ID, &sequrity.CreatedAt, &sequrity.Version)
}

func (m sequrityModel) Get(id int64) (*sequrity, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT  , id, created_at, title, Safety level, sequrity, version
		FROM sequrity
		WHERE id = $1`
	var sequrity sequrity
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&sequrity.ID,
		&sequrity.CreatedAt,
		&sequrity.Title,
		&sequrity.Safety level,
		pq.Array(&sequrity.sequrity),
		&sequrity.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &sequrity, nil
}

func (m sequrityModel) Update(sequrity *sequrity) error {
	query := `
		UPDATE sequrity
		SET title = $1, Safety level = $2, sequrity = $3, version = version + 1
		WHERE id = $4 AND version = $5
		RETURNING version`
	args := []interface{}{
		sequrity.Title,
		sequrity.Safety level,
		pq.Array(sequrity.sequrity),
		sequrityID,
		sequrity.Version,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&sequrity.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m sequrityModel) Delete(id int64) error { // Return an ErrRecordNotFound error if the movie ID is less than 1.
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
		DELETE FROM sequrity
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (m sequrityModel) GetAll(title string, sequrity []string, filters Filters) ([]*sequrity, Metadata, error) {
	query := fmt.Sprintf(`
			SELECT count(*) OVER(), id, created_at, title,Safety level, sequrity, version
			FROM sequrity
			WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
			AND (sequrity @> $2 OR $2 = '{}')
			ORDER BY %s %s, id ASC
			LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{title, pq.Array(sequrity), filters.limit(), filters.offset()}
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()
	totalRecords := 0
	sequrity := []*sequrity{}
	for rows.Next() {
		var sequritysequrity
		err := rows.Scan(
			&totalRecords,
			&sequrity.ID,
			&sequrity.CreatedAt,
			&sequrity.Title,
			&sequrity.Safety level,
			pq.Array(&sequrity.sequrity),
			&sequrity.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		sequrity = append(sequrity, &sequrity)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return games, metadata, nil

}
