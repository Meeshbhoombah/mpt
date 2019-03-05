# mpt
Naive implementation of Modfied Merkle Patricia Trie.

## Structure
The following (`"key"`, `"value"`) pairs...
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

## References
1. [Ethereum Wiki: Patricia Trie](https://github.com/ethereum/wiki/wiki/Patricia-Tree)
1. [Ethereum Yellowpaper: Appendix D](https://ethereum.github.io/yellowpaper/paper.pdf)
1. [Khun Sir's: A Brief Analysis of MPT...](https://medium.com/vitelabs/a-brief-analysis-of-mpt-and-its-application-on-ethereum-c14c3de223f5)
1. [Timothy McCallum's Diving Into...](https://medium.com/cybermiles/diving-into-ethereums-world-state-c893102030ed)
1. [Nevous System's Clojure EVM](https://nervous.io/clojure/crypto/2018/04/04/clojure-evm-iii/)
1. [Understanding the Ethereum Trie](https://easythereentropy.wordpress.com/2014/06/04/understanding-the-ethereum-trie/)
1. [Ethereum Wiki: RLP](https://github.com/ethereum/wiki/wiki/%5BEnglish%5D-RLP)
1. [Vitalik Buterin's Merkling in Ethereum](https://blog.ethereum.org/2015/11/15/merkling-in-ethereum/)

### Implementations
1. [go-ethereum](https://github.com/ethereum/go-ethereum)
1. [begmaroman's](https://github.com/begmaroman/mpt)
1. [ebuchman's](https://github.com/ebuchman/understanding_ethereum_trie)
