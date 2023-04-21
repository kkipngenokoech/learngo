# GO FUNDAMENTALS

here we are going to learn about the following go syntax:
 1. identifiers
 2. keywords
 3. data types
 4. variables
 5. constants
 6. operators

## identifiers

identifiers are normally used for identification purposes. this are normally user defined names of program components.

they can be variable names, constants, statement labels, package names or types.

the rules for identifiers are just as for any other programming language:

1. The name of the identifier must begin with a letter or an underscore `(_)`. And the names may contain the letters ‘a-z’ or ’A-Z’ or digits 0-9 as well as the character `‘_’`.
2.  The name of the identifier should not start with a digit.
3. The name of the identifier is case sensitive
4. Keywords is not allowed to use as an identifier name.
5. There is no limit on the length of the name of the identifier, but it is advisable to use an optimum length of 4 – 15 letters only.

## keywords

this are used in predefined actions there are a couple of predefined keywords:


Keyword	Description
break	Ends the execution of a loop or switch statement
case	Specifies a case in a switch statement
chan	Used for communicating between goroutines
const	Declares a constant value
continue	Skips to the next iteration of a loop
default	Specifies the default case in a switch statement
defer	Defers the execution of a function until the surrounding function returns
else	Specifies an alternative branch in an if/else statement
fallthrough	Allows execution to "fall through" to the next case in a switch statement
for	Begins a loop
func	Declares a function
go	Starts a new goroutine
goto	Transfers control to a labeled statement
if	Begins an if statement
import	Imports a package
interface	Declares an interface
map	Declares a map
package	Declares a package
range	Iterates over elements of an array, slice, or map
return	Returns a value from a function
select	Used for non-blocking communication between goroutines
struct	Declares a struct
switch	Begins a switch statement
type	Declares a new type
var	Declares a variable

## data types

