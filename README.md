# random-string-generator

Regex string generator

## How it works

Give this code a random regex expression and let it generate a string matching the pattern for you.

## Running the code

Build the binary using the makefile

```bash
$ make build
```

Run the binary

```bash
$ ./cmd/main <regex>
```

### Example

```bash
$ ./cmd/main "/[a-z]{3,5}/"
vlbz
```
# random-go-string-generator
