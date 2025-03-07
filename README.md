# Extract a package's structs, functions, or global variables.

Extract structs from a directory to one file.

| Flag | Description                     | Usage                                                |
| :--- | :------------------------------ | :--------------------------------------------------- |
| `-d` | Directory to extract.           | `-d relative/path/to/dir` (`/...` for nested search) |
| `-p` | Set the output package.         | `-p xstruct` (`xstruct` by default)                  |
| `-s` | Enable sorting.                 | `-s`                                                 |
| `>`  | Pipe standard output to a file. | `> file.go`                                          |