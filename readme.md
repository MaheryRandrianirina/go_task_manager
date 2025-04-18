# Go Task Manager

Go Task Manager is a simple task management application built with Go. It allows users to create, update, delete, and list tasks efficiently.

## Features

- Add new tasks with descriptions and categories.
- Update tasks
- Delete tasks.
- List all tasks or filter by category.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/MaheryRandrianirina/go_task_manager.git
   ```
2. Navigate to the project directory:
   ```bash
   cd go_task_manager
   ```
3. Build the project:
   ```bash
   go build
   ```

## Usage

Run the application:
```bash
./go_task_manager
```

### Commands
- **Add Task**: `gtm -n <task name>" -c <task category> -d <task description>` (description is optional)
- **List Tasks**: `gtm -l` (adding `-n <number>` allows to display only <number> tasks)
- **Filter Tasks**. `gtm -l -c <task category>`
- **Delete Task**: `gtm -d <task_id>`
- **Update Task**: `gtm -u <task_id> <task new name> <task new category> <task new description>` (At least one arg should be provided. You can let an element empty if it will not be changed.)

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## TO DO
- [ ] Add due date to task
- [ ] Mark tasks as completed.
- [ ] List all tasks or filter by status (completed/pending)
## Contact

For questions or suggestions, please contact [maheryrandrianirina@gmail.com].