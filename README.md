# Monkey Programming Language

A complete interpreter and compiler implementation for the Monkey programming language, built working through the books "Writing an Interpreter in Go" and "Writing a Compiler in Go" by Thorsten Ball.

## About Monkey

Monkey is a dynamically-typed programming language with a C-like syntax that supports:

- Variable bindings with `let` statements
- Integers, booleans, strings, arrays, and hash maps
- Arithmetic and logical expressions
- Built-in functions for common operations
- First-class and higher-order functions
- Closures with lexical scoping
- Control flow with if/else conditionals
- Implicit return values (last expression in block)

### Quick Example

```monkey
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      return 1;
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};

fibonacci(15);
```

## Language Reference

### Data Types

#### Integers

```monkey
let age = 25;
let negative = -42;
let result = 10 + 5 * 2;  // 20
```

#### Booleans

```monkey
let isTrue = true;
let isFalse = false;
let comparison = 5 > 3;   // true
```

#### Strings

```monkey
let greeting = "Hello, World!";
let name = "Monkey";
let message = greeting + " I'm " + name;  // "Hello, World! I'm Monkey"
```

#### Arrays

```monkey
let numbers = [1, 2, 3, 4, 5];
let mixed = [1, "hello", true, fn(x) { x * 2 }];
let nested = [[1, 2], [3, 4], [5, 6]];

// Array access
let first = numbers[0];     // 1
let last = numbers[4];      // 5
```

#### Hash Maps

```monkey
let person = {
  "name": "John",
  "age": 30,
  "city": "New York"
};

// Hash access
let name = person["name"];           // "John"
let age = person["age"];             // 30

// Mixed key types
let mixed_hash = {
  "string_key": "value1",
  42: "numeric key",
  true: "boolean key"
};
```

### Functions

#### Function Definition and Calls

```monkey
// Simple function
let add = fn(a, b) {
  a + b
};

let result = add(5, 3);  // 8

// Function with explicit return
let multiply = fn(x, y) {
  return x * y;
};

// Anonymous function
let square = fn(x) { x * x };
```

#### Higher-Order Functions

```monkey
// Function that returns a function
let makeMultiplier = fn(factor) {
  fn(x) { x * factor }
};

let double = makeMultiplier(2);
let triple = makeMultiplier(3);

double(5);  // 10
triple(4);  // 12

// Function that takes a function
let applyTwice = fn(func, value) {
  func(func(value))
};

applyTwice(double, 3);  // 12 (double(double(3)))
```

#### Closures

```monkey
let newCounter = fn() {
  let count = 0;
  fn() {
    count = count + 1;
    count
  }
};

let counter = newCounter();
counter();  // 1
counter();  // 2
counter();  // 3

// Each counter maintains its own state
let counter2 = newCounter();
counter2(); // 1
counter();  // 4
```

### Control Flow

#### Conditionals

```monkey
let max = fn(a, b) {
  if (a > b) {
    a
  } else {
    b
  }
};

// Ternary-like usage
let status = if (age >= 18) { "adult" } else { "minor" };

// Nested conditionals
let grade = fn(score) {
  if (score >= 90) {
    "A"
  } else {
    if (score >= 80) {
      "B"
    } else {
      if (score >= 70) {
        "C"
      } else {
        "F"
      }
    }
  }
};
```

### Built-in Functions

#### Array Functions

```monkey
let arr = [1, 2, 3, 4, 5];

len(arr);           // 5 - get array length
first(arr);         // 1 - get first element
last(arr);          // 5 - get last element
rest(arr);          // [2, 3, 4, 5] - all but first
push(arr, 6);       // [1, 2, 3, 4, 5, 6] - add element
```

#### String Functions

```monkey
let text = "Hello";
len(text);          // 5 - get string length
```

#### Output Function

```monkey
puts("Hello, World!");           // prints: Hello, World!
puts("Value:", 42, true);        // prints: Value: 42 true
```

### Operators

#### Arithmetic Operators

