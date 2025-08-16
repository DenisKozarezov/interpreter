# Changelog

<!-- START --->

## [v0.0.1] - 16/08/2025 - Initial Release

### ✨ Language Features
- **Core Syntax**:
   - Arithmetic operators: `+ - * /`
   - Comparison operators: `== != < > <= >=`
   - Logical operators: `&& || !`
   - Operators precedences: `a + b * c`, `(a + b) * c`, `a + (b * c)`, etc...
- **Type System**:
   - Primitive types: `int`, `bool`, `string`
   - Functions: `fn`, closures
- **Control Flow**:
   - Conditionals: `if/else` statements
   - Code blocks with `{ }`
   - Return `return`
- **Statements**
   - `let`
- **Built-In Functions**
   - Function `len` for arrays, strings.

### 🛠️ Core Features
- Added REPL mode for interactive execution
- Implemented basic lexer and parser
- Source line and symbol position tracking for error handling

### 📝 Example Code
```ipret
// Factorial function
let fact = fn(n) {
    if (n <= 1) { return 1; }
    return n * fact(n-1)
};

fact(5);
```