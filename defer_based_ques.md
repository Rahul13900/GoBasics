## 🃏 Golang `defer` Flashcards – Interview Edition

---

### 🔹 **Card 1: Basic Defer Execution**

**Q:** When does a deferred function run?
**A:** After the surrounding function returns, in **LIFO** (last-in, first-out) order.

---

### 🔹 **Card 2: Argument Evaluation Timing**

**Q:** When are deferred function arguments evaluated?
**A:** At the time the **defer statement is declared**, not at execution.

**Example:**

```go
x := 5
defer fmt.Println(x)
x = 10
```

**Output:** `5`

---

### 🔹 **Card 3: LIFO Execution Order**

**Q:** What is the order of deferred function execution?

**Example:**

```go
defer fmt.Println("first")
defer fmt.Println("second")
defer fmt.Println("third")
```

**Output:**

```
third
second
first
```

---

### 🔹 **Card 4: Named Return Value Modification**

**Q:** Can defer modify a named return value?

**Example:**

```go
func test() (result int) {
    defer func() { result++ }()
    return 5
}
```

**Output:** `6`

---

### 🔹 **Card 5: Effect on Unnamed Return Values**

**Q:** Will defer change the return value if return is not named?

**Example:**

```go
func foo() int {
    x := 5
    defer func() { x += 5 }()
    return x
}
```

**Output:** `5`

---

### 🔹 **Card 6: Closures and Reference Capture**

**Q:** Will defer capture the variable by value or reference in a closure?

**Example:**

```go
x := 1
defer func() { fmt.Println(x) }()
x = 10
```

**Output:** `10`

---

### 🔹 **Card 7: Panic and Recover in Defer**

**Q:** Can you handle panics in defer?

**Example:**

```go
defer func() {
    fmt.Println("Recovering:", recover())
}()
panic("crash")
```

**Output:** `Recovering: crash`

---

### 🔹 **Card 8: Defer in a Loop**

**Q:** What happens when defer is used in a loop?

**Example:**

```go
for i := 0; i < 3; i++ {
    defer fmt.Print(i, " ")
}
```

**Output:** `2 1 0 `

---

### 🔹 **Card 9: Defer + Loop Variable Trap**

**Q:** Common trap with closures and loop variables?

**Example:**

```go
for i := 0; i < 3; i++ {
    defer func() { fmt.Print(i, " ") }()
}
```

**Output:** `3 3 3`

---

### 🔹 **Card 10: Defer with Return in Main**

**Q:** Will defer run if `return` is called?

**Example:**

```go
func main() {
    defer fmt.Println("deferred")
    return
}
```

**Output:** `deferred`

---

### 🔹 **Card 11: Nested Defer**

**Q:** Can you nest defer statements?

**Example:**

```go
defer func() {
    defer fmt.Println("nested")
}()
```

**Output:** `nested`

---

### 🔹 **Card 12: Panic with Multiple Defers**

**Q:** What happens if a panic occurs, and multiple defers exist?

**Example:**

```go
defer fmt.Println("1")
defer fmt.Println("2")
panic("fail")
```

**Output:**

```
2
1
panic: fail
```

---

### 🔹 **Card 13: Defer on nil Receiver**

**Q:** What happens if a deferred method is called on a nil receiver?

**Example:**

```go
type Test struct{}
func (t *Test) Hello() {
    fmt.Println("Hello")
}

func main() {
    var t *Test = nil
    defer t.Hello()
}
```

**Output:** `Hello`

---

### 🎓 Interview Tips:

* Know when **values are captured** vs **executed**
* Difference between **named vs unnamed** return values
* Only use `recover()` **inside** a deferred function
* Loop + defer + closures are **common traps**
* Be comfortable predicting tricky outputs involving multiple defers
