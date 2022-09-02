# Exercise 4-9

Imagine a linked list where instead of the node storing a character, the node stores a digit: an int in the range 0-9. We could represent positive numbers of any size using such a linked list; the number 149, for example, would be a linked list in which the first node stores a 1, the second a 4, and the third and last a 9. Write a function *intToList* that takes an integer value and produces a linked list of this sort. **HINT**: You may find it easier to build the linked list backward, so if the value were 149, you would create the 9 node first.
