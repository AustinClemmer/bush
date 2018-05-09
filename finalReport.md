# <u>Bush- The Belly Up SHell</u>

# A project by Austin Clemmer - May 8, 2018

## Abstract

The aim of this project was to create a shell that I would enjoy to use because of some key features I felt were lacking from the shell I'd previously been using for some time.  Bush (the Belly Up SHell) was going to be a simple take on a shell with a lot of the same features covered with the shell written in Systems II lab.  The shell must be able to read what the users keyed into the command line. Upon the pressing of the return key (Enter key), interpret the input to the point where it would either execute the correctly formatted command entered by the user, or return an error with a status message indicating to the user that something was not right, and give them a nod to the type of error encountered with an exit status number.  The shell should also allow for in-line editing on the individual command line level, including the use of the backspace and delete keys to remove characters, and proper handling of the left and right arrow input keys to move a cursor around the line.  Initially job control was going to mimic that of the previous shell written during the Systems II class, but being something that was barely used by me, it got tossed to the wayside during development.  In the end the project turned out to be a very bare bones implementation of the shell, one without traditional job control (foregrounding/backgrounding), or IO redirection, but one that could be used to do everything necessary for a shell to do.  

Keywords: bush, belly up shell, shell, golang, go

## Table of Contents

1. [Abstract](#abstract)
2. [Introduction and Project overview](#introduction-and-project-overview)
   - [Shells](#shells)
   - [Terminals _then_](#terminals-then)
   - [Terminals _now_](#terminals-now)
   - [Designations](#designations)
   - [The Legacy Problem](#the-legacy-problem)
   - [The Belly Up SHell](#the-belly-up-shell)
3. [Design, Development, and Testing](#design-development-and-testing)
   - [Design](#design)
   - [Development](#development)
   - [Testing](#testing)
4. [Results](#results)
5. [Conclusions and Future Work](#conclusions-and-future-work)
6. [References](#references)
7. [All packages/libraries used](#all-libraries/packages)

## Introduction and Project overview

### Shells

A shell is a user interface which allows the user access to the operating system's services. [[1]](#shell)  The name shell comes from the abstract description that the kernel (the core program inside the operating system that has complete control over the CPU, all memory, and devices attached to the system) [[2]](#kernel) is 'wrapped' up by a shell, and the user interaction goes through this level, before directly affecting the system level components (CPU, memory, devices).  Generally, a user will interact with this program (the shell) via a command line terminal emulator.  

### Terminals _then_

The term terminal comes to be known from the early days of computing, when a physical hardware device would be necessary for the input of data, and the displaying of the data.  This screen and keyboard setup was known as the terminal. [[3]](#terminal) This afforded the user something to input information (keyboard/number pad), and showed the user everything they needed to know about what they were doing with the specific system (on-screen display of information).  One popular example of the early terminals is something known as a teletype. [[4]](#teletype) These devices were similar to modern day computer screens in that they were used to display information inputs and outputs from humans and the computations the computing system completed for the user; however, these terminals predated the computer monitor currently known by decades.  The teletype provided a keyboard as the user interface to work with early mainframe computers. They employed punched tape and punched cards to bridge the gap from human to early computer.  These cards or tapes could be used for data storage or represent the data itself. The machines would then display the output, or return something to the user (from a remote source of data or human input). 

### Terminals _now_

The term terminal has since then been abstracted to mainly refer to terminal emulators, since we have moved away from the pure physical hard copy input/display systems of yore.  The term terminal, presently, is known almost ubiquitously to refer to the terminal emulator within a computing system.  The term terminal emulator can also be referred to as a terminal application or just term for short (for confusions sake I will not be using the shorthand 'term'). [[5]](#terminal-emulator) The terminal emulators that run on computing systems presently are just that, an emulation (emulating being the extension of behavior of one computer system to another separate computer system) of the old terminals from the foundations of modern computing.  [[6]](#emulator)

### Designations

There is a lot of confusion surrounding the correct nomenclature for what the terminal is, and what a shell is.  A user will interact with the terminal emulator in order to use the shell. The shell program will pass along the commands to the operating system's kernel leading to the execution and manipulation of operating system services. This interaction will complete the tasks the user has requested. More simply put, the terminal emulator is a program that provides an emulated terminal for displaying and accepting information in the form of user input; the shell is the program that checks to see what the user has entered, and pass it along to the kernel in order to complete some activity.  The two terms terminal and shell are often conflated because they are hard to separate conceptually.  Plainly, the shell runs in the terminal.

### The legacy problem

The problem with most shells is over complication of what is a simple program. The term "legacy problem" is rampant in shell programs.  The legacy problem can be defined as, "the way it has been done, it will continue to be done".  The most prominent shells in today's modern computing world are the same shells that have been around, and in use, for two to three decades.  This project seeks to condense some of the shell programs that are still in use, by only including useful features, or features the developer thought were necessary in a shell.  

### The Belly Up SHell

The intended user base for this project is anyone who uses a shell, is dissatisfied with bloat, or over-complication of simplistic actions in shell use.  The belly up shell has a no bloat user prompt that only displays a triple bracket leader (">>> ") for each refreshed command line.  It will notify you of any permission conflicts, and malformed commands with a returned error/exit status.  The use of 'sudo' will prompt for the users password, and then lead into the desired activities if permission conflicts are present.  The user can look back through the current session's history with the up arrow key in the order of most recently used spanning back to the first command entered during that session.  In addition to this history, a hidden history file will also be created and populated in the user's home directory.  The built-in command 'jobs' will present the user with a list of the current processes through the use of the ps command.  Given the way the jobs feature was implemented it will accept all flags that the ps command accepts.  The command 'cd dir/' will help the user change directories, the default 'cd' with no path will take the user to their home directory as specified at the operating system level.  To list the current contents of a directory, the command 'ls' will provide a color-coded output of the files and directories in the directory it is invoked.  To exit the shell, the command 'exit' will kill the shell, and print a nice departure message to the user.  

## Design, Development, and Testing

### Design

This project is written in the programming language Go. [[7]](#go) Go was chosen for this project based on it's highly concurrent nature, and ability to interact at a system level with the kernel.  The choice to use Go was one that was taken rather lightly on a recommendation from a colleague.  This language turned out to be a great choice overall, and little in the way of similar projects could be found on-line.  

When the user enters the program, they are met with a simple greeting message to welcome them to the Bush experience.  They are then presented with a simple prompt (">>> "), and can begin using the shell to navigate around in the computer, and perform whatever tasks are required during the session.  The user completes these actions by entering in commands to the command line and pressing the return key, this is where the magic happens.

The library that extends the capability of the go language to digest user input, and display a new prompt, is known as "Readline".  Once the user strikes the return key, the string they entered into the command line gets initialized to an internal variable.  This variable is passed along to an executing function to be handled, and to ensure the correct message to reach the operating system.  This string is tokenized into the main command, and its flags/arguments.  Tokenization, in this context, refers to the process of classifying sections of the string. [[8]](#token) An empty string will return the user a new prompt, and do nothing.  The main command and it's arguments are passed to a function in the library in the Go language package "OS" known as "exec".  The package exec runs external commands, remapping stdin and stdout, intentionally not invoking the system shell.  The standard streams, stdin, stdout, and stderr, are preconnected input and output communication channels.  [[9]](#stds) Upon execution of a command with an interactive shell, these streams are connected to the terminal.  Stdin is the representation of the data stream into the program, stdout being the stream of data out of the program, and stderr being the stream which error messages come through.  Specifically, the Command function in the "os/exec" package digests the command and arguments, and returns a command structure.  The shell internally maps the operating system level stdin, stdout, and stderr, to the command stdin, stdout, stderr.  The command is then ran, and error checking occurs to make sure everything was properly formatted, and that no errors occurred.  Each time a command is entered, the "Readline" library takes the raw string from the user, and appends it to a master history file.  If no such file exists, one will be made in the correct place, and the commands will begin to populate it.   The error handling is done with the error interface provided by the "runtime" Go package.  This package allows the developer to distinguish types that are runtime errors from ordinary errors.  After the successful completion of a command, the "Readline" library returns a new prompt to the user, and awaits another command.

### Development

Initially the project used a simple print statement to display a prompt, and digestion of user input was achieved by reading in one rune (a rune being, simplistically, a Go language term for a unicode representation of a key pressed on the keyboard) [[10]](#rune) at a time. The digestion of these runes served to be rather tricky. Upon doing further research it was discovered that not only did it listen for all inputs as runes (even arrow keys), Go made no guarantees that these values were normalized.  This meant that feeding the command structure the exact string the user entered would lead to a very strange input.  The backspace and delete keys were printing their unicode representations in the middle of an otherwise well formed command.  Each time the user struck an arrow key the unicode representations of these keys were also printed to the terminal.  There were no arguments allowed under this model, and it could barely be considered a functioning shell.  It utilized the library "pkg/term" to invoke a virtual terminal to handle the commands, and this was also not exactly what Bush needed to do.  

From this stage, the proof of concept 'skateboard' model had been achieved although it was hardly considered user-friendly.  Because the user could move the cursor to anywhere in the terminal and start printing characters/overwriting characters, the program was written to ignore the arrow keys all together. Backspace didn't actually remove characters, but rather overwrote them with blank spaces in a nifty, but naive, roundabout solution to the problem.  The function to change directories was added in, but would simply do nothing if the directory didn't exist, or was mistyped.  The prompt had some color added to it to appear more like a real shell, but it was very far from being that.  The next update added a working path to be displayed each time a prompt was printed.  The history feature was added in a subsequent update, but the project was still dancing around some very glaring issues with how the user interacted with the program. 

The next iteration used a different method to receive the user input, and allowed for the user to move the arrow keys, and brought with it working backspace/delete.  The problem with this package was that it was written for a very specific purpose (to create command line applications with user suggestions being a main feature), and was at too high a level of abstraction to be properly utilized in the context of this project.  It was functioning exactly how the shell should function, but the prompt would display two suggestions to the user about what 'cd' and 'ls' were and what they did.  This was the early implementation of a feature thought necessary, but which turned into a nuisance.  Each time the user returned a prompt, the suggestion list would show, even when the user didn't prompt it to show by typing one of the beginning letters of the commands.

At this point, the shell was working, but the user experience was more annoying than anything, and the project stagnated as a result of developer frustrations.  Then the "Readline" library was discovered.  The Readline library did everything like a C-style readline function, and was at a low enough level of abstraction to be able to extend the functionality of the shell while not hijacking the core principle of keeping it simple.  backspace, and delete worked exactly how they are expected to work, and the arrow keys moved the cursor like expected.  The built-in commands were then expanded to include a 'jobs' function which lists the current active processes (through the use of the 'ps' command),[[11]](#ps) and accepts the normal 'ps' command flags.  The directory changing functionality was extended to display whether or not the directory was reached or didn't exist.  The program was modularized as much as possible, and extensive error checking was added into the program to make sure the user didn't do anything without receiving some sort of message as to what went wrong with the command.

### Testing

Most initial testing came in the form of extensive use, and would be considered manual testing.  By the time the shell was working in the most rudimentary fashion, it was used as the shell for the developer to develop the shell.  The use of the Go 'testing' package was added into the program to make this process more accessible, and easier to automate.  When it came time to automate testing, the first thing that was tested was the display of the printed output to the terminal on stdout.  If a command is run, and the executor function returns an error, the test will display an appropriate message to the developer about what went wrong.  The next test that was implemented was for the built-in feature of 'cd'.  This feature is expected, when supplied with no path or directory name, to return the user to the home directory.  The command is executed, and the test looks at the user's home directory.  This directory is saved into a variable. Next, the current working directory is queried by the program, and saved into another variable.  If the values of the paths stored in these variables are not the same, the command did not return the user to the correct place, and it returns one of two errors to let the developer know what went wrong.  Another major issue with the 'cd' command was that it just returned a prompt to the user without notifying them that the directory change was a failure.  After the error checking was implemented successfully into the program, a test to make sure an error was happening was implemented.  This test invokes 'cd bogusLongNameDirectory' (in the hopes that the user hasn't added a directory named bogusLongNameDirectory), and returns a logged error message that an error was received as intended by the program.  The next test implemented checks to make sure the 'exit' command actually exits the shell cleanly without an error.  If an error occurs it is displayed to the developer.

## Results

The initial goals of the project were lofty, and unrealistic.  Development of the project took a long time due to poor time management, and was over complicated by the use of the wrong tool set in the beginning.  Had the project used the correct Readline library functionality in the beginning, seemingly hard problems would have just vanished, or been reduced to purely trivial ones. An initial hurdle that was overcome stubbornly was the problem of working with a language that was new to me.   The time spent to learn the tools in the language was very small, and more of a stumble than that of a smooth walk through the new material.   Where I should have taken the time to learn piece by piece, I was trying to follow tutorials, and copy-paste things I thought would work.  The initial feature list was out of control, and I just didn't realize what actually needed to be done.  

Essential features included:

- The Shell must be able to read user input, interpret the input, and execute the user input.
  - This version of the shell prints a prompt to the user, accepts their string as input, and formats this input string to be compliant with the operating system service or built-in command in the shell by tokenizing the command into a command and it's arguments.  If the user enters in a malformed command or something that is not accessible for the user.  The shell will tell the user that the command is not found, or that the permission is denied.
- Upon a user hitting the return key, the input must execute, and return a fresh prompt.  The backspace key must consume characters to the left of the cursor, and the delete key must consume characters to the right of the cursor.
  - In the current version of Bush, pressing the return key will submit the user string to the executor, where it will be formatted properly for execution.  Upon the execution, successful or otherwise, a new prompt will be returned to the user.  The left and right arrow keys will move the cursor on the current line.  When a cursor is placed to the right of a character and the backspace key is pressed, the character on the line to the left of the cursor will be erased.  When the cursor has a character to the right of it, and the delete key is pressed, the character to the right of the cursor will be consumed.   
- Job control that employs a set of built-in command support: bg (places job in background), fg (bring job to the foreground, and make it current job), jobs (lists the active jobs, and their active IDs), suspend (suspends the execution of this shell until it receives a signal to continue), and kill/wait (kill exits the process, and wait 'waits' until the child process exits) built-ins will be implemented.
  - Only some of these features made it into the application.  I began to do some real soul searching about what tools I used with a shell, and the backgrounding/foregrounding functions were nowhere on my list.  Seeing as this project was tailored to custom fit my needs, I scrapped these in the final implementation of the project; although, there is logic in place to handle running a command in the background.  The 'jobs' command lists the currently running processes, and 'kill' works to kill a process when provided a process id, or other standard flags as listed by the man page.
- I/O redirection in the form of the pipe '|' will allow for users to send things between commands, stdin and stdout can be redirected using the '<' and '>' characters, and stderr will be redirected with '/>'.
  - This feature did not make it into the shell, and I am very upset that it didn't.  The use of redirection is very much an useful thing to have in a shell.  It is not necessary for the shell to have redirection, but it is very nice for the user to be able to pass things from one place to another or from a file to program, etc.
- The shell will be able to handle unmatched system calls, and inform the user when errors arise from permission conflicts.
  - This feature is present in the current build of Bush, and the use of 'sudo' to overcome permission conflicts where necessary, is working as expected.  When a string is entered with sudo preceding the actual command, stdout displays a prompt to enter in a user password, and complete whatever action the user needed to have permissions for.  

A 'want to have' goal that wasn't necessarily considered top priority, but was accomplished anyway is the shell being portable to all UNIX systems.  The libraries used, and the way the commands are executed should work ubiquitously across all UNIX- based computing systems.  Readline is even touted to be windows compliant, although nothing else would work. 

All test coverage that was written to test the shell's built in commands is passing, and the current build is perfectly functional.  The features listed in the essential feature list that were actually deemed essential made it into this version of the shell.  

## Conclusions and Future Work

The problem with modern shells is that they often come with a lot of old features, and things that aren't exactly necessary for my usage of them.  The aim of this project was to simplify the shell to the user so that they only interact with the things necessary to get their work done.  The approach was to build a small shell step by step, and only add in the things that were absolutely crucial in my opinion to shell use.  The result was a very bare bones shell that is built for someone who knows exactly what they are doing and exactly what they need to accomplish when using a shell. This shell doesn't hand hold, and doesn't provide much in the way to help the user figure out what is going on or what they 'need' to do in order to successfully navigate the computing system.

I learned a great deal about the language Go during this process, and how shells interact with the kernel.  This was my first time using this language, and I feel like i did myself a disservice in not using something I already knew to get better with it.  I barely scratched the surface with this language.  Using something like C would have helped me sharpen skills I plan to apply in my future. and feel less like a waste of my time now that everything is said and done.  The results of the shell itself are perfectly fine with me.  When I started this project, I would have really loved to have a clock in the prompt. I do not feel like this is at all necessary to my definition of what a shell should be at this time. 

The results of the project are a display in pure utilitarianism.  The shell is as bare and minimalistic as it can pretty much be without being broken.  The user interface is minimalistic and provides the user with a clear division of where the new prompt is, and where the old ones were executed.  

I would love to continue to tweak this shell to be exactly what I am looking for in my work.  IO redirection would be the first thing I implemented in the shell after this. Further development would most likely lead to the completion of the bg/fg commands for foregrounding and backgrounding processes, not because I really want this feature, but if others were to want to use this shell they may want the feature.  I would also really love to make the prompt somewhat customizable, maybe to display the current working directory.   Nothing should affect the shell's usability at this point, it is completely functional, and only open to extension.

## References

- [<a id="shell">1</a>]  Retrieved May 8, 2018 from https://en.wikipedia.org/wiki/Shell_(computing)
- [<a id="kernel">2</a>]  Retrieved May 8, 2018 from https://en.wikipedia.org/wiki/Kernel_(operating_system)
- [<a id="terminal">3</a>]  Retrieved May 8, 2018 from https://en.wikipedia.org/wiki/Computer_terminal
- [<a id="teletype">4</a>]  Retrieved May 8, 2018 from http://www.linfo.org/teletype.html
- [<a id="terminal-emulator">5</a>]  Retrieved May 8, 2018 from https://en.wikipedia.org/wiki/Terminal_emulator
- [<a id="emulator">6</a>]  Retrieved May 8, 2018 from https://en.wikipedia.org/wiki/Emulator
- [<a id="go">7</a>]  Retrieved May 8, 2018 from https://golang.org
- [<a id="token">8</a>]  Retrieved May 8, 2018 from https://www.techopedia.com/definition/13698/tokenization
- [<a id="stds">9</a>]  Retrieved May 8, 2018 from https://en.wikipedia.org/wiki/Standard_streams
- [<a id="rune">10</a>] Retrieved May 8, 2018 from https://blog.golang.org/strings
- [<a id="ps">11</a>] Retrieved May 8, 2018 from http://man7.org/linux/man-pages/man1/ps.1.html


## All libraries/packages

- [Go format (gofmt)](https://golang.org/cmd/gofmt/)
- [Dep](https://github.com/golang/dep)
- [testing](https://golang.org/pkg/testing/)
- [chyzer/readline](https://github.com/chzyer/readline)
- [os](https://golang.org/pkg/os/)
- [os/exec](https://golang.org/pkg/os/exec/)
- [path/filepath](https://golang.org/pkg/path/filepath/)
- [runtime](https://golang.org/pkg/runtime/)
- [strings](https://golang.org/pkg/strings/)

