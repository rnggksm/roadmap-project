# Expense Tracker CLI

A simple command-line tool to track your expenses. Built with Go and [Cobra](https://cobra.dev/).

Project URL: https://roadmap.sh/projects/expense-tracker

## Features

- ✅ Add expenses with description and amount
- ✅ List all expenses with formatted table output
- ✅ View summary of total expenses
- ✅ Filter summary by month
- ✅ Update expenses
- ✅ Delete expenses
- ✅ Persistent storage in JSON format

## Building

```bash
go build -o build/expense-tracker
```

Or build directly from the build directory:

```bash
go build -o expense-tracker
```

## Usage

### Add an expense

```bash
./expense-tracker add --description "Lunch" --amount 20
```

Short flags:
```bash
./expense-tracker add -d "Lunch" -a 20
```

### List all expenses

```bash
./expense-tracker list
```

Output:
```
ID | Date       | Description | Amount
---|------------|-------------|-------
1  | 2026-05-14 | Lunch       | 20.00
2  | 2026-05-14 | Dinner      | 15.50
```

### View expense summary

Show total expenses:
```bash
./expense-tracker summary
```

Output:
```
Total expenses: $35.50
```

Filter by month:
```bash
./expense-tracker summary --month 5
```

Short flag:
```bash
./expense-tracker summary -m 5
```

Output:
```
Total expenses for May: $35.50
```

### Update an expense

```bash
./expense-tracker update --id 1 --description "Lunch updated" --amount 25
```

### Delete an expense

```bash
./expense-tracker delete --id 2
```

Short flag:
```bash
./expense-tracker delete -i 2
```

## Data Storage

Expenses are stored in `data.json` in the working directory. The file is automatically created on first run.

Example `data.json`:
```json
[
  {
    "id": 1,
    "description": "Lunch",
    "amount": 20,
    "created_at": "2026-05-14T20:53:40+07:00",
    "updated_at": "2026-05-14T20:53:40+07:00"
  }
]
```

## Project Structure

```
expense-tracker/
├── cmd/           # Cobra commands
├── handler/       # Business logic handlers
├── internal/
│   └── service/   # Core service layer
├── model/         # Data models
├── build/         # Build output
├── main.go        # Entry point
├── go.mod         # Go module
└── README.md      # This file
```

## Development

### Dependencies

- Go 1.26+
- [Cobra CLI Framework](https://github.com/spf13/cobra)

Install dependencies:
```bash
go get -u github.com/spf13/cobra@latest
```

### Architecture

- **cmd/**: Cobra command definitions and CLI wiring
- **handler/**: Translates CLI input to service calls
- **internal/service/**: Core business logic and persistence
- **model/**: Data structures

## License

Built as part of the roadmap.sh projects.
