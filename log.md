## July 12, 2019
There's a lot of background to this Trie, which I may uncover one day, but 
right now I want to get a baseline Trie deploying with an empty node.

**When the contract is deployed..**
[ ] An empty `Node` is instantiated with the following:
```
Key = ""
Value = _
Children = []
```

Resulting in a **root**.

[ ] Hash the **root**, resulting in a **node reference hash**
[ ] Store a mapping of the **node reference hash** pointing to the **root**
[ ] Emit an event which contains the **node reference hash**

## July 13, 2019
Last I approached building this, I ran into difficulties with "hashing a 
struct" - which I'm beginning to feel, is a misconception I have developed from
years of programming in more abstract languages (Python, Javascript). 

I think my primary issue here stems from my misunderstanding of what a struct is
in Rust. When I was working with Golang I noticed there was no sense of order to
the fields in a struct. The only means was to have the code analyze itself, 
using its field labels to retrive the particular piece of data for a particular
instantiation of the struct. Would later find out this would be known as reflection.

Perhaps a struct is an abstraction made available only through the language.

After doing some poking around this appears to be true:
- [Struct Memory Layout in (C)](https://stackoverflow.com/questions/2748995/struct-memory-layout-in-c)
- [Data structure alignment](https://en.wikipedia.org/wiki/Data_structure_alignment)

