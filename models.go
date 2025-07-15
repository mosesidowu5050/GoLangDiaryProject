package main

import (
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
