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

func (r *RecordRepository) UpdateReocrd(record types.Record) error {
	if record.UserName == "" {
		var user types.User
		result, err := db.Query("SELECT * FROM user WHERE id = ?", record.UserID)
		if err != nil {
			return types.ErrServerQueryError
		}
		defer result.Close()

		result.Next()
		err = result.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return types.ErrNotFound
		}

		stmt, err := db.Prepare("UPDATE record SET user_id  = ?, user_name = (SELECT name FROM user WHERE id = ?), subject = ?, score = ? WHERE id = ?")
		if err != nil {
			return types.ErrServerQueryError
		}
		defer stmt.Close()

		_, err = stmt.Exec(record.UserID, record.UserID, record.Subject, record.Score, record.ID)
		if err != nil {
			return types.ErrServerQueryError
		}
	} else {
		stmt, err := db.Prepare("UPDATE record SET user_name = (SELECT name FROM user WHERE id = ?) WHERE user_id = ?")
		if err != nil {
			return types.ErrServerQueryError
		}
		defer stmt.Close()

		_, err = stmt.Exec(record.UserID, record.UserID)
		if err != nil {
			return types.ErrServerQueryError
		}
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
