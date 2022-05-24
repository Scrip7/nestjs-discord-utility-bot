Do not abuse `array.forEach()`. `forEach` is typically not the right method to use, and there's almost always a better method you could be using instead.

`forEach()` is **not** for:
‣ making changes to each item in an array or transforming it (use `map()` instead)
‣ removing certain items or cherry-picking items (use `filter()` instead)
‣ adding up, collating, or reducing values into a single value (use `reduce()` instead)
‣ anything asynchronous (use `map()`+`Promise.all()` or `for-of` instead)
‣ find out about more array methods at MDN: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array

`forEach()` is for:
‣ calling an external side-effect function on each item in the array which has no effect on the array or the local context of the function
‣ nothing else