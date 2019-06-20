# Speaking Notes
1. Hashing
1. Merkle Trees
1. Trie
1. Patricia Trie

## 40 Years
1. Cryptography
1. Computer Science
1. Economics
1. Game Theory

1 Year in 7.5 seconds

## Cryptographic Hash Function
1. Deterministic
`h(k)` is always `h(k)`

1. Quick Computation
1. Pre-Image Resistance
Given `h(k)` - impossible to determine `k`

1. Second Pre-Image Resistance
Given `k`, `h(k)`, and `h(v)` - impossible to determine `v`

1. Collision Resistance
`h(k)` is never `h(v)`

1. Reflective
Slight change to `k` is huge change in `h(k)`

## Merkle Tree
- Used in Bitcoin and blockchain projects
- Simplified Payment Verification

- Chunk = unit of information
- Seperate each chunk into into buckets (how = impl specefic)
- Hash each chunk -> bucket -> root 

### Merkle Proof
- Change ripples across entire Merkle
- Prove position and value
- Only need three pieces of data = (tx B, hash of tx A, root to Right)

## Trie
- Store words breaking down to the char
- Perhaps stored entire word in leaf to avoid recursive pop to assemble word
- 26 slots for each node

## Modified Merkle Patricia Trie
- Three types of node (abstract)
    + Branch
    + Extension
    + Leaf
- Implementation, only two types of nodes
    + One with 17 members
    + Another
