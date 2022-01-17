# Table of Contents

- [Table of Contents](#table-of-contents)
- [Challenges](#challenges)
  - [Algorithms](#algorithms)
    - [1. Palindromes](#1-palindromes)
    - [2. Anagrams](#2-anagrams)
    - [3. Swap integers without extra variable](#3-swap-integers-without-extra-variable)
    - [4. Big number subtraction](#4-big-number-subtraction)
    - [5. Reverse every word in sentence](#5-reverse-every-word-in-sentence)
    - [6. Regex simple match and extract](#6-regex-simple-match-and-extract)
    - [7. Max difference of 2 elements in array](#7-max-difference-of-2-elements-in-array)
    - [8. Power with recursion](#8-power-with-recursion)
    - [9. Palindrome number](#9-palindrome-number)
    - [10. String compression](#10-string-compression)
  - [Data structures](#data-structures)
    - [1. Queue from Stacks](#1-queue-from-stacks)
    - [2. Linked Lists](#2-linked-lists)
      - [2.1 Linked List Shuffle](#21-linked-list-shuffle)
      - [2.2 Linked List Midpoint](#22-linked-list-midpoint)
      - [2.3 Linked List Loops](#23-linked-list-loops)
      - [2.4 Linked List `fromLast(n)`](#24-linked-list-fromlastn)
      - [2.5 Linked List Reverse](#25-linked-list-reverse)
    - [3. Tree Elements per Level](#3-tree-elements-per-level)
    - [4. Binary Search Trees](#4-binary-search-trees)
      - [4.1 BST Contains](#41-bst-contains)
      - [4.2 BST Traversal](#42-bst-traversal)
      - [4.3 BST Validation](#43-bst-validation)

# Challenges

## Algorithms

### 1. Palindromes

given a string, determine if it's a palindrome

e.g.
- `palindrome("madam")` is true
- `palindrome("madame")` is false
- `palindrome("car")` is false
- `palindrome("racecar")` is true

### 2. Anagrams

given 2 strings, determine if they're anagrams

special characters such as space, comma, period, are included, so
- `anagram("hello.", "hell.o")` is true
- `anagram("hello.", "hello"` is false

### 3. Swap integers without extra variable

Given 2 integers, swap the value of those two variables without using extra variable(s).

### 4. Big number subtraction

given 2 numbers as a string, with values bigger than any integer type

make a function that can subtract these 2 big numbers and return the result as string

e.g.
- `bigNumSubtract("11443333311111111100", "1144422222221111")` returns `"11442188888888889989"`

bonus points if other operations are supported, such as addition

### 5. Reverse every word in sentence

create a function to reverse each word in a sentence

but

- the order of the words is unaffected
- punctuations are unaffected

also

- sentences will only contain `[a-zA-Z0-9]` and optional punctuation at the end

e.g.

```
input = "Hello World!"
output = "olleH dlroW!"
```

### 6. Regex simple match and extract

Given a string `"http://tokopedia.com/discovery/get-this-string/new"` \
Extract "get-this-string" with a regex.

For example: \
Input = `"http://tokopedia.com/discovery/hohoho/new"` \
Expected Output = "hohoho"

### 7. Max difference of 2 elements in array

Given an integer array, find the maximum difference between two elements in it such that the smaller element appears before the larger element.

e.g.

- input `{ 2, 7, 9, 5, 1, 3, 5 }` will output `7`, because the pair is `{2,9}`, not `{9,1}`
  - because "the smaller element appears before the larger element"

also

- we don't care what the pair is, we just want the difference

### 8. Power with recursion

implement power function like `math.Pow` using recursion

### 9. Palindrome number

check if a number is palindrome

but

- you may not cast it to a string

### 10. String compression

create a function that compresses duplicate consecutive chars in a string
into a number indicating the count of that consecutive char

e.g.
```
abbcccddddqwe -> ab2c3d4qwe
aaaaaaaaaab -> a10b
```

## Data structures

### 1. Queue from Stacks

simulate a queue's behavior, using stacks

prerequisite(s):
- queue
- stack

### 2. Linked Lists

#### 2.1 Linked List Shuffle

shuffle a doubly linked list

prerequisite(s):
- singly linked list
- doubly linked list

#### 2.2 Linked List Midpoint

find the midpoint (middle node) in a __singly__ linked list, **but**
- no using a counter variable
- no using the size of the list
- only iterate through the list once at maximum

also
- if the length of list is even, there'd be 2 items in the middle, we'll return the first one (the one closer to the list's head)

prerequisite(s):
- singly linked list
- doubly linked list


#### 2.3 Linked List Loops

determine if a loop in a singly linked list exists or not

prerequisite(s):
- singly linked list
- doubly linked list

#### 2.4 Linked List `fromLast(n)`

make a function that when called with an integer param, will return the node at that index, starting from the last node, **but**
- no using the size method
- must use singly linked list

prerequisite(s):
- singly linked list
- doubly linked list

#### 2.5 Linked List Reverse

reverse a singly linked list

prerequisite(s):
- singly linked list

### 3. Tree Elements per Level

make a function that takes a node, and returns an array/slice of element counts per level, starting from that node. like:

```
    5
   /|\
  1 3 9
 /| | |
2 6 4 8
```

now if we pass the node with the value `5` to our function, it will output `[1,3,4]` because
- there's 1 element in the first level, the "root" node
- there're 3 elements in the second level, 1, 3, and 9
- and also 4 elements in the third level, 2, 6, 4, and 8

but:
- must use a regular tree

also:
- the node values don't matter

### 4. Binary Search Trees

#### 4.1 BST Contains

given a binary search tree, make a function that checks if a node with a certain data/value/key exists

prerequisite(s):
- binary search trees

#### 4.2 BST Traversal

implement bst preorder, inorder, and postorder traversals, **and**
- must be able to pass a function to be called at every node, not just print

prerequisite(s):
- binary search trees

#### 4.3 BST Validation

given a node, validate wether it is a valid BST or not, for example
```
    10
    /\
   5  20
    \
    15
```
is not a valid bst, because even though 15 is correctly on the right branch of 5, it shouldn't be on the left branch of 10

instead it should be on the right branch of 10 like this
```
    10
    /\
   5  20
     /
    15
```
now this is a valid bst

prerequisite(s):
- binary search trees
