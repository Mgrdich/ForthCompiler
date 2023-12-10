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

Happy Forthing

