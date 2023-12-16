# GForth Compiler
minor GForth compiler written in Go. 

Supported Architectures
* `x64`

GForth functionalities can be checked https://learnxinyminutes.com/docs/forth/

### How to build the compiler
* Install `Go` version <= 1.21.1
* run `make`

It will build the forth compiler that is written with `Go`.
It gives you an executable with the name `mf`.

### How to test the language
* `mf ./samples/test1.mf` it will compile this file and give you and executable with the same name `test1`
* Then run it `./test1`

### What's to Come
* More functionality in the language
* More Parameters in the CLI
* Writing test suites

Happy Forthing


### Supported Operations

Take each line as a clean slate

* Stack push Eg: `1 2 3` 
* Stack print `1 2 3` -> `<3> 1 2 3 ok`
* Stack Pop and print `1 2 3 .` -> `3`
* Arithmetic Operation Addition `1 2 + .` -> `3`
* Arithmetic Operation Subtraction `1360 23 - .` -> `1337 ok`
* Arithmetic Operation Multiplication `6 7 * .` -> `42 ok`
* Keyword Operation Negation `99 negate .` -> `-99 ok`
* Keyword Operation Absolute `-99 abs .` -> `99 ok`
* Keyword Operation Maximum `52 26 max .` -> `52 ok`
* Keyword Operation Minimum `52 26 min .` -> `26 ok`
* Keyword Operation Duplication `3 dup .s` -> `<2> 3 3 ok`
* Keyword Operation Swap `3 5 swap .s` -> `<2> 5 3 ok`
* Keyword Operation Rotate `6 5 4 rot .s` -> `<3> 4 5 6`
* Keyword Operation Nip `6 5 4 nip .s` -> `<2> 6 4`
* Keyword Operation Tuck `1 2 3 4 tuck .s` -> `<5> 1 2 4 3 4`
* Keyword Operation Over `1 2 3 4 over .s` -> `<5> 1 2 3 4 3`
