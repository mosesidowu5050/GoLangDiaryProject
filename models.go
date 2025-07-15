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

func (diaries *Diaries) AddDiary(username string, password string) {

	for _, diary := range diaries.DiaryList {
		if diary.Username == username {
			fmt.Printf("Diary with username %s already exists.\n", username)
			return
		}
		if diary.Username == "" && diary.Password == "" {
			fmt.Printf("Diary with username %s is empty, please fill it.\n", username)
			return
		}
		if diary.Password != password {
			fmt.Printf("Password for the diary with username %s does not match.\n", username)
			return
		}
	}

	diary := Diary{Username: username, Password: password, IsLocked: true}
	diaries.DiaryList = append(diaries.DiaryList, diary)
	fmt.Printf("Diary with username %s has been added successfully.\n", username)
}

func (diaries *Diaries) FindDiary(username string) *Diary {
	for check := range diaries.DiaryList {
		if diaries.DiaryList[check].Username == username {
			return &diaries.DiaryList[check]
		}
	}
	return nil
}

func (diaries *Diaries) DeleteDiary(username string, password string) {
	for check := range diaries.DiaryList {
		if diaries.DiaryList[check].Username == username &&
			diaries.DiaryList[check].Password == password {
			diaries.DiaryList = append(diaries.DiaryList[:check], diaries.DiaryList[check+1:]...)
			fmt.Printf("Diary with username %s has been removed successfully.\n", username)
			return
		}
		fmt.Printf("Diary with username %s not found.\n", username)
	}
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
	if diary.IsLocked {
		fmt.Print("Diary is locked, cannot delete entry. \n")
		return
	}

	for check, entry := range diary.Entries {
		if entry.ID == deleteId {
			diary.Entries = append(diary.Entries[:check], diary.Entries[check+1:]...)
			fmt.Printf("Entry with ID %d deleted successfully.\n", deleteId)
			break
		}
		fmt.Printf("Entry with ID %d not found.\n", deleteId)
	}

}

func (diary *Diary) FindEntryById(entryId int) (*Entry, bool) {
	if diary.IsLocked {
		fmt.Println("Diary is locked, cannot find entry.")
		return nil, false
	}

	for _, entry := range diary.Entries {
		if entry.ID == entryId {
			return &entry, true
		}
	}
	return nil, false
}

func (diary *Diary) UpdateEntry(entryId int, title string, body string) {
	if diary.IsLocked {
		fmt.Print("Diary is locked, cannot update entry. \n")
		return
	}

	for checkId := range diary.Entries {
		if diary.Entries[checkId].ID == entryId {
			diary.Entries[checkId].Title = title
			diary.Entries[checkId].Body = body
			diary.Entries[checkId].DateCreated = time.Now()
			fmt.Printf("Entry with ID %d updated successfully.\n", entryId)
			return
		}
		fmt.Printf("Entry with ID %d not found.\n", entryId)
	}
}
