# GForth Compiler
minor GForth compiler written in Go. for arch `x64`

To check GForth functionalities check
https://learnxinyminutes.com/docs/forth/

### How to build the compiler
* Install `Go` version <= 1.21.1
* run `make`

It will build the forth compiler that is written with `Go`.
It gives you an executable with the name `mf`.

### How to test the language
* `mf ./test1.mf` it will compile this file and give you and executable with the same name `test1`
* Then run it `./test1`

Happy Forthing

