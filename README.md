# Task Tracker CLI

A simple command-line task tracker built with Go. Tasks are stored locally in `data.json`.

Project reference: https://roadmap.sh/projects/task-tracker

## Requirements

- Go 1.20+ (or your installed Go version from `go.mod`)

## Run the app

Run from the project root:

```bash
go run . <command>
```

Example:

```bash
go run . add "Buy groceries"
```

## Build binary

Build into the `build` directory:

```bash
mkdir -p build
go build -o build/task-cli .
```

Run the binary:

```bash
./build/task-cli add "Buy groceries"
./build/task-cli list
```

## Commands

### Add a task

```bash
go run . add "Task description"
```

### Update task description

```bash
go run . update <id> "New description"
```

### Delete a task

```bash
go run . delete <id>
```

### Mark task in progress

```bash
go run . mark-in-progress <id>
```

### Mark task done

```bash
go run . mark-done <id>
```

### List all tasks

```bash
go run . list
```

### List tasks by status

```bash
go run . list todo
go run . list in-progress
go run . list done
```

## Data file

- File: `data.json`
- Auto-created on first run if it does not exist
- Starts as an empty list (`[]`)
- Task IDs start from `1`

## Notes

- Always use `go run .` so all files in the package are compiled.
- Status values used by this app are `todo`, `in-progress`, and `done`.
