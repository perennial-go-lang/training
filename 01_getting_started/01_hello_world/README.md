# Hello, World

This is a very basic program that will help us understand the way of structuring a typical Go program. In Go, each program is a part of a package. Packages are used as a way to structure code in Go. Every .go file belongs to one and only one package. A package may contain more than one .go files too. It is not necessary for the file name to match the name of the package it belongs to.

The package to which a .go file belongs to is written on the very first line of the program. This is followed by import declaration. We can specify any packages our program is dependent on using the import declaration. We can import any package in our program except the main package. The main package indicates the entrypoint of the program. The execution of any go program begins from the main function defined in the main package of the program.

