package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	diaries := Diaries{}

	for {
		fmt.Println("\n=== Diary App ===")
		fmt.Println("1. Create Diary")
		fmt.Println("2. Login to Diary")
		fmt.Println("0. Exit")

		userChoice := getInput("Choose an option: ")
		choice, err := strconv.Atoi(userChoice)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		if choice < 0 || choice > 2 {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		switch choice {
		case 1:
			simulateLoading("Loading diary setup ", 5)
			createDiary(&diaries)
		case 2:
			simulateLoading("Accessing Login ", 3)
			loginToDiary(&diaries)
		case 0:
			simulateLoading("Existing app ", 3)
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func createDiary(diary *Diaries) {
	username := getInput("Enter your username: ")
	password := getInput("Enter your password: ")
	diary.AddDiary(username, password)
}

func loginToDiary(diary *Diaries) {
	username := getInput("Enter your username: ")
	foundDiary := diary.FindDiary(username)
	if foundDiary != nil {
		fmt.Printf("Welcome back, %s!\n", username)
		diaryMenu(foundDiary)
	} else {
		fmt.Println("Invalid username.")
	}
}

func diaryMenu(diary *Diary) {

	if diary.IsLocked {
		fmt.Println("This diary is currently locked.")
		password := getInput("Enter password to unlock: ")
		if !diary.UnlockDiary(password) {
			simulateLoading("Failed to unlock diary with invalid password. Returning to main menu.", 3)
			return
		}
		fmt.Println("Diary unlocked successfully.")
	}

	for {
		fmt.Println("\n--- Diary Menu ---")
		fmt.Println("1. Add Entry")
		fmt.Println("2. View Entries")
		fmt.Println("3. Update Entry")
		fmt.Println("4. Delete Entry")
		fmt.Println("5. Lock Diary & Exit")
		fmt.Println("0. Logout")

		var choice int
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		if choice < 0 || choice > 5 {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		switch choice {
		case 1:
			if diary.IsLocked {
				fmt.Println("Diary is locked. Unlock to add entries.")
				continue
			}
			title := getInput("Enter entry title: ")
			body := getInput("Enter entry body: ")

			diary.CreateEntry(title, body)
			fmt.Println("Entry created.")

		case 2:
			if diary.IsLocked {
				fmt.Println("Diary is locked. Unlock to view entries.")
				continue
			}
			for _, entry := range diary.Entries {
				fmt.Printf("ID: %d\nTitle: %s\nBody: %s\nDate: %s\n\n",
					entry.ID, entry.Title, entry.Body, entry.DateCreated.Format("02 Jan 2006 15:04"))
			}

		case 3:
			userId := getInput("Enter your ID to update: ")
			id, err := strconv.Atoi(userId)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a valid number.")
				continue
			}

			title := getInput("Enter entry title: ")
			body := getInput("Enter entry body: ")

			diary.UpdateEntry(id, title, body)

		case 4:
			userId := getInput("Enter your ID to update: ")
			id, err := strconv.Atoi(userId)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a valid number.")
				continue
			}
			diary.DeleteEntryById(id)

		case 5:
			diary.IsLocked = true
			fmt.Println("Diary locked. Returning to main menu.")
			return

		case 0:
			fmt.Println("Logging out...")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func simulateLoading(message string, seconds int) {
	fmt.Print(message)
	for count := 0; count < seconds; count++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()
}
