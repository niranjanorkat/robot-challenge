# 🦾 Robot Simulator CLI

A command-line simulator to manage warehouses and robots, allowing you to enqueue tasks, move robots, manage crates, and view a visual grid representation of the warehouse.

---

## 📦 Features

- Add and view **normal** or **crate-enabled** warehouses.
- Add robots of various types (`N` for normal, `D` for diagonal).
- Assign movement tasks to robots (e.g. `N`, `E`, `S`, `W`, `G`, `D`).
- View the current robot and crate positions using a visual **grid**.
- Cancel active robot tasks and view task status.
- Show all robots and crates in a warehouse.

---

## 🛠️ Usage

Run the REPL using:

```bash
go run main.go
```

You'll enter a prompt like this:

```bash
Robot Simulator CLI
Commands: add_warehouse, show_warehouse, ...
rcli>
```

---

## 🧾 Command Reference

### Warehouse Management

| Command                        | Description                                      |
|-------------------------------|--------------------------------------------------|
| `add_warehouse [type]`        | Add warehouse of type `n` (default) or `c`       |
| `show_warehouse`              | Show all added warehouses                        |

### Robot Management

| Command                            | Description                                      |
|------------------------------------|--------------------------------------------------|
| `add_robot W<id> [type]`           | Add robot to warehouse. Type: `N` or `D`         |
| `show_robots W<id>`                | List robots in a warehouse                       |
| `move_robot W<id> R<id> <cmds>`    | Send movement commands (e.g. `N S E E G`)            |

### Crate Management

| Command                             | Description                                         |
|-------------------------------------|-----------------------------------------------------|
| `add_crate W<id> x y`               | Add crate to `(x, y)` of crate-enabled warehouse    |
| `show_crates W<id>`                 | List crates in warehouse                           |

### Grid & Visualization

| Command             | Description                             |
|---------------------|-----------------------------------------|
| `show_grid W<id>`   | Display robot and crate grid            |

### Task Management

| Command                              | Description                                      |
|--------------------------------------|--------------------------------------------------|
| `show_tasks W<id> R<id>`             | View active tasks for a robot                    |
| `cancel_task W<id> R<id> <task_id>`  | Cancel a specific robot task                    |

### System

| Command    | Description       |
|------------|-------------------|
| `exit`     | Exit the program  |

---

## 📌 Notes

- Warehouses are 10x10 grids.
- Only crate-enabled warehouses support `add_crate` or display crate grids.
- Diagonal robots support diagonals like `NE`, `SW` if implemented.
- Commands like `G` (grab crate) or `D` (drop crate) only work on crate-enabled grids.

---

## 🧪 Example

```bash
rcli> add_warehouse c
Crate-enabled warehouse added.
rcli> add_robot W1 N
Robot added to warehouse.
rcli> move_robot W1 R1 NEEGDD
Move command sent successfully.
rcli> show_grid W1
Robot Grid:
. . R1 . . . . . . .
...

Crate Grid:
. . C . . . . . . .
...
```

---
