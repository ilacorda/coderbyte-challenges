package main

import (
	"fmt"
)

func main() {
	testStr := "hellosannasmith"
	fmt.Println("The longest palindromic substring is:", SearchingChallenge(testStr))

	word := "I love !! dogs"
	fmt.Println("The longest word", LongestWord(word))

	fmt.Println("Singly Linked List Challenge")

	var singlyLinkedList SinglyLinkedList

	values := []string{"First", "Second", "Third"}
	for _, value := range values {
		singlyLinkedList.Append(value)
	}

	fmt.Println("The head of the linked list is:", singlyLinkedList.Head.Value)
	fmt.Println("the tail of the  linked list is ", singlyLinkedList.Tail.Value)

	//Codeland Username Validation
	fmt.Println("Codeland Username Validation")
	username := "u__hello_world123"
	fmt.Println("The username is valid:", CodelandUsernameValidation(username))

	invalidUsername := "2hello_world123_"
	fmt.Println("The username is valid:", CodelandUsernameValidation(invalidUsername))

	// reverse string
	fmt.Println("Reverse String")
	str := "hello"
	fmt.Println("The reversed string is:", FirstReverse(str))

	// fizzbuzz challenge
	fmt.Println("FizzBuzz Challenge")
	input := 15
	fmt.Println(FizzBuzz(input))

	// question marks
	fmt.Println("Question Marks")
	value := "14???4"
	fmt.Println(QuestionsMarks(value))

	// Letter changes
	fmt.Println("Letter Changes")
	providedStr := "hello*3"
	fmt.Println(LetterChanges(providedStr))

	// Simple Adding
	fmt.Println("Simple Adding")
	param := 12
	fmt.Println(SimpleAdding(param))

	// Number Reverse
	fmt.Println("Number Reverse")
	numbers := "1 2 3"
	fmt.Println(NumberReverse(numbers))

	// SimpleSymbols
	fmt.Println("Simple Symbol")
	symbol := "*&bla++" // this will return false
	fmt.Println(SimpleSymbols(symbol))

}
