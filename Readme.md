## TodoApp (Go CLI)

A simple cross‑platform command‑line Todo application written in Go. Todos are stored locally in a JSON file (`todos.json`) and displayed as a formatted table.

### Features
- **Add**: Create new todos with a title
- **List**: Pretty table output with creation and completion timestamps
- **Edit**: Rename a todo by index
- **Toggle**: Mark a todo as completed/incomplete (tracks completion time)
- **Delete**: Remove a todo by index
- **Local storage**: JSON file; no external DB required

## Requirements
- **Go**: 1.22+ (as specified in `go.mod`)
- **OS**: Windows, macOS, or Linux

## Install / Build
Clone this repository and build the binary:

```bash
git clone https://github.com/marshallfungai/todo-tracker-app
cd todoApp
go build -o todo
```

Run directly with Go (useful during development):

```bash
go run . -list
```

## Usage
Flags are parsed using Go's `flag` package. Use `-h` for built-in help.

```bash
./todo -h
```

### Commands
- **-add "TITLE"**: Add a new todo with the given title
- **-list**: Show all todos
- **-edit INDEX:NEW_TITLE**: Rename a todo
- **-toggle INDEX**: Toggle completion state
- **-del INDEX**: Delete a todo

### Examples
Add two todos, list them, toggle, edit, and delete:

```bash
./todo -add "Write project README"
./todo -add "Refactor storage layer"
./todo -list

./todo -toggle 0
./todo -edit 1:"Refactor persistence layer"
./todo -del 0
./todo -list
```

Example list output:

```
#  Title                         Completed  Created At           Completed At
0  Write project README         Yes        2025-08-09 12:55:14  Sat, 09 Aug 2025 13:01:59 EAT
1  Refactor persistence layer   No         2025-08-09 12:59:10  N/A
```

## Data persistence
- Data is stored in a JSON file named `todos.json` in the working directory.
- The file is created automatically the first time you save (e.g., after the first successful command that modifies todos).
- To change the storage path/name, update the filename in `main.go`:

```go
storage := NewStorage[Todos]("todos.json")
```

### JSON format
```json
[
  {
    "Title": "Write project README",
    "Completed": false,
    "CreatedAt": "2025-08-09T12:55:14.629642844+03:00",
    "CompletedAt": null
  }
]
```

## Development
- Ensure dependencies are in place: `go mod tidy`
- Run locally: `go run . -list`
- Build: `go build -o todo`

## Notes
- The app currently uses a local JSON file for storage (not sqlite3).
- Time values are stored in RFC3339 format in the JSON file and displayed human‑readable in the table.