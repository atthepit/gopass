package main

import (
	"bufio"
	"bytes"
	"code.google.com/p/gopass"

	"fmt"
	"github.com/atotto/clipboard"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mesmerismo/gopass/pass_generator"
	_ "github.com/ziutek/mymysql/godrv"
	"os"
)

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func create_user(user_name string) {
	var master_password string
	var err error
	is_secure := false

	fmt.Println("Nice to meet you, " + user_name + "!")
	for !is_secure {
		fmt.Println("Please, write a good master password so I can keep your passwords secure: ")
		master_password, err = gopass.GetPass("")
		check_err(err)

		is_secure, score := pass_generator.Check_password(master_password)

		if !is_secure {
			fmt.Println("Your password is not very secure...")
			fmt.Printf("It has only scored %d out of 100\n", score)
			fmt.Println("Do you want to use it anyway? (Y/n)")
			var response []byte
			bio := bufio.NewReader(os.Stdin)
			response, _, err = bio.ReadLine()
			check_err(err)

			if string(response) == "Y" {
				fmt.Println("OK! Let's continue then......")
				break
			}
		} else {
			break
		}
	}

	hashed_pass, salt := pass_generator.Hash_password(master_password, nil)

	pass_generator.Create_user(user_name, hashed_pass, salt)

	fmt.Println("Done!")
	fmt.Println("Now you can store some passwords executing \"gopass -n new_site_name -u " + user_name + "\"")
}

func check_user(user string) []byte {

	num, master, salt := pass_generator.Get_user(user)
	if num != 1 {
		fmt.Println("I don't know any user named " + user)
		os.Exit(0)
	}

	master_password, err := gopass.GetPass("Hi, " + user + "! Write your master password now: ")
	check_err(err)

	hashed_pass, _ := pass_generator.Hash_password(master_password, salt)

	if !bytes.Equal(hashed_pass, master) {
		fmt.Println("Sorry but that password it's incorrect. Try again!")
		os.Exit(0)
	}
	fmt.Println("Correct!")
	return master
}

func create_site(site, user string) {

	master := check_user(user)

	fmt.Println("Creating password......")
	var new_pass string
	is_secure := false
	for !is_secure {
		new_pass = pass_generator.Generate_random_password(16)
		is_secure, _ = pass_generator.Check_password(new_pass)
	}

	err := clipboard.WriteAll(new_pass)
	check_err(err)

	encrypted_pass := pass_generator.Encrypt(new_pass, master)
	fmt.Println(" Done!")

	pass_generator.Create_site(site, user, encrypted_pass)

	fmt.Println("Your password for " + site + " has been created and copied to your clipboard!")
}

func get_site_password(master []byte, site, user string) []byte {

	pass := pass_generator.Get_password(site, user)
	decrypted_pass := pass_generator.Decrypt(pass, master)

	return decrypted_pass
}

func get_site(site, user string) {
	master := check_user(user)

	fmt.Println("Getting your password........")
	decrypted_pass := get_site_password(master, site, user)
	fmt.Println(" Done!")
	err := clipboard.WriteAll(string(decrypted_pass))
	check_err(err)

	fmt.Println("Your password for " + site + " has been copied to your clipboard!")
}

func list_sites(user string) {
	check_user(user)
	names := pass_generator.List_sites(user)
	if len(names) > 0 {
		fmt.Println("I know your password for this sites: ")
		for i := 0; i < len(names); i++ {
			fmt.Println("- " + names[i])
		}
	} else {
		fmt.Println("Sorry but I don't know any of your passwords...")
	}

}

func delete_site(site, user string) {
	check_user(user)

	fmt.Println("Are you sure you want to delete your " + site + " password? (Y/n)")
	fmt.Println("(you should change the password in that site first!)")
	var response []byte
	bio := bufio.NewReader(os.Stdin)
	response, _, err := bio.ReadLine()
	check_err(err)

	if string(response) != "Y" {
		os.Exit(1)
	}

	pass_generator.Delete_site(site, user)

	fmt.Println("Done!")
}

