# Challenge: Decode A Message

A message has been encoded as a text stream that is to be read character by character. The stream contains a series of comma-delimited integers, each a positive number capable of being represented by a Go `int`. However, the character represeneted by a particular integer depends on the current decoding mode. There are three modes:
uppercase, lowercase, and punctuation.

In uppercase mode, each integer represent an uppercase letter: The integer modulo 27 indicated the letter of the alphabet (where 1 = A and so on). So an input value of 143 iin uppercase mode would yield the letter H because 143 modulo 27 is 8 and H is the eight letter in the alphabet.

The lowercase mode works the same but with lowercase letters; the remainder of dividing the integer by 27 represents the lowercase letter (1 = a and so on). So an input value of 56 in lowercase mode would yield the letter b because 56 modulo 27 is 2 and b is the second letter in the alphabet.

In punctuation mode, the integer is instead considered modulo 9, with the interprenetation given by the following Table:

| Number | Symbol  |
|--------|---------|
| 1      | !       |
| 2      | ?       |
| 3      | ,       |
| 4      | .       |
| 5      | (space) |
| 6      | ;       |
| 7      | "       |
| 8      | '       |

So 19 would yield an exclamation point because 19 modulo 9 is 1.

At the beginning of each message, the decoding mode is uppercase letters. Each time the modulo operation (by 27 or 9, depending on mode) results in 0, the decoding mode switches. If the current mode is uppercase, the mode switches to lowercase letters. If the current mode is lowercase, the mode switches to punctuation, and if it is punctuation, it switches back to uppercase.
