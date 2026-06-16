# 📌 Task CLI

A lightweight and simple command-line task manager built for productivity and speed.

`task-cli` allows you to create, update, delete, and track tasks directly from your terminal with an intuitive and minimal interface.

---

## ✨ Features

- ➕ Add new tasks
- ✏️ Update existing tasks
- ❌ Delete tasks
- 🚧 Mark tasks as **in progress**
- ✅ Mark tasks as **done**
- 📋 List all tasks
- 🔍 Filter tasks by status

---

# 🚀 Installation

Clone the repository:

```bash
git clone https://github.com/your-username/task-cli.git
cd task-cli
```

Build the project:

```bash
go build -o task-cli
```

Run the executable:

```bash
./task-cli
```

---

# 📖 Usage

```bash
task-cli <command> [arguments]
```

---

# 🛠 Commands

| Command                     | Description                                                   |
| --------------------------- | ------------------------------------------------------------- |
| `add <description>`         | Add a new task                                                |
| `update <id> <description>` | Update an existing task description                           |
| `delete <id>`               | Delete a task                                                 |
| `mark-in-progress <id>`     | Mark a task as `in-progress`                                  |
| `mark-done <id>`            | Mark a task as `done`                                         |
| `list [status]`             | List tasks _(Optional status: `todo`, `in-progress`, `done`)_ |

---

# 💡 Examples

## Add a Task

```bash
task-cli add "Buy groceries"
```

### Output

```bash
Task added successfully (ID: 1)
```

---

## Update a Task

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

---

## Delete a Task

```bash
task-cli delete 1
```

---

## Mark Tasks

### Mark as In Progress

```bash
task-cli mark-in-progress 1
```

### Mark as Done

```bash
task-cli mark-done 1
```

---

# 📋 Listing Tasks

## List All Tasks

```bash
task-cli list
```

---

## List Tasks by Status

### Done Tasks

```bash
task-cli list done
```

### Todo Tasks

```bash
task-cli list todo
```

### In Progress Tasks

```bash
task-cli list in-progress
```

---

# 🧩 Example Workflow

```bash
# Adding a new task
task-cli add "Buy groceries"

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

---

# 📁 Project Structure

```text
task-cli/
├── main.go
├── todo.json
├── README.md
└── ...
```

## 🌍 Project Inspiration

This project is based on the Task Tracker challenge from Roadmap.sh:

👉 https://roadmap.sh/projects/task-tracker

It is a great backend beginner project focused on:

- CLI application development
- File-based data persistence
- CRUD operations
- Task state management
- Clean software architecture

---

# 🤝 Contributing

Contributions, issues, and feature requests are welcome.

Feel free to fork the repository and submit a pull request.

---

# 📜 License

This project is licensed under the MIT License.

---

# ⭐ Support

If you found this project useful, consider giving it a star on GitHub ⭐
