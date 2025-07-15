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
