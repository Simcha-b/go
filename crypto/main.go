package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("WELCOME TO THE ENIGMA!", "","red", true)
	myFigure.Print()
	fmt.Println("This is a simple Caesar cipher implementation.")
	fmt.Println("You can encrypt and decrypt text using a key.")
	for {
		fmt.Println("Chose what you want to do:")
		fmt.Println("1. Encrypt text")
		fmt.Println("2. Decrypt text with a key")
		fmt.Println("3. Decrypt text without a key")
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
			encryptedText, nomMove := encryptText(text)
			writeToFile(encryptedText)
			fmt.Printf("Encrypted text with key: %d\n", nomMove)
			fmt.Println("Encrypted text is written to output.txt")
			fmt.Println("Encrypt successful!!")
		case 2:
			fmt.Println("You chose to decrypt text")
			fmt.Println("Please enter the key (number of positions to shift): ")
			// Read the key from user input
			var kay int
			fmt.Scanf("%d", &kay)
			bufio.NewReader(os.Stdin).ReadString('\n')
			text := readFromFile()
			if text == "" {
				fmt.Println("No text found in input.txt, please provide input.")
				return
			}
			if kay < 1 || kay > 25 {
				fmt.Println("Invalid key, please enter a number between 1 and 25")
				return
			}
			// Create the mapping based on the key
			kayMap := createKayMap(kay)
			// Decrypt the text using the same mapping
			decryptedText := decryptText(text, kayMap)
			writeToFile(decryptedText)
			fmt.Println("Decrypted successful!!")

		case 3:
			decryptTextWithoutKay(readFromFile())
			fmt.Println("Decrypted successful!!, check the results")
		default:
			fmt.Println("Invalid choice, please choose 1 or 2")
			return
		}
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
func encryptText(text string) (string, int) {
	numMove := rand.Intn(26) + 1
	fmt.Printf("Num: %d\n", numMove)
	kayMap := createKayMap(numMove)
	encrypted := []rune{}
	for _, char := range text {
		if val, ok := kayMap[char]; ok {
			encrypted = append(encrypted, val)
		} else {
			encrypted = append(encrypted, char)
		}
	}
	return string(encrypted), numMove
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
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(text + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	writer.Flush()
	fmt.Println("Text appended to output.txt")
}

func decryptTextWithoutKay(text string) {

	wg := sync.WaitGroup{}

	mt := sync.Mutex{}

	for i := 1; i < 26; i++ {
		wg.Add(1)

		go func(i int) {

			defer wg.Done()

			tempText := decryptText(text, createKayMap(i))

			mt.Lock()
			writeToFile(fmt.Sprintf("\n==== Try #%d ====\nShift: %d\nDecrypted Text:\n%s\n====================", i, i, tempText))
			mt.Unlock()

			log.Printf("Try #%d with shift %d:\n", i, i)
			fmt.Println(tempText)
		}(i)

	}
	wg.Wait()
}


