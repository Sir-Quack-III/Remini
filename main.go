package main

import (
	"Remini/crypt"
	"fmt"
	"strings"
)

func main() {
	var (
		inp string
		iar []string
		//pt  string
		//psd string
		//eps string
	)

	//fmt.Print("Enter Text: ")
	//fmt.Scanln(&pt)
	//
	//fmt.Print("Enter Password: ")
	//fmt.Scanln(&psd)

	//pt = "hi my name is isaac lopez"
	//psd = "mysupersecurepassword"

	fmt.Println(`________________________________
 ___              _        _ 
| _ \ ___  _ __  (_) _ _  (_)
|   // -_)| '  \ | || ' \ | |
|_|_\\___||_|_|_||_||_||_||_|
	  Password Manager
-------------------------------`)
	for {
		fmt.Print("> ")
		fmt.Scanln(&inp)

		iar = strings.Split(inp, " ")
		fmt.Println(inp)

		if inp == "exit" {
			break
		}

		switch iar[0] {
		case "encrypt":
			fmt.Println(crypt.EncryptAES(iar[1], crypt.EncryptSHA256(iar[2])))
		}
	}
}
