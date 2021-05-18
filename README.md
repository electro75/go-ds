## Data Structures in Golang

- This repo contains the code written while learning about the various data structures in golang.

### Structs

- Similar to "objects" in JS.
- All structs have a name and different properties.
- Are used to create a "custom type" that only exists within the program
- Example : 
```
    type fruit struct {
        name string
        taste string
    }

    // I
    banana := fruit{"banana", "sweet"} // in this form of definition the order of values matters
    // name ==> "banana"
    // tast ==> "good"

    // II
    avocado := fruit{name: "avocado", tast: "unknown"}

    // III
    var mango fruit
    mango.name = "mango"
    mango.taste = "sweet"
``` 