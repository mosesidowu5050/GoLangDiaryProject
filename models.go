package main

import (
	"fmt"
	"time"
)

type Entry struct {
	ID          int
	Title       string
	Body        string
	DateCreated time.Time
}

type Diary struct {
	Username string
	Password string
	IsLocked bool
	Entries  []Entry
}

type Diaries struct {
	DiaryList []Diary
}

func (diary *Diary) UnlockDiary(password string) bool {
	if diary.Password != password {
		return false
	}
	diary.IsLocked = false
	return true
}

func (diary *Diary) LockDiary(password string) bool {
	if diary.Password != password {
		return false
	}
	diary.IsLocked = true
	return true
}

func (diary *Diary) CreateEntry(title string, body string) bool {
	if diary.IsLocked {
		fmt.Print("Diary is locked, cannot create entry. \n")
		return false
	}

	entry := Entry{
		ID:          len(diary.Entries) + 1,
		Title:       title,
		Body:        body,
		DateCreated: time.Now(),
	}
	diary.Entries = append(diary.Entries, entry)
	return true
}

func (diary *Diary) DeleteEntryById(deleteId int) {
	for check, entry := range diary.Entries {
		if entry.ID == deleteId {
			diary.Entries = append(diary.Entries[:check], diary.Entries[check+1:]...)
			fmt.Printf("Entry with ID %d deleted successfully.\n", deleteId)
			break
		}
		fmt.Printf("Entry with ID %d not found.\n", deleteId)
	}

}
