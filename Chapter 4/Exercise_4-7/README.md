# Exercise 4-7

Following up on the previous exercise, implement the *concatenate* function. Note that if we make a call *concatenate(s1, s2)*, where both parameters are pointers to the first nodes of their respective linked lists, the function should create a copy of each of the nodes in *s2* and append them to the end of *s1*. That is, the function should not simply point the *next* field of the last node in the *s1*'s list to the first node of *s2*'s list.
