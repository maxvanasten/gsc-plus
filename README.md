# GSC Plus

## GSC, Slightly improved

### About

This is primarily a learning project for me. Its a superset of GSC that slightly improves (subjectively, ofcourse) the syntax of GSC for Call of Duty scripting.

### Syntax changes

#### Variable Declarations

```
let varname = "Hello, World!";
```

#### Declaring a function

```
fn some_function(arg1, arg2) {
    return arg1 + arg2;
}
```

#### Running functions on objects

This code runs a function on an object:
```
level > some_function();
```

This code runs a function threaded on an object:
```
level >> some_function();
```
