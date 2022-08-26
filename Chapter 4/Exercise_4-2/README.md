# Exercise 4-2

For our dynamically allocated strings, create a function *substring* that takes three parameters: an *arrayString*, a starting position integer, and an integer length of characters. The function returns a pointer to a new dynamically allocated string array. This string array contains the characters in the original string, starting at the specified position for the specified length. The original string is unaffected by the operation. So if the original string was *abcdefg*, the position was 3, and the length was 4, then the new string would contain *cdef*.

*NOTE*: Since Go has built-in *slices* that grow and shrink dynamically I will write this challenge using slices instead of C++ LinkedLists or Vectors.
