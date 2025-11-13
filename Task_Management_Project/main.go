package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	username string
	name     string
	password string
	email    string
}

type taskStatus struct {
	username   string
	task       string
	status     string
	start_time string
	end_time   string
	task_id    int
}

func main() {

	var choice = 0
	scanner := bufio.NewScanner(os.Stdin)
Loop:
	for {
		fmt.Println("User choices: ")
		fmt.Println("**************")
		fmt.Println("1.Admin Login.")
		fmt.Println("2.User Login.")
		fmt.Println("3.Registration.")
		fmt.Println("4.Exit.")
		fmt.Println("*******************")
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			databaseConnection(choice, scanner)
		case 2:
			databaseConnection(choice, scanner)
		case 3:
			databaseConnection(choice, scanner)
		case 4:
			break Loop
		}
	}
}
func userRegistration(db *sql.DB) {

	var userInfo user
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Println("Enter an unique username: ")
		scanner.Scan()
		userInfo.username = scanner.Text()

		rows, err := db.Query("select username from userInfo")
		if err != nil {

			log.Fatal("query failed: %v", err)
		}

		var cnt int //to check for duplicate username
		for rows.Next() {

			var tempUser user
			rows.Scan(&tempUser.username)
			if userInfo.username == tempUser.username {

				cnt++
				fmt.Println("Username already exists!!")
			} else if tempUser.username == "" {

				break
			}
		}
		if cnt == 0 {

			break
		}
	}

	fmt.Println("Enter your name: ")
	scanner.Scan()
	userInfo.name = scanner.Text()
	fmt.Println("Enter your password: ")
	scanner.Scan()
	userInfo.password = scanner.Text()
	fmt.Println("Enter your email: ")
	scanner.Scan()
	userInfo.email = scanner.Text()

	query := "insert into userInfo(username, name, password, email) values(?, ?, ?, ?)"
	_, err := db.Exec(query, userInfo.username, userInfo.name, userInfo.password, userInfo.email)
	if err != nil {

		log.Fatal("database insertion failed!!")
	}
}
func userLogin(db *sql.DB, scanner *bufio.Scanner) {

	var username string
	var password string
	for {

		fmt.Println("Enter your username: ")
		fmt.Scanln(&username)
		fmt.Println("Enter your password: ")
		fmt.Scanln(&password)

		query := "select count(*) from userInfo where username = ? and password = ?"
		var cnt int
		err := db.QueryRow(query, username, password).Scan(&cnt)
		if err != nil {

			if err == sql.ErrNoRows {

				fmt.Println("User not found!!")
			} else {
				log.Fatal("query failed: %v", err)
			}
		}
		if cnt > 0 {
			fmt.Println("Login successful!!")
			break
		}

	}
	choice := 0
	for {

		fmt.Println("Your choices are: ")
		fmt.Println("******************")
		fmt.Println("1.View tasks.")
		fmt.Println("2.Complete a task.")
		fmt.Println("3.Update your Information.")
		fmt.Println("4.Exit.")
		fmt.Println("******************")
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {

		case 1:
			query := "select task, status, start_time, end_time from taskStatus where username = ?"
			rows, err := db.Query(query, username)
			if err != nil {

				log.Fatal("data retrieve failed: %v", err)
			}
			fmt.Println("task                                          status               stat_time                   end_time")
			fmt.Println("*****                                        ********             *************                **********")
			for rows.Next() {

				var task taskStatus
				rows.Scan(&task.task, &task.status, &task.start_time, &task.end_time)
				fmt.Printf("%v                        %v              %v                 %v\n", task.task, task.status, task.start_time, task.end_time)
			}
		case 2:
			task := 0
			for {
				fmt.Println("Enter task id to complete a task: ")
				fmt.Scanln(&task)
				cnt := 0
				query := "select count(*) from taskStatus where task_id = ?"
				err := db.QueryRow(query, task).Scan(&cnt)
				if err != nil {

					log.Fatal("task retrieve failed: %v", err)
				}
				if cnt > 0 {
					query := "update taskStatus set status  = 'completed', end_time = ? where task_id = ?"
					_, err := db.Exec(query, time.Now().Format("02-01-06::03:04pm"), task)
					if err != nil {

						log.Fatal("update failed: %v", err)
					}
					break
				}

			}
		case 3:
			var newUsername string
			var newPassword string
			var newName string
			var newEmail string

			fmt.Println("Enter a new username: ")
			fmt.Scanln(&newUsername)
			fmt.Println("Enter a new password: ")
			fmt.Scanln(&newPassword)
			fmt.Println("Enter new name: ")
			scanner.Scan()
			newName = scanner.Text()
			fmt.Println("Enter a new email: ")
			fmt.Scanln(&newEmail)

			query := "update userInfo set username = ?, password = ?, name = ?, email = ? where username = ?"
			query2 := "update taskStatus set username = ? where username = ?"
			_, err := db.Exec(query, newUsername, newPassword, newName, newEmail, username)
			db.Exec(query2, newUsername, username)
			if err != nil {

				log.Fatal("update failed: %v", err)
			}
		}

	}
}

