**Type aliases vs. Interfaces**

Both type aliases and interfaces can be used to describe objects and are largely interchangeable in most cases:
```typescript
type TypeAlias = {
  foo: number
}

interface Interface {
  foo: number
}
```

_However, there are some subtle differences between the two:_

**1.**  Interfaces have declaration merging, while type aliases do not. This means that you can later add properties to an interface, which is not possible with type aliases. This is particularly useful for public types for a library, where the consumers of the library may need to modify the types. 
**2.**  Type aliases have implicit index signatures and can be assigned to types with compatible index signatures, while the equivalent interfaces cannot. 
**3.**  Properties of interfaces can reference the interface itself using this; the same isn't possible with type aliases.
**4.**  Type aliases can create mapped types, while interfaces cannot.
**5.**  In addition to describing objects, type aliases can be used to create aliases for primitives, unions, tuples, and anything else you could think of, while interfaces can only describe objects.