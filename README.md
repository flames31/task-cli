# Task Tracker CLI (Go)

**Task Tracker CLI** is a simple command-line application to track and manage your daily tasks. Built in Go using urfave/cli library, it helps you practice core programming skills like file I/O, CLI argument parsing, and JSON data management — while creating a practical tool for task tracking.

---

## Features

- Add new tasks
- Update or delete tasks
- Mark tasks as `todo`, `in-progress`, or `done`
- List all tasks, or filter by status
- All task data is stored in a local `tasks.json` file
- Zero external libraries used

---

## Task Structure

Each task includes the following properties:

```json
{
  "id": 1,
  "description": "Buy groceries",
  "status": "todo",
  "createdAt": "2024-05-21T12:00:00Z",
  "updatedAt": "2024-05-21T12:00:00Z"
}
