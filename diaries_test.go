package main

import "testing"

func TestThatAddToDiaries(test *testing.T) {

	diaries := Diaries{}
	diaries.AddDiary("User1", "password")

	if len(diaries.DiaryList) != 1 {
		test.Errorf("Expected 1 diary, got %d", len(diaries.DiaryList))
	}

	diaries.AddDiary("user1", "abcd")
	if len(diaries.DiaryList) != 1 {
		test.Errorf("Duplicate diary should not be added")
	}
}

func TestFindDairyByUsername(test *testing.T) {

	diaries := Diaries{}
	diaries.AddDiary("User2", "password2")
	diaries.AddDiary("User3", "password3")

	diary := diaries.FindDiary("User2")
	if diary == nil || diary.Username != "User2" {
		test.Error("Expected to find diary with username 'user2'")
	}
	notFound := diaries.FindDiary("unknown")
	if notFound != nil {
		test.Error("Expected nil for unknown username")
	}
}

func TestDeleteDiary(test *testing.T) {
	diaries := Diaries{}
	diaries.AddDiary("User4", "password4")
	diary := diaries.FindDiary("User4")

	if diary == nil || diary.Username != "User4" {
		test.Error("Expected to find diary with username 'user4'")
	}

	diaries.DeleteDiary("User4", "password4")
	if len(diaries.DiaryList) != 0 {
		test.Error("Expected diary with username 'user4' to be deleted")
	}
}
