package main

import (
	"fmt"
)

func main() {
	// pointer

	// firstName is going to be a pointer to string value
	var firstName *string
	fmt.Println(firstName) // Output is <nil>. Why? Because it is an empty pointer and it is uninitialized as of now.

	// var firstName2 *string

	// firstName2 = "Arthur" // trying to assign a string to the firstName2

	// fmt.Println(firstName2) // Will not work since we are assigning String to a pointer data type

	// var firstName3 *string

	// We should try dereferencing operator. to dereference a pointer,
	// you're basically saying I've got this pointer that's pointing at some data.
	// I want you to reach through the pointer, grab that data, and give that back to me.
	// This is called dereferencing.

	// *firstName3 = "Arthur" // for now this pointer is uninitialized.

	// Even this will not work, we get another error, segmentation violation and code panic's and exits.
	// And the reason for this is that we're actually trying to assign the string Arthur to an uninitialized
	// pointer. So since the pointer firstName isn't pointing to anything, it's uninitialized, Go is not
	// going to allow us to assign a value to that.

	// fmt.Println(firstName3)

	// Now to fix the above error
	// We need to use the new() method of Go to initialize the pointer
	var firstName4 *string = new(string)
	*firstName4 = "Arthur"
	fmt.Println(firstName4) // This will give the memory address, this is the value that pointer is holding.
	// And what value is the pointer holding? Well, it's holding the address in memory that's holding the
	// string Arthur for us.
	fmt.Println(*firstName4)

	firstName5 := "Tomchi"
	fmt.Println(firstName5)

	// Go allows us to create a pointer that is pointing to that variable
	ptr := &firstName5
	fmt.Println(ptr, *ptr) // printing the pointer and the value that it stores

	firstName5 = "Tom"
	fmt.Println(ptr, *ptr)

}
