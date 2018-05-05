# bush
## bush- Belly Up SHell

- A capstone project by Austin Clemmer

Bush is a very rudimentary implementation of a shell, written in the language Go.
It has been written with all UNIX systems in mind. 

### User Installation Instructions
1. Clone or download repo into your local go/src directory
2. Use "go build" in the bush/ directory to build and install the application
3. Use "go run main.go" or "./bush" in the project directory to launch the application

If the user so chooses, the "go install" command will place the executable in the user's 
path, and allow for invokation outside of the source directory.  This application can be 
launched, after install, by simply typing "bush".

### Developer Instructions
If a developer so chooses to work on this project, they can do so by downloading the source,
and editing the 'main.go' file.

Testing can be accomplished by running 'go test' using the go 'testing' package.  The tests
can be found in the 'main_test.go', and can be extended to include more as development continues.
Tests can be ran with 'go test -v' to display a more verbose description, and display log messages,
of the tests currently implemented. 

Quality control of this repository has been taken on with the use of the go 'go/format' linter, which
implements the standard formatting of Go source code.

In terms of package management, the tool Dep has been included in this project, to make managing dependencies trivial.

### Resources
[Project Proposal](projectProposal.md)
[Final Technical Report (draft)](finalReport.md)