```monkey
let a = 10;
let b = 3;

a + b;    // 13 - addition
a - b;    // 7  - subtraction
a * b;    // 30 - multiplication
a / b;    // 3  - division (integer division)
```

#### Comparison Operators

```monkey
5 == 5;     // true  - equality
5 != 3;     // true  - inequality
5 > 3;      // true  - greater than
3 < 5;      // true  - less than
```

#### Logical Operators

```monkey
!true;      // false - negation
!false;     // true
!(5 > 3);   // false
```

## Comprehensive Examples

### Recursive Functions

```monkey
// Factorial
let factorial = fn(n) {
  if (n <= 1) {
    1
  } else {
    n * factorial(n - 1)
  }
};

factorial(5);  // 120

// Tree traversal
let sumTree = fn(tree) {
  if (len(tree) == 0) {
    0
  } else {
    tree[0] + sumTree(rest(tree))
  }
};

sumTree([1, 2, 3, 4, 5]);  // 15
```

### Array Processing

```monkey
// Map function implementation
let map = fn(arr, func) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0) {
      accumulated
    } else {
      iter(rest(arr), push(accumulated, func(first(arr))))
    }
  };
  iter(arr, [])
};

let double = fn(x) { x * 2 };
map([1, 2, 3, 4], double);  // [2, 4, 6, 8]

// Filter function implementation
let filter = fn(arr, predicate) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0) {
      accumulated
    } else {
      let head = first(arr);
      let tail = rest(arr);
      if (predicate(head)) {
        iter(tail, push(accumulated, head))
      } else {
        iter(tail, accumulated)
      }
    }
  };
  iter(arr, [])
};

let isEven = fn(x) { x / 2 * 2 == x };
filter([1, 2, 3, 4, 5, 6], isEven);  // [2, 4, 6]
```

### Hash Map Operations

```monkey
// Working with nested data
let users = [
  {"name": "Alice", "age": 25, "city": "NYC"},
  {"name": "Bob", "age": 30, "city": "LA"},
  {"name": "Charlie", "age": 35, "city": "Chicago"}
];

let getUser = fn(users, name) {
  let iter = fn(arr) {
    if (len(arr) == 0) {
      null
    } else {
      let user = first(arr);
      if (user["name"] == name) {
        user
      } else {
        iter(rest(arr))
      }
    }
  };
  iter(users)
};

let alice = getUser(users, "Alice");
puts("Found user:", alice["name"], "age", alice["age"]);
```

### Complex Closures

```monkey
// Function factory with multiple closures
let createCalculator = fn(initial) {
  let value = initial;

  {
    "add": fn(x) { value = value + x; value },
    "multiply": fn(x) { value = value * x; value },
    "get": fn() { value },
    "reset": fn() { value = initial; value }
  }
};

let calc = createCalculator(10);
calc["add"](5);        // 15
calc["multiply"](2);   // 30
calc["get"]();         // 30
calc["reset"]();       // 10
```

## Architecture

This implementation includes both an **interpreter** and a **compiler** for the Monkey language:

### Interpreter Info

- **Lexer**: Tokenizes source code into meaningful tokens
- **Parser**: Builds an Abstract Syntax Tree (AST) using Pratt parsing
- **Evaluator**: Tree-walking interpreter that directly executes the AST

### Compiler Info

- **Lexer & Parser**: Same as interpreter (reused)
- **Compiler**: Generates bytecode instructions from the AST
- **Virtual Machine**: Stack-based VM that executes the bytecode

## Project Structure

```
monkey-lang/
├── ast/           # Abstract Syntax Tree definitions
├── code/          # Bytecode instructions and operations
├── compiler/      # Bytecode compiler and symbol table
├── evaluator/     # Tree-walking interpreter
├── lexer/         # Tokenization
├── object/        # Runtime object system and built-ins
├── parser/        # Recursive descent parser
├── repl/          # Read-Eval-Print Loop
├── token/         # Token definitions
└── vm/            # Virtual machine and execution frames
```

