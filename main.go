package main

import "fmt"

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

}
