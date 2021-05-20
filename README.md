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

### Maps
- Very similar to `structs` but also different. The differences are:
    - Keys of a map should be of the same type, so should the values. In a struct the values can be of completely different types.
    - values cannot be accessed using the dot notation
    - In a map, all the keys are indexed which makes iteration over a map possible. In Struct keys are not indexed so iteration over keys is not possible.
    - A Map is __Reference Type__ whereas a struct is a __Value Type__.
    - Generally, a map is used to represent a collection of related properties. A Struct is used to represent one **thing** and its different properties.
    - In a Map, all the keys need not be fixed at compile time. All the keys are supposed to be fixed at compile time in a Struct.

- Declaring Maps :
```
// I
fruits := map[string]string {
    "banana": "good",
    "avocado": "unknown"
}

// II
var fruits map[string]string // empty map

// III
fruits := make(map[string]string) //empty map

// adding a new key
fruits["mango"] = "sweet"

// deleting a key
delete(fruits, "mango")
```


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