package repository

import (
	"github.com/OAOv/restful_CRUD/types"
)

type RecordRepository struct{}

func (r *RecordRepository) CreateRecord(record types.Record) error {
	stmt, err := db.Prepare("INSERT INTO record (id, user_id, user_name, subject, score) VALUES (?, ?, (SELECT name FROM user WHERE id = ?), ?, ?)")
	if err != nil {
		return types.ErrServerQueryError
	}
	defer stmt.Close()

	if record.ID == "" {
		record.ID = "0"
	}
	_, err = stmt.Exec(record.ID, record.UserID, record.UserID, record.Subject, record.Score)
	if err != nil {
		return types.ErrServerQueryError
	}

	return nil
}

func (r *RecordRepository) GetRecords() ([]types.Record, error) {
	var records []types.Record

	result, err := db.Query("SELECT * FROM record")
	if err != nil {
		return nil, types.ErrServerQueryError
	}
	defer result.Close()

	for result.Next() {
		var record types.Record
		err := result.Scan(&record.ID, &record.UserID, &record.UserName, &record.Subject, &record.Score)
		if err != nil {
			return nil, types.ErrInvalidParams
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *RecordRepository) GetRecord(id string) (types.Record, error) {
	var record types.Record
	result, err := db.Query("SELECT * FROM record WHERE id = ?", id)
	if err != nil {
		return record, types.ErrServerQueryError
	}
	defer result.Close()

	result.Next()
	err = result.Scan(&record.ID, &record.UserID, &record.UserName, &record.Subject, &record.Score)
	if err != nil {
		return record, types.ErrNotFound
	}

	return record, nil
}

func (r *RecordRepository) GetRecordByUser(id string) ([]types.Record, error) {
	var records []types.Record
	result, err := db.Query("SELECT * FROM record WHERE user_id = ?", id)
	if err != nil {
		return nil, types.ErrServerQueryError
	}
	defer result.Close()

	for result.Next() {
		var record types.Record
		err = result.Scan(&record.ID, &record.UserID, &record.UserName, &record.Subject, &record.Score)
		if err != nil {
			return nil, types.ErrInvalidParams
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *RecordRepository) UpdateReocrd(id string, record map[string]interface{}) error {
	var data types.Record
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Query("SELECT * FROM record WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return types.ErrServerQueryError
	}
	result.Next()
	err = result.Scan(&data.ID, &data.UserID, &data.UserName, &data.Subject, &data.Score)
	if err != nil {
		return types.ErrNotFound
	}
	result.Close()

	isFirst := true
	sql := "UPDATE record SET"
	for key, value := range record {
		if isFirst {
			sql += " " + key + " = \"" + value.(string) + "\""
			isFirst = false
		} else {
			sql += ", " + key + " = \"" + value.(string) + "\""
		}
	}
	sql += ", user_name = (SELECT name FROM user WHERE id = " + record["user_id"].(string) + ") WHERE id = " + id
	_, err = tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		return types.ErrServerQueryError
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *RecordRepository) DeleteRecord(id string) error {
	stmt, err := db.Prepare("DELETE FROM record WHERE id = ?")
	if err != nil {
		return types.ErrServerQueryError
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return types.ErrServerQueryError
	}

	return nil
}
