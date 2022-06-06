# Simple logging system

This simple logging system provides logging levels, file storage and file rotation.
It also provides a command to display colorized logs. 

By default, the logging output is printed to stdout.

## Installation

Since clog provides a command to display logs with colors, the clog package must be
installed with the command line instruction 

```bash
$ go install github.com/chmike/clog/cmd/...@latest
```

This installs the package and the commands `clogClr` and `testClog`. 

## Usage

The command `testClog` generates various logging messages output to stdout. The first 
character of a line specifies the log level.

 - F : Fatal error message causing the program to stop.
 - E : Error message requiring human intervention to fix the problem.
 - W : Warning message reporting an error that is automatically fixed by the program.
 - I : Information message is for key milestones achievements.
 - _ : (white space) It's just a normal information message.
 - . : Debug message of level 0
 - - : Debug message of level 1
 - = : Debug message of level 2

The command `clogClr` accept as optional argument a log file to read, otherwise it
will read logging messages from stdin. It will then colorize lines according to the 
logging level. Executing the command `testClog | clogClr` will demonstrate its 
effect. 

I handy usage pattern when logging messages are written to files is to use the command
`tail -F log_latest | clogClr`. Where `log_latest` is the link file to the latest 
log file created by the clog package. 
