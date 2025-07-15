package main

import "fmt"

func main() {
	diaries := Diaries{}

	for {
		fmt.Println("\n=== Diary App ===")
		fmt.Println("1. Create Diary")
		fmt.Println("2. Login to Diary")
		fmt.Println("0. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		if choice < 0 || choice > 2 {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		switch choice {
		case 1:
			createDiary(&diaries)
		case 2:
			loginToDiary(&diaries)
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func createDiary(diary *Diaries) {
	fmt.Print("Enter your username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Enter your password: ")
	var password string
	fmt.Scanln(&password)

	diary.AddDiary(username, password)
	fmt.Printf("Diary for %s created successfully.\n", username)
}

func loginToDiary(diary *Diaries) {
	fmt.Print("Enter your username: ")
	var username string
	fmt.Scanln(&username)

	foundDiary := diary.FindDiary(username)

	if foundDiary != nil {
		fmt.Printf("Welcome back, %s!\n", username)
		diaryMenu(foundDiary)
	} else {
		fmt.Println("Invalid username.")
	}
}

func diaryMenu(diary *Diary) {
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

		switch choice {
		case 1:
			if diary.IsLocked {
				fmt.Println("Diary is locked. Unlock to add entries.")
				continue
			}
			var title, body string
			fmt.Print("Title: ")
			fmt.Scanln(&title)
			fmt.Print("Body: ")
			fmt.Scanln(&body)
			diary.CreateEntry(title, body)
			fmt.Println("Entry created.")

		case 2:
			if diary.IsLocked {
				fmt.Println("Diary is locked. Unlock to view entries.")
				continue
			}
			for _, entry := range diary.Entries {
				fmt.Printf("ID: %d\nTitle: %s\nBody: %s\nDate: %s\n\n", entry.ID, entry.Title, entry.Body, entry.DateCreated.Format("02 Jan 2006 15:04"))
			}

		case 3:
			var id int
			var title, body string
			fmt.Print("Entry ID to update: ")
			fmt.Scanln(&id)
			fmt.Print("New Title: ")
			fmt.Scanln(&title)
			fmt.Print("New Body: ")
			fmt.Scanln(&body)
			diary.UpdateDiary(id, title, body)

		case 4:
			var id int
			fmt.Print("Entry ID to delete: ")
			fmt.Scanln(&id)
			diary.DeleteEntry(id)

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
