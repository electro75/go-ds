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

- #### Pointers
    - A __pointer__ stores the memory address of the variable it points at.
    - `*pointer` will give the value stored in the address
    - Example:
    ```
    fruit := "banana"
    fruitPointer := &fruit
    // fruitPointer ==> address of the value "banana"
    // *fruitPointer ==> "banana"
    ```
    - `*` in front of a pointer signifies the value whereas `*` in front of a type implies a __pointer to that type__ or type.

- Slices don't behave the same way as structs do with regards receiver functions as they are a **Reference Type**
    - A receiver function called with slice will modify the value of the slice.
    - The copy passed to the function will still refer to the underlying array.

## Types :

| Value         | Reference     |
| -----------   | ------------  |
| int           | slices        |
| float         | maps          |
| string        | channels      |
| bool          | pointers      |
| structs       | functions     |

* Value types need to be used with pointers for mutation with receiver functions
* Reference type can be mutated with the receiver functions