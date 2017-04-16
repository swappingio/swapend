package db

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Song struct {
	Id       int64
	SongID   string
	Created  time.Time
	Name     string
	Uploader int64
	Versions []string
}

func CreateVersion(SongUUID string) string {

	newUUID := uuid.New().String()

	res, err := db.Exec("UPDATE songs SET versions = array_append(versions, $1) WHERE songid = $2",
		newUUID, SongUUID)
	if err != nil {
		fmt.Println(err)
	}

	affectedRows := res.RowsAffected()

	if affectedRows != 1 {
		return ""
	}

	_, err = db.Exec("INSERT INTO versions (created, fileid, uploader) VALUES (now(), $1, $2)", newUUID, 1)
	if err != nil {
		fmt.Println(err)
	}
	return newUUID
}

func CreateSong(s Song) (string, error) {
	var newUUID = uuid.New().String()
	_, err := db.Exec("INSERT INTO songs (songid, name, uploader, created) VALUES ($1, $2, $3, now())",
		newUUID, "LUL", 1)
	if err != nil {
		fmt.Println(err)
	}
	return newUUID, nil
}

func GetSong(SongUUID string) (Song, error) {
	var s Song
	err := db.QueryRow("SELECT * FROM songs WHERE songid = $1", SongUUID).Scan(
		&s.Id, &s.SongID, &s.Created, &s.Name, &s.Uploader, &s.Versions)
	if err != nil {
		fmt.Println(err)
	}
	return s, nil
}
