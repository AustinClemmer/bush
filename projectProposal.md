# Metadata 

#### bush- the Belly Up SHell    --    A project by Austin Clemmer



# Project Overview

The bush project seeks to satisfy my own need for a shell highly tailored to fit my personal use preferences.  

	#### Essential Features

 - The Shell must be able to read user input, interpret the input, and execute the user input.
- Upon a user hitting the return key, the input must execute, and return a fresh prompt.  The backspace key must consume characters to the left of the cursor, and the delete key must consume characters to the right of the cursor.
- Job control that employs a set of built-in command support: bg (places job in background), fg (bring job to the foreground, and make it current job), jobs (lists the active jobs, and their active IDs), suspend (suspends the execution of this shell until it receives a signal to continue), and kill/wait (kill exits the process, and wait 'waits' until the child process exits) built-ins will be implemented.
- I/O redirection in the form of the pipe '|' will allow for users to send things between commands, stdin and stdout can be redirected using the '<' and '>' characters, and stderr will be redirected with '/>'.
- The shell will be able to handle unmatched system calls, and inform the user when errors arise from permission conflicts.

####Want to have

- Shell that is portable to all UNIX systems, achievable by careful design of system calls.
- File globbing, and support for wildcard variables in the calls.
- Prompt customization in the form of a fully customizable prompt (able to execute system level programs and display their information in the prompt, have full color support, and real-time refresh/draw integration).
- Mouse mode that copies to a system level clipboard for interaction outside of the terminal emulator with the text.
- Shell that is easily scriptable with a straightforward document describing how to accomplish this.

#### Bells and whistles

- Windows support for the shell.
- Tab completion/ suggestions for mismatched command.

Users will want to use this shell because it condenses all of the high level functionality of other shells into a straightforward and easy to use shell.  The selling point of the shell will be usability and human feedback as a guide to effective shell use.



# Similar Existing Work

Many shells exist out there in the wild.  There is a project online that aims to provide a small scale proof-of-concept shell that is written in the Go language.  [https://github.com/driusan/gosh] My project would be different in the minutia of how to handle jobs, I/O redirection, and would provide customizability.



# Previous Experience

I have previously worked on getting a rudimentary shell together in Systems II last semester.  This project is in a different language, but one that is similar to C in that it can interact with systems level commands.



# Technology

- With the 'testing' library I can automate testing for the shell.
- The 'dep' package manager will help manage dependencies, and the software as a deployable portable unit.
- The 'GoFmt' tool to formatting will be applied throughout the software development stage.  It uses tabs for indentation, and blanks for alignment. 
- Scripting will allow for automation of the automation tools that clean and format the code.



# Risk Areas

The biggest risk area for this project is going to be the fact that I am unexperienced writing in this language.  The project will provide an outlet for me to learn that language, but I have little experience with it going into the project.  I am still fuzzy on how the clipboard interacts with the application, and how integrating a mouse for copy/pasting features will be handled.  I have no experience with the testing package for Go, and no experience with the dep tool for package management (as I have never had to think about deploying a standalone application).  Simply writing some small simple programs and trying these tools out on it would be sufficient to getting my feet wet/ start learning the tools.

