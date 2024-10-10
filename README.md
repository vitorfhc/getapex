# getapex

`getapex` is a simple command-line tool written in Go that extracts the apex domain from a given URL. It supports reading input URLs from standard input and outputs the apex domain, which is the primary registered domain name.

## Installation

To install `getapex`, make sure you have Go installed and run:

```bash
go install github.com/vitorfhc/getapex@latest
```

## Usage

`getapex` reads URLs from standard input, so you can pipe in a list of URLs or type them in manually:

```bash
echo "example.com" | getapex
```

Output:
```
example.com
```

You can also pass multiple URLs from a file:

```bash
cat file.txt | getapex
```

Output:
```
example.com
domain.com
google.com
```
