# Day 2: Variables

## What I Learned
- `var` declares 1 or more variables.
- You can declare multiple variables at once.
- Go will infer the type of initialized variables.
- Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an `int` is `0`.
- `:=` is shorthand for declaring and initializing variables, e.g., `var f string = "apple"` can be written as `f := "apple"`. This syntax is only available inside functions.
- `const` defines constants.
- Use `reflect.TypeOf` to check the type of a variable.

## Code
See [variables.go](variables.go).
