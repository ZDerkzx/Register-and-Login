package main

import (
	"fmt"
)

var Users = make(map[string]int)
var Logged bool = false
var ActualAccountName string
var ActualAccountPass int8

func changePassword() {
	if Logged {
		var newPassword int8
		var confirm string

		for key, value := range Users {
			if key == ActualAccountName && value == int(ActualAccountPass) {
				fmt.Println("-- Succes --")
				fmt.Println("Enter New Password")
				fmt.Scan(&newPassword)
				fmt.Println("¿Are You Sure? Y/N")
				fmt.Scan(&confirm)

				switch confirm {
				case "Y", "y":
					Users[ActualAccountName] = int(newPassword)
					ActualAccountPass = newPassword
					Logged = false
				case "N", "n":
					accountSettings()
				}
			}
		}
	}
}

func changeName() {
	if Logged {
		var newName string
		var confirm string

		for key, value := range Users {
			if key == ActualAccountName && value == int(ActualAccountPass) {
				fmt.Println("-- Succes --")
				fmt.Println("Enter New Name: ")
				fmt.Scan(&newName)
				fmt.Println("¿Are You Sure? Y/N")
				fmt.Scan(&confirm)
				switch confirm {
				case "Y", "y":
					delete(Users, ActualAccountName)
					ActualAccountName = newName
					Users[newName] = int(ActualAccountPass)
					fmt.Println("Account Name Now: ", ActualAccountName)
					fmt.Println("Password: ", ActualAccountPass)
					Logged = false
					return
				case "N", "n":
					accountSettings()
					return
				}
			}
		}
	} else {
		fmt.Println("You are not Logged")
		return
	}
}

func accountSettings() {
	if Logged {
		var option int8
		var menu = [3]string{
			"1.Change Name",
			"2.Change Password",
			"3.Logout",
		}
		fmt.Println("---_ Settings _---")
		fmt.Println(menu)
		fmt.Scan(&option)
		switch option {
		case 1:
			changeName()
		case 2:
			changePassword()
		case 3:
			Logged = false
			return
		}
	}
}

func addUser() {
	var username string
	var userpassword int8
	fmt.Println("Ingresa tu Username:")
	fmt.Scan(&username)
	fmt.Println("Ingresa tu contraseña:")
	fmt.Scan(&userpassword)
	Users[username] = int(userpassword)
	fmt.Println(Users)
	return
}

func loginUser() {
	var loginname string
	var loginpassword int8
	fmt.Println("-- Login --")
	fmt.Println("Ingrese su nombre")
	fmt.Scan(&loginname)
	fmt.Println("Ingrese su contraseña")
	fmt.Scan(&loginpassword)
	for llave, valor := range Users {
		if llave == loginname && valor == int(loginpassword) {
			fmt.Println("Has Ingresado a tu cuenta!")
			Logged = true
			ActualAccountName = loginname
			ActualAccountPass = loginpassword
			accountSettings()
		}
	}
}

var menu = [3]string{
	"1.Register",
	"2.Login",
	"3.Verify Logged",
}

func main() {
	var option int8
	for true {
		fmt.Print(menu)
		fmt.Scan(&option)
		switch option {
		case 1:
			addUser()
		case 2:
			loginUser()
		case 3:
			fmt.Println("Logged: ", Logged)
		}
	}

}
