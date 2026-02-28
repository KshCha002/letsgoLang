# Important pointers for topics:

1. Basic Language Fundamentals
Package declaration (package main)
Import statements and package management
Basic main() function
Comments (single-line )

2. Variables and Data Types
Variable declarations (var, short declaration :=)
Basic types: int, string, bool, float32, float64
Type inference and explicit typing
Zero values initialization

3. Arrays and Slices
Array declaration and initialization ([5]int, [3]int{1,21,3})
Slice creation ([]int{-11,22,3,13,54})
Slice operations: append(), slicing (slice[1:2])
make() for slice creation
len() function for length
Slice element removal techniques

4. Maps (Key-Value Pairs)
Map creation with make(map[int]string)
Map literal syntax (map[int]int{})
Adding/accessing map elements
delete() function (commented)
Iterating over maps with range

5. Control Flow
Conditionals: if-else statements
Switch statements: with case, default, fallthrough
Loops:
for loops with index
range loops for iteration
Loop control: continue, break
Labels and goto statements

6. Functions
Function declaration and definition
Multiple parameter functions (func k(v1 int, v2 int, v3 int))
Multiple return values ((int, string))
Variadic functions (func kki(v1 ...int))
Function calls and return value handling

7. Structs and Methods
Struct definition (type User struct, type struct1 struct)
Struct field access (exported vs unexported fields)
Struct initialization and literal syntax
Methods: Functions with receivers (func (s *struct1) Getstuct())
Pointer receivers for methods

8. Pointers
Pointer declaration (var ptr *int)
Address operator (&)
Dereference operator (*)
Pointer manipulation and value modification

9. Input/Output
Standard Output: fmt.Println(), fmt.Printf() with format specifiers
User Input: bufio.NewReader(), ReadString()
Format specifiers (%T, %d, %s, %g, %v)

10. String and Type Conversion
strconv.ParseInt() for string to integer conversion
strconv.ParseFloat() for string to float conversion
strings.TrimSpace() for string manipulation
Type casting and conversion

11. Time Package
time.Now() for current time
time.Date() for specific date creation
Time formatting with Format()
time.Local timezone

12. Defer Statements
defer keyword usage
LIFO (Last In, First Out) execution order
Defer in function cleanup

13. Random Number Generation
math/rand package
rand.Seed() for seeding
rand.Intn() for random integers

14. Go Modules and Workspace
go.mod files for module definition
go.work files for multi-module workspaces
Cross-package imports and usage


15. Error Handling Basics
Error variable handling (err variable)
Basic error checking patterns (if err != nil) I have covered all these basic topics , give me list of interview questions related to these and first tell me how likely these are to be asked in an interview

## Questions
Perfect 💯
I’ll answer **every question concisely but interview-ready**, then add **important extra pointers at the end**.

---

# 1️⃣ Basics & Package

### 1. Why is `package main` special?

Because only `package main` can produce an executable binary. It must contain `main()` function.

---

### 2. What happens if `main()` is missing?

Compilation error: program entry point missing.

---

### 3. Purpose of `import`?

To include external or standard library packages into your file.

---

### 4. Can a Go file have multiple packages?

No. One file → one package.

---

### 5. What is `init()`?

Special function executed automatically before `main()`. Used for initialization.

---

# 2️⃣ Variables & Types

### 1. Difference between `var x int = 5` and `x := 5`

* `var` → explicit declaration.
* `:=` → short declaration + type inference.
* `:=` only inside functions.

---

### 2. Where can you NOT use `:=`?

Outside functions (package level).

---

### 3. What are zero values?

Default values when variable declared but not initialized:

* int → 0
* string → ""
* bool → false
* slice/map/pointer → nil

---

### 4. Is Go statically typed?

Yes. Type is determined at compile time.

---

### 5. What if variable declared but not used?

Compilation error.

---

# 3️⃣ Arrays vs Slices

### 1. Difference?

* Array → fixed size, value type.
* Slice → dynamic size, reference type (wrapper over array).

---

### 2. Is slice passed by value or reference?

Passed by value, but underlying array reference is copied. So elements can be modified.

---

### 3. Length vs Capacity?

* len → number of elements.
* cap → total capacity before reallocation.

---

### 4. Append beyond capacity?

New underlying array created (usually doubled capacity).

---

### 5. How does `make()` work?

Allocates underlying array and returns slice.

Example:

```go
make([]int, 5, 10)
```

---

### 6. Remove element from slice?

```go
s = append(s[:i], s[i+1:]...)
```

---

### 7. Can slices be compared?

No (except to nil).

---

### Classic Output:

```go
a := []int{1,2,3}
b := a
b[0] = 100
```

Output: `[100 2 3]`

Because both share same underlying array.

---

# 4️⃣ Maps

### 1. Zero value?

nil

---

### 2. Access non-existing key?

Returns zero value of value type.

---

### 3. Check if key exists?

