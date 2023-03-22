package block

import "fmt"

func PrintBlockData(blockData BlockData) {
	fmt.Println("[-] Block Data")

	fmt.Println("[+] From: ", blockData["from"])
	fmt.Println("[+] To: ", blockData["to"])
	fmt.Println("[+] Amount: ", blockData["amount"])
}
