# Unix Word Count Utility in Go

`wc` command in Unix systems, implemented in Go.

## Steps to use the project:

1. `make build` &rarr; creates the `go-wc` executable
2. `sudo ln -s "$PWD/go-wc" /usr/local/bin/go-wc` &rarr; creates a symbolic link to the `go-wc` executable in `$PATH` so we don't have to use the `./` prefix everytime we run the app.
3. `go-wc [-c | -l | -w] <file-name>` or `cat <file-name> | go-wc [-c | -l | -w]` or simply `go-wc [-c | -l | -w]` and you can type directly to the terminal, and the app will start counting words.
4. Optional flags:
   * `-c` &rarr; Prints bytes in the given stream
   * `-l` &rarr; Prints lines in the given stream
   * `-w` &rarr; Prints words in the given stream
