Golang Line Group Processor with Timeout

This Golang script reads lines from standard input (STDIN) and groups them according to a specified number of lines. The script then sends the full group to the specified command's standard input (stdin) and outputs both stdin and stderr. The next group will only be processed once the previous group has finished execution. Additionally, a timeout flag is provided to limit the command execution time.

Usage:

cat file-with-lines.txt | go run main.go -l <lines> [-t <timeout>] <command> <command arguments>

Flags:
  -l <lines>      : Number of lines to group
  -t <timeout>    : Timeout value for command execution (optional)

How to Build:

To build the program, simply run the following command:

go build main.go

This will create an executable named 'main' in the current directory.