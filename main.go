package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
	this purpose of this function is to check for grammatical syntax, where nouns that beggining with vowels or

consonants have the right articles before them.
*/
func add_article(s []string) []string {
	for i, words := range s {
		switch words {
		case "a":
			if s[i+1][0] == 'a' || s[i+1][0] == 'e' || s[i+1][0] == 'i' || s[i+1][0] == 'o' || s[i+1][0] == 'u' || s[i+1][0] == 'h' {
				s[i] = "an"
			}
		case "A":
			if s[i+1][0] == 'a' || s[i+1][0] == 'e' || s[i+1][0] == 'i' || s[i+1][0] == 'o' || s[i+1][0] == 'u' || s[i+1][0] == 'h' {
				s[i] = "An"
			}
		default:
		}
	}
	return s
}

/*
	for the case of single quotes speechmarks, it is supposed to make sure qouted strings are apearing in the right

format with respect to the string which they are attached.
*/
func deal_w_apos(s []string) []string {
	count := 0
	for i, word := range s {
		if count == 0 && word == "'" {
			count++
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}
	for i, word := range s {
		if word == "'" {
			s[i-1] = s[i-1] + word
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

/*
this function is mainly focused on checking the string for the arrangment and order of the punctuation marks
and apply changes according to its position with reference to the next or previous word.
*/
func punctuate(s []string) []string {
	markers := []string{",", ".", "?", "!", ":", ";"}
	for i, word := range s {
		for _, mark := range markers {
			if strings.HasPrefix(word, mark) && !strings.HasSuffix(word, mark) {
				s[i] = word[1:]
				s[i-1] += mark
			}
		}
	}
	for i, word := range s {
		for _, mark := range markers {
			if strings.HasPrefix(word, mark) && strings.HasSuffix(word, mark) && i == len(s)-1 {
				s[i-1] += word
				s = s[:len(s)-1]
				break
			}
		}
	}
	for i, word := range s {
		for _, mark := range markers {
			if strings.HasPrefix(word, mark) && strings.HasSuffix(word, mark) && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}

// this function will capitalise the first letter of the string when required
func toTitle(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

/*
this function converts the given string from the value expressed in its format which is in hexadecimal value

	to is actual decimal value and convert back to string format
*/
func hex2str(s string) string {
	hex_num, err := strconv.ParseInt(s, 16, 64)
	if err == nil {
		num := strconv.FormatInt(hex_num, 10)
		s = num
	}

	return s
}

/*
this function converts the given string from the value expressed in its format which is in binary value

	to is actual decimal value and convert back to string format
*/
func bin2str(s string) string {
	bin_num, err := strconv.ParseInt(s, 2, 64)
	if err == nil {
		num := strconv.FormatInt(bin_num, 10)
		s = fmt.Sprint(num)
	}

	return s
}

/*
this is the important function in which i am calling all the functions made to manipulate strings. This is also the function that is
called in the main function of this program to conduct the manipulation required and return the required value.
*/
func recieveParam(str string) []string {
	s := strings.Fields(str)
	for i, word := range s {
		if word == "(hex)" {
			s[i-1] = hex2str(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}
		if word == "(bin)" {
			s[i-1] = bin2str(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}
		if word == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}
		if word == "(up," {
			new_byte := []byte(s[i+1])
			num, err := strconv.Atoi(string(new_byte[0]))
			if err == nil {
				for j := i - num; j < i; j++ {
					s[j] = strings.ToUpper(s[j])
				}
				s = append(s[:i], s[i+2:]...)
			}
		}
		if word == "(cap)" {
			s[i-1] = toTitle(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}
		if word == "(cap," {
			new_byte := []byte(s[i+1])
			num, err := strconv.Atoi(string(new_byte[0]))
			if err == nil {
				for j := i - num; j < i; j++ {
					s[j] = toTitle(s[j])
				}
				s = append(s[:i], s[i+2:]...)
			}
		}
		if word == "(low)" {
			s[i-1] = strings.ToLower(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}
		if word == "(low," {
			new_byte1 := []byte(s[i+1])
			num1, err := strconv.Atoi(string(new_byte1[0]))
			if err == nil {
				for k := i - num1; k < i; k++ {
					s[k] = strings.ToLower(s[k])
				}
				s = append(s[:i], s[i+2:]...)
			}
		}

	}
	for i := 0; i < len(s)-1; i++ {
		s = punctuate(s)
	}
	s = deal_w_apos(s)
	s = add_article(s)

	return s
}

/*
this function is focused on creating the file in which we want our output to be and write the contents to the created file.
the contents is that which will be the result of the manipulation done on other functions.
*/
func writeToFile(filename, contents string) error {
	theFile, err := os.Create(filename)
	if err != nil {
		fmt.Print(err)
	}
	defer theFile.Close()

	_, err = theFile.WriteString(contents)
	if err != nil {
		fmt.Print(err)
	}
	return err
}

/*
	this main involves extracting arguments from the command line and operate on them in a specified manner
	and display necessary error prompts if the criteria laid is not adhered to. the objective is to open a file to read its contents

which is in the form of a string content and manipulates it  according to a set formating rules to display the corrected text in an output file
that is created and the resulting string after manipulation is written to the output
*/
func main() {
	args := os.Args[1:]
	if len(os.Args) < 3 {
		fmt.Println("not enough arguments")
	}
	filename1 := args[0]
	if !strings.HasSuffix(filename1, ".txt") {
		fmt.Println("Error: File must have a .txt extension")
		os.Exit(1)
	}
	opendFile, err := os.Open(filename1)
	if err != nil {
		log.Fatal(err)
	}
	defer opendFile.Close()

	scanner := bufio.NewScanner(opendFile)

	line1 := ""
	for scanner.Scan() {
		line1 = scanner.Text() + " "
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	content := recieveParam(line1)

	filename := args[1]
	finalStr := strings.Join(content, " ") + "\n"

	err = writeToFile(filename, finalStr)
	if err != nil {
		fmt.Print(err)
		// }
		// doc := os.WriteFile(args[1], []byte(finalStr), 0o644)
		// if doc != nil {
		// 	panic(doc)
		// }
	}
}
