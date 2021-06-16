# Table of Contents

- [Table of Contents](#table-of-contents)
- [Challenges](#challenges)
  - [1. Queue from Stacks](#1-queue-from-stacks)
  - [2. Linked List Shuffle](#2-linked-list-shuffle)
  - [3. Linked List Midpoint](#3-linked-list-midpoint)
  - [4. Linked List Loops](#4-linked-list-loops)
  - [5. Linked List `fromLast(n)`](#5-linked-list-fromlastn)
  - [6. Tree Elements per Level](#6-tree-elements-per-level)
  - [7. BST Contains](#7-bst-contains)
  - [8. BST Traversal](#8-bst-traversal)
  - [9. BST Validation](#9-bst-validation)

# Challenges

## 1. Queue from Stacks

simulate a queue's behavior, using stacks

prerequisite(s):
- queue
- stack

## 2. Linked List Shuffle

shuffle a doubly linked list

prerequisite(s):
- singly linked list
- doubly linked list

## 3. Linked List Midpoint

find the midpoint (middle node) in a __singly__ linked list, **but**
- no using a counter variable
- no using the size of the list
- only iterate through the list once at maximum

also
- if the length of list is even, there'd be 2 items in the middle, we'll return the first one (the one closer to the list's head)

prerequisite(s):
- singly linked list
- doubly linked list


## 4. Linked List Loops

determine if a loop in a singly linked list exists or not

prerequisite(s):
- singly linked list
- doubly linked list

## 5. Linked List `fromLast(n)`

make a function that when called with an integer param, will return the node at that index, starting from the last node, **but**
- no using the size method
- must use singly linked list

prerequisite(s):
- singly linked list
- doubly linked list

## 6. Tree Elements per Level

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

## 7. BST Contains

given a binary search tree, make a function that checks if a node with a certain data/value/key exists

prerequisite(s):
- binary search trees

## 8. BST Traversal

implement bst preorder, inorder, and postorder traversals, **and**
- must be able to pass a function to be called at every node, not just print

prerequisite(s):
- binary search trees

## 9. BST Validation

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
