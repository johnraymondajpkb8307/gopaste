// storage/sqlite.go
package storage

import (
	"database/sql"
	"fmt"
	"gopaste/backend/models"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(dbPath string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// 创建表
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			created_at INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS pastes (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at INTEGER NOT NULL,
			expires_at INTEGER NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
    `)
	if err != nil {
		return nil, err
	}

	return &SQLiteStore{db: db}, nil
}

func (s *SQLiteStore) SavePaste(paste *models.Paste) error {
	_, err := s.db.Exec(
		"INSERT INTO pastes (id, content, created_at, expires_at) VALUES (?, ?, ?, ?)",
		paste.ID, paste.Content, paste.CreatedAt, paste.ExpiresAt,
	)
	return err
}

func (s *SQLiteStore) GetPaste(id string) (*models.Paste, error) {
	paste := &models.Paste{}
	err := s.db.QueryRow(
		"SELECT id, content, created_at, expires_at FROM pastes WHERE id = ? AND expires_at > ?",
		id, time.Now().Unix(),
	).Scan(&paste.ID, &paste.Content, &paste.CreatedAt, &paste.ExpiresAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("paste not found")
	}
	return paste, err
}

func (s *SQLiteStore) CreateUser(username, password string) (*models.User, error) {
	user := &models.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: password,
	}

	_, err := s.db.Exec(
		"INSERT INTO users (id, username, password) VALUES (?, ?, ?)",
		user.ID, user.Username, user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *SQLiteStore) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := s.db.QueryRow(
		"SELECT id, username, password FROM users WHERE username = ?",
		username,
	).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *SQLiteStore) UsernameExists(username string) bool {
	var exists bool
	err := s.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)",
		username,
	).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
