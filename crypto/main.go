package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the ENIGMA!!!!")
	moveNum := 3 
	kayMap := createKayMap(moveNum)
	fmt.Println("Chose what you want to do:")
	fmt.Println("1. Encrypt text")
	fmt.Println("2. Decrypt text")
	var choice int
	fmt.Scanf("%d", &choice)
	bufio.NewReader(os.Stdin).ReadString('\n')
	switch choice {
	case 1:
		fmt.Println("You chose to encrypt text")
		// text := readFromUser()
		text := readFromFile()
		if text == "" {
			fmt.Println("No text found in input.txt, please provide input.")
			return
		}
		encryptedText := encryptText(text, kayMap)
		writeToFile(encryptedText)
		fmt.Println("Encrypt successful!!")
	case 2:
		fmt.Println("You chose to decrypt text")
		// text := readFromUser()
		text := readFromFile()
		if text == "" {
			fmt.Println("No text found in input.txt, please provide input.")
			return
		}
		// Decrypt the text using the same mapping
		decryptedText := decryptText(text, kayMap)
		writeToFile(decryptedText)
		fmt.Println("Decrypted successful!!")
	default:
		fmt.Println("Invalid choice, please choose 1 or 2")
		return
	}
}

func readFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the text that you want to encrypt: ")
	input, _ := reader.ReadString('\n')
	// fmt.Println(input)
	return input
}

// createKayMap creates a mapping for the Caesar cipher based on the move number
func createKayMap(moveNum int) map[rune]rune {
	abcLower := []rune("abcdefghijklmnopqrstuvwxyz")
	abcUpper := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	mp := make(map[rune]rune)
	// Create a mapping for lowercase and uppercase letters
	for i := 0; i < len(abcLower); i++ {
		mp[abcLower[i]] = abcLower[(i+moveNum)%len(abcLower)]
	}
	// Create a mapping for uppercase letters
	for i := 0; i < len(abcUpper); i++ {
		mp[abcUpper[i]] = abcUpper[(i+moveNum)%len(abcUpper)]
	}

	return mp
}

// encryptText encrypts the input text using the provided mapping
func encryptText(text string, kayMap map[rune]rune) string {
	encrypted := []rune{}
	for _, char := range text {
		if val, ok := kayMap[char]; ok {
			encrypted = append(encrypted, val)
		} else {
			encrypted = append(encrypted, char)
		}
	}
	return string(encrypted)
}

// decryptText decrypts the input text using the provided mapping
func decryptText(text string, kayMap map[rune]rune) string {
	// Create a reverse mapping for decryption
	reverseMap := make(map[rune]rune)
	for k, v := range kayMap {
		reverseMap[v] = k
	}
	decrypted := []rune{}
	for _, char := range text {
		if val, ok := reverseMap[char]; ok {
			decrypted = append(decrypted, val)
		} else {
			decrypted = append(decrypted, char)
		}
	}
	return string(decrypted)
}

// readFromFile reads the content of input.txt and returns it as a string
func readFromFile() string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()
	res := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res += line + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return res
}

// writeToFile writes the given text to output.txt
func writeToFile(text string) {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	writer.Flush()
	fmt.Println("Output written to output.txt")
}


