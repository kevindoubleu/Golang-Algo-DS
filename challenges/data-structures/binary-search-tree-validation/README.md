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