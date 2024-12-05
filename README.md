# Advent of Code 2024

My solutions for [Advent of Code 2024](https://adventofcode.com/2024) implemented in Go.

## Project Structure

```
aoc-24/
├── Makefile        # Build and run commands
├── inputs/         # Daily puzzle inputs
│   └── day01.txt
├── src/           # Solution implementations
│   └── day01.go
└── README.md
```

## Requirements

- Go 1.23.2 or higher
- Make

## Running Solutions

You can run solutions in several ways:

### Run a specific day

```bash
# Run day 1
make run-1
# or
make run-01

# Run day 15
make run-15
```

### Run all solutions

```bash
make all
```

### Clean build artifacts

```bash
make clean
```

## Solution Structure

Each day's solution follows this pattern:
- Source file: `src/dayXX.go`
- Input file: `inputs/dayXX.txt`

Solutions are organized as standalone Go programs that read their input from the corresponding file in the `inputs` directory.

## Help

Run `make help` to see available commands.