## Features

### Language Features

- **Data Types**: Integers, booleans, strings, arrays, hash maps
- **Functions**: First-class functions with closures
- **Control Flow**: If/else conditionals
- **Built-ins**: `len`, `first`, `last`, `rest`, `push`, `puts`
- **Operators**: Arithmetic (`+`, `-`, `*`, `/`), comparison (`==`, `!=`, `>`, `<`), logical (`!`)

### Implementation Features

- **Dual Execution**: Both interpreted and compiled execution modes
- **Memory Management**: Garbage collection through Go's runtime
- **Error Handling**: Comprehensive error reporting for syntax and runtime errors
- **REPL**: Interactive shell for testing and exploration
- **Persistent State**: REPL maintains variables and functions across sessions

## Building and Running

### Prerequisites

- Go 1.19 or later

### Build

```bash
go build -o monkey main.go
```

### Run REPL

```bash
./monkey
```

Or directly with Go:

```bash
go run main.go
```

## Testing

The project includes comprehensive test suites for all components:

```bash
# Run all tests
go test ./...

# Run tests for specific components
go test ./lexer
go test ./parser
go test ./compiler
go test ./vm
go test ./evaluator
```

### Test Coverage

- **Lexer**: Token recognition and edge cases
- **Parser**: AST construction and operator precedence
- **Compiler**: Bytecode generation for all language constructs
- **VM**: Bytecode execution and stack operations
- **Evaluator**: Expression evaluation and environment handling

## REPL Usage Examples

### Basic Operations

```
>> let name = "Monkey";
Monkey
>> let age = 1;
1
>> let message = "Hello, " + name + "! You are " + puts(age) + " year old.";
1
Hello, Monkey! You are null year old.
>> message
Hello, Monkey! You are null year old.
```

### Working with Arrays

```
>> let fruits = ["apple", "banana", "cherry"];
[apple, banana, cherry]
>> len(fruits)
3
>> first(fruits)
apple
>> rest(fruits)
[banana, cherry]
>> push(fruits, "date")
[apple, banana, cherry, date]
```

### Function Definition and Usage

```
>> let greet = fn(name) { "Hello, " + name + "!" };
>> greet("World")
Hello, World!
>> let add = fn(a, b) { a + b };
>> add(5, 3)
8
```

### Higher-Order Functions

```
>> let twice = fn(f, x) { f(f(x)) };
>> let addOne = fn(x) { x + 1 };
>> twice(addOne, 5)
7
```

## Implementation Details

### Compiler Architecture

The compiler generates bytecode for a stack-based virtual machine:

- **Opcodes**: 20+ instruction types covering all language operations
- **Constants Pool**: Efficient storage of literals and compiled functions
- **Symbol Table**: Manages variable scoping (global, local, builtin, free)
- **Closures**: Captures free variables for lexical scoping

### VM Architecture

- **Stack Machine**: 2048-slot evaluation stack
- **Frames**: Call stack for function execution contexts
- **Globals**: 65536-slot global variable storage
- **Built-ins**: Native function implementations

### Performance

The compiler/VM approach significantly outperforms the tree-walking interpreter:

- **Fibonacci(35): ~100x faster execution**
- **Reduced memory allocation** through bytecode reuse
- **Optimized instruction dispatch**

### Language Limitations

Current limitations of the Monkey language:

- No loops (while, for) - use recursion instead
- Integer division only (no floating-point numbers)
- No string interpolation
- No module/import system
- No exception handling
- Limited built-in functions

## Acknowledgments

This implementation is based on the books:

- **"Writing an Interpreter in Go"** by Thorsten Ball
- **"Writing a Compiler in Go"** by Thorsten Ball

I found these books super helpful they walk you through building a programming language from scratch in a really clear way. If you're interested in how languages work under the hood like I was, I'd definitely recommend checking them out!

## License

This project is for educational purposes. Please refer to the original books for commercial usage guidelines.