func reset_site_password(site, user string) {
	master := check_user(user)

	decrypted_pass := get_site_password(master, site, user)

	err := clipboard.WriteAll(string(decrypted_pass))
	check_err(err)

	fmt.Println("The previous password has been copied to your clipboard.")
	fmt.Println("Be sure you change it in the site before continuing!")
	fmt.Println("Continue? (Y/n)")
	var response []byte
	bio := bufio.NewReader(os.Stdin)
	response, _, err = bio.ReadLine()
	check_err(err)

	if string(response) != "Y" {
		os.Exit(1)
	}

	fmt.Print("Creating new password........")
	var new_pass string
	is_secure := false
	for !is_secure {
		new_pass = pass_generator.Generate_random_password(16)
		is_secure, _ = pass_generator.Check_password(new_pass)
	}
	err = clipboard.WriteAll(new_pass)
	check_err(err)
	fmt.Println(" Done!")

	encrypted_pass := pass_generator.Encrypt(new_pass, master)

	pass_generator.Update_password(site, user, encrypted_pass)

	fmt.Println("Your new password for " + site + " has been created and copied to your clipboard!")
}

func delete_user(user string) {
	check_user(user)
	fmt.Println("Remember: This action can not be undone")
	fmt.Println("Are you sure you want to delete your account? (Y/n)")
	var response []byte
	bio := bufio.NewReader(os.Stdin)
	response, _, err := bio.ReadLine()
	check_err(err)

	if string(response) != "Y" {
		fmt.Println("Yay! Thanks for staying! :D")
		os.Exit(1)
	}

	fmt.Println("Goodbye, " + user + ". I will miss you :_(")
	pass_generator.Delete_user(user)
}

func display_help() {
	fmt.Println("This is what I can do:")
	fmt.Println("-Create a new user")
	fmt.Println("\tgopass -u new_user_name")
	fmt.Println("-Store a new password for a user")
	fmt.Println("\tgopass -n new_site_name -u user_name")
	fmt.Println("-Get a user's site password")
	fmt.Println("\tgopass -s site_name -u user_name")
	fmt.Println("-Reset a user's site password ")
	fmt.Println("\tgopass -r site_name -u user_name")
	fmt.Println("-Forget all I know about a user")
	fmt.Println("\tgopass -d user_name")
	fmt.Println("-Destroy a user's site password")
	fmt.Println("\tgopass -d site_name -u user_name")
	fmt.Println("-List all I know about a user")
	fmt.Println("\tgopass -l user_name")
}

func main() {
	pass_generator.Create_DB()
	var option, param string
	num_params := len(os.Args)

	if num_params >= 2 {
		option = os.Args[1]
		if num_params >= 3 {
			param = os.Args[2]
		}
	}

	switch option {
	case "-u":
		if param != "" {
			create_user(param)
		} else {
			fmt.Println("You must specify who you are: gopass -u username")
		}
		break
	case "-n":
		site_name := param
		if num_params == 5 {
			if os.Args[3] == "-u" {
				user_name := os.Args[4]
				create_site(site_name, user_name)
			} else {
				fmt.Println("You must tell me who you are: gopass -n sitename -u username")
				os.Exit(0)
			}
		} else {
			fmt.Println("I only need this info: gopass -n sitename -u username")
			os.Exit(0)
		}
		break
	case "-s":
		site_name := param
		if num_params == 5 {
			if os.Args[3] == "-u" {
				user_name := os.Args[4]
				get_site(site_name, user_name)
			} else {
				fmt.Println("You must tell me who you are: gopass -s sitename -u username")
				os.Exit(0)
			}
		} else {
			fmt.Println("I only need this info: gopass -s sitename -u username")
			os.Exit(0)
		}
		break
	case "-r":
		site_name := param
		if num_params == 5 {
			if os.Args[3] == "-u" {
				user_name := os.Args[4]
				reset_site_password(site_name, user_name)
			} else {
				fmt.Println("You must tell me who you are: gopass -r sitename -u username")
				os.Exit(0)
			}
		} else {
			fmt.Println("I only need this info: gopass -r sitename -u username")
			os.Exit(0)
		}
		break
	case "-l":
		if num_params == 3 {
			list_sites(os.Args[2])
		} else {
			fmt.Println("I only need this info: gopass -l username")
			os.Exit(0)
		}
		break
	case "-d":
		if num_params == 3 {
			user_name := param
			delete_user(user_name)
		} else {
			site_name := param
			if num_params == 5 {
				if os.Args[3] == "-u" {
					user_name := os.Args[4]
					delete_site(site_name, user_name)
				} else {
					fmt.Println("You must tell me who you are: gopass -d sitename -u username")
					os.Exit(0)
				}
			} else {
				fmt.Println("I only need this info: gopass -d sitename -u username")
				os.Exit(0)
			}
		}
		break
	case "-h":
		display_help()
		break
	default:
		fmt.Println("Sorry but I do not know how to do that :S")
		display_help()
		break
	}
}
