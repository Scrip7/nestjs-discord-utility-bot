Consider using `fs/promises` instead of the older `fs` module in Node 14+
```typescript
import fs from "fs/promises"; // or use require

const contents = await fs.readFile("./foo.txt", "utf-8");
```
Benefits of using `fs/promises` over `fs`:
▫️ Promise-based API fits into other `async` code easier
▫️ All the benefits of using `async`/`await` over other asynchronous strategies
▫️ Easier logic to understand code flow than using the callback API
▫️ Does not incur the performance cost of blocking the thread like the `__Sync` functions present in old `fs`

Documentation: https://nodejs.org/api/fs.html#promises-api