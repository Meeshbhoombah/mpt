# mpt
Naive implementation of Modfied Merkle Patricia Trie.

## Structure
The following (`"key"`, `"value"`) pairs:
```
("a711355", "dog")
("a77d337", "cat")
("a7f9365", "frog")
("a77d397", "bat")
```

Would result in:
```
|- "a7"
    |-"0"
    |-"1"
        |-"1355"
            |-"dog"
    |-"2"
    |-"3"
    |-"4"
    |-"5"
    |-"6"
    |-"7"
        |-"d3"
            |-"0"
            |-"1"
            |-"2"
            |-"3"
                |-"7"
                    |-"cat"
            |-"4"
            |-"5"
            |-"6"
            |-"7"
            |-"8"
            |-"9"
                |-"7"
                    |-"bat"
            |-"a"
            |-"b"
            |-"c"
            |-"d"
            |-"e"
            |-"f"
            |-"value"
    |-"8"
    |-"9"
    |-"a"
    |-"b"
    |-"c"
    |-"d"
    |-"e"
    |-"f"
        |-"9365"
    |-"value"

```
