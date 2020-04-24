package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

// AnswerData : アンサーデータテーブル
type AnswerData struct {
    gorm.Model
    Name       string
    QuestionNo int
    Language   string
    SourceCode string
    Message    string
    Result     string
    Error      string
}

func (a *AnswerData) CreateAtTime() string {
    return a.CreatedAt.Format("2006/01/02 15:04:05")
}

func openDB(dest string, addDBCh chan *AnswerData) error {
    for data := range addDBCh {
        db, err := gorm.Open("sqlite3", dest)
        if err != nil {
            return err
        }

        // Migrate the schema
        db.AutoMigrate(&AnswerData{})

        // Create
        db.Create(&data)

        db.Close()
    }

    return nil
}
