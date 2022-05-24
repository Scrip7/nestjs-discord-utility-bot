**Indentation is a key part to readable code**

If your code has extremely poor formatting, it can be hard to comprehend and overall becomes less readable.
You should always follow the following rules when indenting your code:

- A block is fixed amount of space (2 spaces, 4 spaces, 1 tab)
- Each new scope should be indented by a new block

Example of good indentation:
```javascript
function xyz() {
  if (true) {
    console.log("Foo")
    
    if (false) {
      console.log("Bar")
    }
  }
}
```