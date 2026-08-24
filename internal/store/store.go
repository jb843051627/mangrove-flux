package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jb843051627/mangrove-flux/internal/model"
	_ "modernc.org/sqlite"
)

type Store struct{ db *sql.DB }

func Open(path string) (*Store, error) {
	if path == "" || path == ":memory:" {
		return nil, model.ErrInvalid
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	s := &Store{db: db}
	if err := s.schema(context.Background()); err != nil {
		_ = db.Close()
		return nil, err
	}
	return s, nil
}

func (s *Store) schema(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS records (kind TEXT NOT NULL, id TEXT NOT NULL, payload BLOB NOT NULL, revision INTEGER NOT NULL DEFAULT 0, updated_at TEXT NOT NULL, PRIMARY KEY(kind,id)); CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY AUTOINCREMENT, subject TEXT NOT NULL, action TEXT NOT NULL, created_at TEXT NOT NULL); CREATE INDEX IF NOT EXISTS records_kind_updated ON records(kind, updated_at);`)
	return err
}

func (s *Store) Close() error { return s.db.Close() }

func (s *Store) SaveContext(ctx context.Context, kind, id string, value any) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, `INSERT INTO records(kind,id,payload,revision,updated_at) VALUES(?,?,?,0,?) ON CONFLICT(kind,id) DO UPDATE SET payload=excluded.payload, revision=records.revision+1, updated_at=excluded.updated_at`, kind, id, raw, time.Now().UTC().Format(time.RFC3339Nano))
	return err
}

func (s *Store) Save(kind, id string, value any) error {
	return s.SaveContext(context.Background(), kind, id, value)
}

func (s *Store) LoadContext(ctx context.Context, kind, id string, value any) error {
	var raw []byte
	err := s.db.QueryRowContext(ctx, `SELECT payload FROM records WHERE kind=? AND id=?`, kind, id).Scan(&raw)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w: %s/%s", model.ErrNotFound, kind, id)
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, value)
}

func (s *Store) Load(kind, id string, value any) error {
	return s.LoadContext(context.Background(), kind, id, value)
}

func (s *Store) List(ctx context.Context, kind string, decode func([]byte) error) error {
	rows, err := s.db.QueryContext(ctx, `SELECT payload FROM records WHERE kind=? ORDER BY updated_at, id`, kind)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var raw []byte
		if err := rows.Scan(&raw); err != nil {
			return err
		}
		if err := decode(raw); err != nil {
			return err
		}
	}
	return rows.Err()
}

func (s *Store) DeleteContext(ctx context.Context, kind, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM records WHERE kind=? AND id=?`, kind, id)
	return err
}
func (s *Store) Delete(kind, id string) error { return s.DeleteContext(context.Background(), kind, id) }

func (s *Store) Audit(ctx context.Context, subject, action string) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO events(subject,action,created_at) VALUES(?,?,?)`, subject, action, time.Now().UTC().Format(time.RFC3339Nano))
	return err
}

func (s *Store) WithTx(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		_ = tx.Commit()
		return fmt.Errorf("%w: %v", model.ErrTransaction, err)
	}
	return tx.Commit()
}

func TxSave(tx *sql.Tx, kind, id string, value any) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`INSERT INTO records(kind,id,payload,revision,updated_at) VALUES(?,?,?,0,?) ON CONFLICT(kind,id) DO UPDATE SET payload=excluded.payload, revision=records.revision+1, updated_at=excluded.updated_at`, kind, id, raw, time.Now().UTC().Format(time.RFC3339Nano))
	return err
}
