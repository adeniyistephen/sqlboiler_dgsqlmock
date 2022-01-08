package main

//go:generate sqlboiler --wipe mysql

import "mysql/example/user"

func main() {
	// Open DB instance and run listed functions
	user.OpenDB()
}