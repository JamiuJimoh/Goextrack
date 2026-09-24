```
╔═══════════════════════════════════════════════════════════════════════════╗
║                                                                           ║
║ ██████╗  ██████╗ ███████╗██╗  ██╗████████╗██████╗  █████╗  ██████╗██╗  ██╗║
║██═════╝ ██╔═══██╗██╔════╝╚██╗██╔╝╚══██╔══╝██╔══██╗██╔══██╗██╔════╝██╗ ██╔╝║
║██║  ███╗██║   ██║█████╗   ╚███╔╝    ██║   ██████╔╝███████║██║     █████╔╝ ║
║██║   ██║██║   ██║██╔══╝   ██╔██╗    ██║   ██╔══██╗██╔══██║██║     ██╔═██╗ ║
║╚██████╔╝╚██████╔╝███████╗██╔╝ ██╗   ██║   ██║  ██║██║  ██║╚██████╗██║  ██╗║
║ ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝║
║                                                                           ║
╚═══════════════════════════════════════════════════════════════════════════╝
```

A simple, interactive command-line expense tracker built with Go.

GOEXTRACK allows users to record, view, and manage expenses directly from the terminal.

## Features

- Add new expenses interactively
- List recorded expenses
- Retrieve expenses by ID
- Delete expenses by ID
- Persist expense data locally
- Interactive command-line interface

> Features may evolve as development continues.

## Requirements

- [Go](https://go.dev/dl/) 1.27 or later

## Installation

Install

```bash
go install github.com/JamiuJimoh/Goextrack/cmd/goextrack-repl@latest
```
and run
```bash
goextrack-repl
```

Or

Clone the repository

```bash
git clone https://github.com/JamiuJimoh/Goextrack.git
```

Navigate into the project directory:

```bash
cd Goextrack
```

## Running the Application

Run the repl application directly with:

```bash
go run ./cmd/goextrack-repl/main.go
```
or

```bash
make repl-run
```

Alternatively, build an executable:

```bash
go build -o goextrack ./cmd/goextrack-repl/main.go 
```
or

```bash
make repl-build
```

Then run it:

```bash
./goextrack
```

## Usage

After launching GOEXTRACK, use the available commands to interact with your expenses.

Example commands:

```text
add       Add a new expense
list      List all expenses
get       Get an expense by ID
delete    Delete an expense by ID
help      Display the help menu
exit      Exit the application
```

### Example Session

```text
> add

Enter amount: 58963

-------------------------------------
Categories
-------------------------------------
1. food
2. transport
3. housing
4. utilities
5. health
6. education
7. entertainment
8. shopping
9. other

Enter category: 4
Enter description: electricity bill

You have added 1 new entries. Do you want to:
1. Add more
2. Save

> 2

Expenses saved successfully.

> list

ID                                    CATEGORY       AMOUNT     DESCRIPTION             DATE
──                                    ────────       ──────     ───────────             ────
01a0ce9c-3178-7f5c-8e4d-5fffa79f237a  utilities      ₦58963.00  electricity bill        2026-09-23
```

## Development

Run the test suite:

```bash
go test ./...
```
