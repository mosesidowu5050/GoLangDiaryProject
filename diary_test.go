package main

import (
	"testing"
)

func TestUnlockDiary(test *testing.T) {

	diary := Diary{Username: "testuser", Password: "password", IsLocked: true}
	success := diary.UnlockDiary("password")

	if !success || diary.IsLocked {
		test.Errorf("Expected diary to be unlocked with correct password.")
	}

	if diary.UnlockDiary("wrongpassword") {
		test.Errorf("Expected diary to remain locked with incorrect password.")
	}
}

func TestDiaryIsLocked(test *testing.T) {

	diary := Diary{Username: "testuser", Password: "password", IsLocked: false}
	success := diary.LockDiary("password")

	if !success {
		test.Errorf("Expected diary to be locked with correct password.")
	}
	if !diary.IsLocked {
		test.Errorf("Diary should be locked after successful lock attempt.")
	}
}

func TestCreateEntry(test *testing.T) {

	diary := Diary{Username: "testuser", Password: "password", IsLocked: false}
	diary.CreateEntry("Test Entry", "This is a test entry.")

	if len(diary.Entries) != 1 {
		test.Errorf("Expected 1 entry in diary, got %d", len(diary.Entries))
	}
	if len(diary.Entries) == 0 {
		test.Errorf("Expected diary to have entries, but it is empty.")
	}
}

func TestCannotCreateEntryWhenDiaryLocked(test *testing.T) {
	diary := Diary{Username: "testuser", Password: "password", IsLocked: true}
	success := diary.CreateEntry("Test Entry", "This is a test entry.")

	if success {
		test.Errorf("Expected entry creation to fail on locked diary.")
	}
	if len(diary.Entries) != 0 {
		test.Errorf("Expected no entries in diary, got %d", len(diary.Entries))
	}
}

func TestCreateMultipleEntries(test *testing.T) {
	diary := Diary{Username: "testuser", Password: "password", IsLocked: false}
	diary.CreateEntry("First Entry", "This is the first entry.")
	diary.CreateEntry("Second Entry", "This is the second entry.")

	if len(diary.Entries) != 2 {
		test.Errorf("Expected 2 entries in diary, got %d", len(diary.Entries))
	}
	if diary.Entries[0].Title != "First Entry" || diary.Entries[1].Title != "Second Entry" {
		test.Errorf("Expected specific titles for entries, got %s and %s", diary.Entries[0].Title, diary.Entries[1].Title)
	}
}

func TestDeleteEntry(test *testing.T) {
	diary := Diary{Username: "testuser", Password: "password", IsLocked: false}
	success := diary.CreateEntry("Test Entry", "This is a test entry.")

	if !success {
		test.Errorf("Expected entry creation to succeed.")
	}

	diary.DeleteEntryById(1)

	if len(diary.Entries) != 0 {
		test.Errorf("Expected no entries in diary after deletion, got %d", len(diary.Entries))
	}
}

func TestCannotDeleteEntryWhenDiaryIsLocked(test *testing.T) {
	diary := Diary{Username: "testuser", Password: "password", IsLocked: true}
	success := diary.CreateEntry("Test Entry", "This is a test entry.")

	if success {
		test.Errorf("Expected entry creation not to succeed.")
	}

	diary.DeleteEntryById(1)
	if !diary.IsLocked {
		test.Errorf("Dairy cannot be deleted when locked.")
	}
}

func TestFindEntryById(test *testing.T) {
	diary := Diary{Username: "testuser", Password: "password", IsLocked: false}
	successOne := diary.CreateEntry("Test Entry1", "This is a test entry 1.")
	successTwo := diary.CreateEntry("Test Entry2", "This is a test entry 2.")
	entry, foundEntry := diary.FindEntryById(1)

	if !successOne && !successTwo {
		test.Errorf("Expected entry creation to succeed.")
	}

	if len(diary.Entries) != 2 {
		test.Errorf("Expected 1 entry in diary, got %d", len(diary.Entries))
	}

	if !foundEntry {
		test.Errorf("Expected entry to be found, but entry ID not found.")
	}
	if entry.Title != "Test Entry1" {
		test.Errorf("Expected title 'Test Entry1', got '%s'", entry.Title)
	}
}

func TestFindEntryByIdNotFound(test *testing.T) {
	diary := Diary{Username: "testuser", Password: "password", IsLocked: false}
	diary.CreateEntry("Test Entry", "This is a test entry.")

	entry, foundEntry := diary.FindEntryById(999)
	if foundEntry {
		test.Errorf("Expected entry not to be found, but it was.")
	}
	if entry != nil {
		test.Errorf("Expected entry to be nil when not found, got %v", entry)
	}
}

func TestUpdateEntryWithIdTitleBody(test *testing.T) {
	diary := Diary{Username: "testuser", Password: "password", IsLocked: false}
	diary.CreateEntry("Test Entry", "This is a test entry.")

	entry, foundEntry := diary.FindEntryById(1)
	if !foundEntry {
		test.Errorf("Expected entry not to be found, but it was.")
	}
	if entry.Title != "Test Entry" {
		test.Errorf("Expected title 'Test Entry', got '%s'", entry.Title)
	}

	diary.UpdateEntry(2, "Updated Entry", "Entry is updated successfully")
	if len(diary.Entries) != 1 {
		test.Errorf("Expected 1 entry in diary, got %d", len(diary.Entries))
	}
}
