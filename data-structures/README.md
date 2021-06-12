# Table of Contents

- [Table of Contents](#table-of-contents)
- [Time complexities](#time-complexities)
  - [1. Queue and Stack](#1-queue-and-stack)
  - [2. Linked List](#2-linked-list)
    - [2.1 Linked List (Singly)](#21-linked-list-singly)
    - [2.2 Linked List (Doubly)](#22-linked-list-doubly)

# Time complexities

Table of time complexities for common operations in each DS implementations in this repo

## 1. Queue and Stack

| operation | complexity |
| --------- | ---------- |
| push      | O(1)       |
| pop       | O(1)       |
| peek      | O(1)       |
| lookup    | O(n)       |

## 2. Linked List

| operation                       | complexity |
| ------------------------------- | ---------- |
| get(index)                      | O(n)       |
| pop(index)                      | O(n)       |
| push(index)                     | O(n)       |
| size, if kept as an attribute   | O(1)       |
| size, if dynamically calculated | O(n)       |

### 2.1 Linked List (Singly)

| operation | complexity |
| --------- | ---------- |
| append    | **O(n)**   |
| prepend   | O(1)       |
| popFirst  | O(1)       |
| popLast   | **O(n)**   |

append, popLast: [improved in doubly](#2.2-linked-list-doubly)

### 2.2 Linked List (Doubly)

| operation | complexity |
| --------- | ---------- |
| append    | **O(1)**   |
| prepend   | O(1)       |
| popFirst  | O(1)       |
| popLast   | **O(1)**   |