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

prerequisite(s):
- trees