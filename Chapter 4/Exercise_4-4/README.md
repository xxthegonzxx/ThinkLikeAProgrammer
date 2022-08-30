# Exercise 4-4

Change the implementation of our strings such that *location[0]* in the array stores the size of the array (and therefore *location[1]* stores the first actual character in the string), rather than using a null-character terminator. Implement each of the three functions, *append*, *concatenate*, and *characterAt*, taking advantage of the stored size information whenever possible. Because we'll no longer be using the null-termination convention expected by the standard output stream, you'll need to write your own *output* function that loops through its string parameter, displaying the characters.

**NOTE**: *This is not idiomatic Go as this book was heavily influenced by C++. Some of it feels wonky and forced to comply with the requirements of the challenge.*
