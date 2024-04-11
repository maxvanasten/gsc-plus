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

### Current working file:

```
#include common_scripts\utility;
#include maps\mp\gametypes_zm\_hud_util;
#include maps\mp\zombies\_zm_utility;

fn get_hello_string(name) {
	let first_part = "Hello, ";
	let second_part = name + "!";

	return first_part + second_part;
}

fn hello_on_obj() {
	print("Hello, " + self.name + "!");
}

fn greet_person(person) {
	person > hello_on_obj();
	person >> hello_on_obj();
}
```

Compiles to:

```
#include common_scripts\utility;
#include maps\mp\gametypes_zm\_hud_util;
#include maps\mp\zombies\_zm_utility;

get_hello_string(name)
{
	first_part = "Hello, ";
	second_part = name + "!";
	return first_part + second_part;
}

hello_on_obj()
{
	print("Hello, "+self.name+"!");
}

greet_person(person)
{
	person hello_on_obj();
	person thread hello_on_obj();
}
```
