# Simple logging system

This simple logging system provides logging levels, file storage and file rotation.
It also provides a command to display colorized logs. 

By default, the logging output is printed to stdout.

A command to display colorized clog logs is provided. It can be tested with the
command `go install ./... && testClog | clogClr`.

It can display a clog file content with the command `clogClr logFile`. 

