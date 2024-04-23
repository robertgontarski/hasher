package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store interface {
	UpdateHashOnEmailByID(id int, hash string) error
	UpdateHashOnPhoneByID(id int, hash string) error
	UpdateHashOnNameByID(id int, hashName string, hashSurname string) error
}

type MysqlStore struct {
	Ctx context.Context
	Db  *sql.DB
	Store
}

func NewMysqlStore(ctx context.Context, db *MysqlDatabase) *MysqlStore {
	return &MysqlStore{
		Ctx: ctx,
		Db:  db.GetDb(),
	}
}

func (s *MysqlStore) UpdateHashOnEmailByID(id int, hash string) error {
	ctx, cancel := context.WithTimeout(s.Ctx, time.Millisecond*300)
	defer cancel()

	if _, err := s.Db.QueryContext(ctx, "UPDATE clinet SET hash_email = ? WHERE id = ?", hash, id); err != nil {
		return fmt.Errorf("UpdateHashOnEmailByID: %v", err)
	}

	return nil
}

func (s *MysqlStore) UpdateHashOnPhoneByID(id int, hash string) error {
	ctx, cancel := context.WithTimeout(s.Ctx, time.Millisecond*300)
	defer cancel()

	if _, err := s.Db.QueryContext(ctx, "UPDATE client SET hash_number = ? WHERE id = ?", hash, id); err != nil {
		return fmt.Errorf("UpdateHashOnPhoneByID: %v", err)
	}

	return nil
}

func (s *MysqlStore) UpdateHashOnNameByID(id int, hashName string, hashSurname string) error {
	ctx, cancel := context.WithTimeout(s.Ctx, time.Millisecond*300)
	defer cancel()

	if _, err := s.Db.QueryContext(ctx, "UPDATE client SET hash_name = ?, hash_surname = ? WHERE id = ?", hashName, hashSurname, id); err != nil {
		return fmt.Errorf("UpdateHashOnNameByID: %v", err)
	}

	return nil
}
