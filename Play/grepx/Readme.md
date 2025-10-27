# GREPX just yet another grep implementation

CLi - like grep filter. just reading the file or Stdin and printing the output where found substring, supporting flags.

## Features:
- Input: `grepx -i -n -f /path/to/file 'foo'` or `.... < file`
- Flags:
  - -i - case insensitive
  - -n - print line number
  - -f - read from file (if there is no file then reading from Stdin)
- Exit code: 0 if found, 1 if not found, 2 - input erros or wrong arguments

