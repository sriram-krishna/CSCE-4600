# Project 2: Shell Builtins

## Description

For this project I have added  more commands to a simple shell. 

The shell was already written, And I chose five (5) shell builtins (or shell-adjacent) commands to rewrite into Go, and integrate into the Go shell.

As an example, two shell builtins have already been added to the package builtins:

- `cd`
- `env`

Additionally, the five requested builtins that have been added to the package are:
- `echo`
- `clear`
- `pwd`
- `date`
- `whoami`

## Steps

1. Clone down the example input/output and skeleton `main.go`:

    `git clone https://github.com/sriram-krishna/CSCE-4600`
 
2. Copy the `Project2` files to your own git project.

    1. In your `go.mod`, replace "sriram-krishna" in the module line with your GitHub name, e.g.:

      - "module github.com/sriram-krishna/CSCE-4600" changes to "module github.com/CoolStudent123/CSCE4600"
  
    2. In the `main.go`, replace "sriram-krishna" in the imports with your package path, e.g.:

      - "github.com/sriram-krishna/CSCE-4600/Project2/builtins" changes to "github.com/CoolStudent123/CSCE4600/Project2/builtins"

3. Build the project by typing `go build`.
