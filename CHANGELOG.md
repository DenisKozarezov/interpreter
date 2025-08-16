# Changelog

<!-- START --->

## [v0.0.1] - 16/08/2025 - Initial Release

### ✨ Language Features
- **Core Syntax**:
   - Arithmetic operators: `+ - * /`
   - Comparison operators: `== != < > <= >=`
   - Logical operators: `&& || !`
   - Operators precedence: `a + b * c`, `(a + b) * c`, `a + (b * c)`, etc...
- **Type System**:
   - Primitive types: integers, booleans, strings
   - Functions (`fn`) and closures
- **Control Flow**:
   - Conditionals: `if/else` statements
   - Code blocks with `{ }`
   - Return `return`
- **Statements**
   - `let`
- **Built-In Functions**
   - Function `len` for arrays and strings

### 🛠️ Core Features
- Implemented basic lexer and parser
- Implemented basic AST evaluation
- Implemented errors handling
- Source line and symbol position tracking for errors output

### 📝 Example Code
```ipret
// Factorial function
let fact = fn(n) {
    if (n <= 1) { return 1; }
    return n * fact(n-1)
};

fact(5);
```