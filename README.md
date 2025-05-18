# Push-Swap

A sorting algorithm project that efficiently sorts a stack of integers using a limited set of operations.

## Overview

This project implements two programs:

1. **push_swap**: Calculates and outputs the minimum sequence of operations needed to sort a stack of integers in ascending order.
2. **checker**: Verifies if a given sequence of operations correctly sorts a stack of integers.

The goal is to sort a stack of integers using two stacks (a and b) and a specific set of allowed operations, while minimizing the number of operations used.

## Allowed Operations

| Operation | Description |
|-----------|-------------|
| `pa` | Push the top element from stack b to stack a |
| `pb` | Push the top element from stack a to stack b |
| `sa` | Swap the first 2 elements of stack a |
| `sb` | Swap the first 2 elements of stack b |
| `ss` | Execute both `sa` and `sb` |
| `ra` | Rotate stack a (shift up all elements by 1) |
| `rb` | Rotate stack b (shift up all elements by 1) |
| `rr` | Execute both `ra` and `rb` |
| `rra` | Reverse rotate stack a (shift down all elements by 1) |
| `rrb` | Reverse rotate stack b (shift down all elements by 1) |
| `rrr` | Execute both `rra` and `rrb` |

## Algorithm

The sorting algorithm implemented in this project is based on:
- Using stack b as a temporary storage area
- Partitioning the data using strategic pivots
- Optimizing the number of operations for different ranges of input sizes

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/push-swap.git
cd push-swap

# Build the project
go build -o push-swap cmd/push-swap/main.go
go build -o checker cmd/checker/main.go
```

## Usage

### push_swap

```bash
# Format: ./push-swap [integers]
./push-swap "4 67 3 87 23"
```

This will output the sequence of instructions needed to sort the stack.

### checker

```bash
# Format: ./checker [integers]
# The checker will wait for instructions from stdin
./checker "4 67 3 87 23"
```

You can also pipe the output of push_swap into checker:

```bash
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
```

If the checker displays "OK", it means the sequence of operations sorted the stack correctly.

## Examples

```bash
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | wc -l
6

$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```

## Error Handling

Both programs display "Error" followed by a newline in case of:
- Non-integer arguments
- Duplicate values
- Invalid instructions (for checker)

## Project Structure

```
push-swap/
├── cmd/
│   ├── push-swap/
│   │   └── main.go
│   └── checker/
│       └── main.go
├── internal/
│   ├── stack/
│   │   └── stack.go
│   ├── operations/
│   │   └── operations.go
│   └── algorithm/
│       └── algorithm.go
├── test/
│   └── ...
├── go.mod
├── go.sum
└── README.md
```

## Performance

The algorithm's efficiency has been optimized for different input sizes:
- 3 numbers: Maximum 3 operations
- 5 numbers: Maximum 12 operations
- 100 numbers: Less than 700 operations
- 500 numbers: Less than 5500 operations

## What I Learned

Through this project, I gained experience with:
- Implementation of non-comparative sorting algorithms
- Stack data structures and manipulation
- Algorithm optimization techniques
- Go programming principles and best practices
- Error handling and input validation

## Collabs

[@Serigne Saliou Mbacke Mbaye](https://github.com/SSMM0498)
[@Louis Sebastian Malack](https://github.com/steb1)