func adminLogin(db *sql.DB, scanner *bufio.Scanner) {

	var admin user

	for {

		fmt.Println("Enter your username: ")
		scanner.Scan()
		admin.username = scanner.Text()
		fmt.Println("Enter your password: ")
		scanner.Scan()
		admin.password = scanner.Text()

		rows, err := db.Query("select * from adminInfo")
		if err != nil {

			log.Fatal("data retrieve failed: %v", err)
		}
		var adminData user
		for rows.Next() {
			rows.Scan(&adminData.username, &adminData.password)
		}
		if admin.username == adminData.username && admin.password == adminData.password {

			fmt.Println("login successful!!")
			break
		} else {
			fmt.Println("Wrong credentials!!")
		}
	}

	choice := 0
Loop:
	for {
		fmt.Println("Your choices are: ")
		fmt.Println("*******************")
		fmt.Println("1.View employee info.")
		fmt.Println("2.Assign tasks.")
		fmt.Println("3.View task status.")
		fmt.Println("4.Exit.")
		fmt.Println("*******************")
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {

		case 1:
			viewEmployeeInfo(db)
		case 2:
			taskAssign(db, scanner)
		case 3:
			viewTaskStatus(db)
		case 4:
			break Loop
		}
	}

}
func viewEmployeeInfo(db *sql.DB) {

	rows, err := db.Query("select username, name, email from userInfo")
	if err != nil {

		log.Fatal("data retrieve failed: %v", err)
	}

	fmt.Println("Username              Name              Email")
	fmt.Println("*********            ******            ********")
	for rows.Next() {

		var userInfo user
		rows.Scan(&userInfo.username, &userInfo.name, &userInfo.email)
		fmt.Printf("%v          %v            %v \n", userInfo.username, userInfo.name, userInfo.email)
	}
}

func taskAssign(db *sql.DB, scanner *bufio.Scanner) {

	var user taskStatus
	for {
		fmt.Println("Enter username of employee: ")
		fmt.Scanln(&user.username)
		query := "select count(username) from userInfo where username = ?"
		cnt := 0
		err := db.QueryRow(query, user.username).Scan(&cnt)
		if err != nil {

			log.Fatal("query failed: %v", err)
		}
		if cnt > 0 {
			break
		} else {
			fmt.Println("Wrong credentials or username doesn't exist!!")
		}
	}

	for {
		fmt.Println("Enter a task for the employee: ")
		scanner.Scan()
		user.task = scanner.Text()
		query := "select count(task) from taskStatus where task = ?"
		cnt := 0
		err := db.QueryRow(query, user.task).Scan(&cnt)
		if err != nil {

			log.Fatal("query failed: %v", err)
		}
		if cnt == 0 {
			break
		} else {
			fmt.Println("Task already assigned to another employee!!")
		}
	}

	for {
		fmt.Println("Enter an unique task id: ")
		fmt.Scanln(&user.task_id)
		query := "select count(*) from taskStatus where task_id = ?"
		cnt := 0
		err := db.QueryRow(query, user.task_id).Scan(&cnt)
		if err != nil {

			log.Fatal("query failed: %v", err)
		}
		if cnt == 0 {
			break
		} else {
			fmt.Println("Task id already assigned to another employee!!")
		}
	}

	currentTime := time.Now().Format("02-01-06::03:04pm")
	var endTime string
	query := "insert into taskStatus values(?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, user.username, user.task, "active", currentTime, endTime, user.task_id)
	if err != nil {
		log.Fatal("database insertion failed: %v", err)
	}

}
func viewTaskStatus(db *sql.DB) {

	rows, err := db.Query("select * from taskStatus")
	if err != nil {

		log.Fatal("data retrieve failed: %v", err)
	}
	fmt.Println("(username) - (task) - (status) - (stat_time) - (end_time) - (task_id)")
	for rows.Next() {

		var task taskStatus
		rows.Scan(&task.username, &task.task, &task.status, &task.start_time, &task.end_time, &task.task_id)
		fmt.Printf("%v - %v - %v - %v - %v - %v\n", task.username, task.task, task.status, task.start_time, task.end_time, task.task_id)
	}

}

func databaseConnection(choice int, scanner *bufio.Scanner) {

	dns := "go:admin21@tcp(127.0.0.1:3306)/taskManagement"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("database open unsuccessful: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("database connection failed: %v", err)
	}
	fmt.Println("database connection successful!!")

	switch choice {

	case 1:
		adminLogin(db, scanner)
	case 2:
		userLogin(db, scanner)
	case 3:
		userRegistration(db)

	}
}