```go
v, ok := m[key]
```

---

### 4. Are maps thread-safe?

No.

---

### 5. Can map keys be slices?

No. Keys must be comparable types.

---

### 6. Read from nil map?

Returns zero value.

---

### 7. Write to nil map?

Runtime panic.

---

# 5️⃣ Control Flow

### 1. Does Go have while?

No. Use `for`.

---

### 2. How switch works?

Matches case, executes matching block.

---

### 3. Does switch auto-break?

Yes.

---

### 4. What is fallthrough?

Forces next case to execute.

---

### 5. Labels?

Used with break/continue/goto.

---

# 6️⃣ Functions

### 1. Why multiple return values?

For returning result + error.

---

### 2. Named return?

```go
func f() (x int) {
    x = 10
    return
}
```

---

### 3. Variadic?

```go
func sum(nums ...int)
```

---

### 4. Functions as arguments?

Yes (first-class functions).

---

### 5. Default parameters?

No.

---

# 7️⃣ Structs & Methods

### 1. Struct vs class?

Struct = data structure. No inheritance.

---

### 2. Method receiver?

Function attached to type.

---

### 3. Value vs Pointer receiver?

* Value → copy.
* Pointer → modifies original.

---

### 4. When use pointer receiver?

* Modify struct
* Avoid copying large struct
* Consistency

---

### 5. Exported fields?

Capital letter → accessible outside package.

Example:

```go
Name string  // exported
name string  // private
```

---

# 8️⃣ Pointers

### 1. Pointer arithmetic?

No (unlike C).

---

### 2. new() vs make()?

* new → allocates memory, returns pointer.
* make → initializes slice, map, channel.

---

### 3. Everything passed by value?

Yes.

---

### 4. Dereference nil pointer?

Runtime panic.

---

### 5. Why pointers with struct?

Avoid copying large structs.

---

# 9️⃣ Error Handling

### 1. Type of error?

Built-in interface:

```go
type error interface {
    Error() string
}
```

---

### 2. Why no exceptions?

Explicit error handling → cleaner, predictable flow.

---

### 3. Idiomatic error handling?

```go
if err != nil {
    return err
}
```

---

### 4. Custom error?

```go
errors.New("msg")
```

---

### 5. panic vs error?

* error → expected failure.
* panic → unrecoverable failure.

---

# 🔟 Defer

### 1. When executed?

After function returns.

---

### 2. LIFO?

Yes.

---

### 3. Output?

```go
for i := 0; i < 3; i++ {
    defer fmt.Println(i)
}
```

Output:

```
2
1
0
```

---

# 1️⃣1️⃣ String & Conversion

### 1. Conversion vs assertion?

* Conversion → change type.
* Assertion → extract concrete type from interface.

---

### 2. String to int?

Using `strconv.ParseInt()` from Go standard library

---

### 3. ParseInt fails?

Returns error.

---

### 4. Why TrimSpace?

Remove newline from user input.

---

# 1️⃣2️⃣ Time Package

From Go standard library

### 1. Format time?

```go
t.Format("2006-01-02")
```

---

### 2. Why 2006-01-02?

Reference date used by Go (Mon Jan 2 15:04:05 MST 2006).

---

# 1️⃣3️⃣ Random

From `math/rand`

### 1. Why Seed?

To generate different sequence.

---

### 2. No seed?

Same sequence every run.

---

# 1️⃣4️⃣ Go Modules

### 1. go.mod?

Defines module name and dependencies.

---

### 2. go.sum?

Dependency checksum verification.

---

### 3. How dependencies resolved?

Using semantic versioning.

---

### 4. go.work?

Used for multi-module workspace.

---

### 5. Semantic versioning?

MAJOR.MINOR.PATCH

---

## 🔥 IMPORTANT EXTRA POINTERS (VERY HIGH INTERVIEW VALUE)

These are commonly asked but you didn’t explicitly list:

---

###  ✅ 1. Difference between make and new (VERY COMMON)

* new → returns pointer
* make → initializes slice/map/channel

---

###  ✅ 2. nil vs empty slice

```go
var s []int       // nil slice
s := []int{}      // empty slice
```

len same (0), but nil check differs.

---

###  ✅ 3. Slice internal structure

Slice contains:

* pointer to array
* length
* capacity

---

###  ✅ 4. Value receiver can call pointer receiver?

Yes. Go automatically handles it.

---

###  ✅ 5. Map iteration order?

Random. Not guaranteed.

---

###  ✅ 6. Are arrays comparable?

Yes (if elements comparable).
Slices → not comparable.

---

###  ✅ 7. What causes memory leak in Go?

* Goroutines not closed
* Large slices referencing small part of big array

---

###  ✅ 8. Does Go have inheritance?

No. Uses composition.

---

## 🎯 Final Advice

Interviewer wants to see you confidently explain:

* Slice internals
* Map nil behavior
* Pointer receiver reasoning
* Error handling philosophy


