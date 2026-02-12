# Todo CLI

A powerful, feature-rich command-line todo application built with Go.

## Features

- ✅ **Task Management**
  - Add tasks with title and description
  - Edit existing tasks
  - Delete tasks
  - Mark tasks as complete/incomplete

- 📋 **Viewing & Organization**
  - List all tasks
  - Filter by status (pending, completed, overdue)
  - Filter by priority, project, or tags
  - Sort by priority, date, or creation time
  - Search tasks by keyword

- 🏷️ **Categorization**
  - Assign priorities (low, medium, high)
  - Set due dates and deadlines
  - Categorize with tags
  - Organize by projects

- 💾 **Data Management**
  - Local storage (JSON or YAML)
  - Automatic backup
  - Restore from backups

- 🎨 **User Experience**
  - Color-coded output
  - Interactive mode with arrow key navigation
  - Command aliases for quick access
  - Tab completion support

## Installation

### Option 1: Install via Homebrew (macOS)

The easiest way to install on macOS:

```bash
# Add the tap
brew tap gourangadassamrat/tap

# Install the app
brew install todo-cli
```

That's it! You can now run the app using the `todo` command.

### Option 2: Download Pre-built Binaries

Download the latest release for your platform from the [Releases page](https://github.com/GourangaDasSamrat/todo-cli-go/releases).

Available platforms:

- **Linux** - `todo-{version}-linux-amd64.tar.gz`
- **macOS (Intel)** - `todo-{version}-darwin-amd64.tar.gz`
- **macOS (Apple Silicon)** - `todo-{version}-darwin-arm64.tar.gz`
- **Windows** - `todo-{version}-windows-amd64.zip`

#### Installation Steps:

```bash
# Linux/macOS: Extract and install
tar -xzf todo-v1.0.0-linux-amd64.tar.gz
sudo mv todo-linux-amd64 /usr/local/bin/todo
chmod +x /usr/local/bin/todo

# Windows: Extract the zip and add to PATH
```

### Option 3: Build from Source

#### Prerequisites

- Go 1.21 or higher

#### Build

```bash
git clone https://github.com/GourangaDasSamrat/todo-cli-go.git
cd todo-cli-go
go mod download
go build -o todo
```

#### Install Globally

```bash
go install
```

Or copy the binary to your PATH:

```bash
# Linux/macOS
sudo cp todo /usr/local/bin/

# Windows
# Copy todo.exe to a directory in your PATH
```

## Quick Start

See [docs/Quickstart.md](docs/Quickstart.md) for a comprehensive quick start guide.

### Add a Task

```bash
# Simple task
todo add --title "Buy groceries"

# Task with details
todo add \
  --title "Complete project proposal" \
  --description "Write and submit Q1 proposal" \
  --priority high \
  --project "Work" \
  --tags "urgent,deadline" \
  --due "2024-03-15 17:00"
```

### List Tasks

```bash
# List all tasks
todo list

# Filter by status
todo list --status pending
todo list --status completed
todo list --status overdue

# Filter by priority
todo list --priority high

# Sort tasks
todo list --sort priority
todo list --sort date --asc
```

### Edit a Task

```bash
todo edit --id <task-id> --title "New title" --priority medium
```

### Complete a Task

```bash
# Mark as complete
todo complete --id <task-id>

# Mark as incomplete
todo complete --id <task-id> --incomplete
```

### Search Tasks

```bash
todo search "grocery"
```

### Delete a Task

```bash
todo delete --id <task-id>
```

### Backup & Restore

```bash
# Create backup
todo backup

# List backups
todo restore

# Restore from backup
todo restore tasks_backup_2024-03-10_14-30-00.json
```

### Interactive Mode

```bash
todo interactive
```

This launches a menu-driven interface where you can:

- Navigate with arrow keys
- Select options with Enter
- Manage tasks interactively

## Command Reference

### Global Flags

- `--storage, -s`: Storage type (json or yaml) - default: json

### Commands and Aliases

| Command       | Aliases                 | Description                   |
| ------------- | ----------------------- | ----------------------------- |
| `add`         | `a`, `new`, `create`    | Add a new task                |
| `list`        | `ls`, `l`               | List tasks                    |
| `edit`        | `e`, `update`, `modify` | Edit a task                   |
| `delete`      | `del`, `rm`, `remove`   | Delete a task                 |
| `complete`    | `done`, `finish`, `c`   | Mark task complete/incomplete |
| `search`      | `find`, `s`             | Search tasks                  |
| `backup`      | `bak`                   | Create backup                 |
| `restore`     | `rst`                   | Restore from backup           |
| `interactive` | `i`, `int`              | Interactive mode              |

### Add Command Flags

```
--title, -t        Task title (required)
--description, -d  Task description
--priority, -p     Priority: low, medium, high (default: low)
--project          Project name
--tags             Comma-separated tags
--due              Due date (YYYY-MM-DD or YYYY-MM-DD HH:MM)
```

### List Command Flags

```
--status           Filter by status: pending, completed, overdue
--priority         Filter by priority: low, medium, high
--project          Filter by project name
--tags             Filter by tags
--sort             Sort by: priority, date, created, title
--asc              Sort in ascending order
```

### Edit Command Flags

```
--id, -i           Task ID (required)
--title, -t        New title
--description, -d  New description
--priority, -p     New priority
--project          New project
--tags             New tags
--due              New due date (empty to clear)
```

### Complete Command Flags

```
--id, -i           Task ID (required)
--incomplete, -u   Mark as incomplete instead of complete
```

## Examples

See [docs/EXAMPLES.md](docs/EXAMPLES.md) for detailed usage examples.

### Daily Workflow

```bash
# Morning: See what's pending
todo list --status pending --sort priority

# Add a new task
todo add -t "Review pull requests" -p high --due "2024-03-10 15:00"

# Complete a task
todo complete --id 1678901234-abc123

# End of day: Check completed tasks
todo list --status completed
```

### Project Management

```bash
# Add project tasks
todo add -t "Design database schema" --project "API Redesign" -p high
todo add -t "Write API documentation" --project "API Redesign" -p medium

# View project tasks
todo list --project "API Redesign" --sort priority

# Filter by tags
todo list --tags "urgent"
```

### Search and Filter

```bash
# Search for tasks
todo search "meeting"

# View overdue tasks
todo list --status overdue --sort date

# High priority pending tasks
todo list --status pending --priority high
```

## Project Structure

```
todo-cli-go/
├── .github/
│   └── workflows/
│       └── release.yml          # Automated release workflow
├── cmd/                         # CLI commands
│   ├── root.go                 # Root command
│   ├── add.go                  # Add task command
│   ├── list.go                 # List tasks command
│   ├── edit.go                 # Edit task command
│   ├── delete.go               # Delete task command
│   ├── complete.go             # Complete task command
│   ├── search.go               # Search command
│   ├── backup.go               # Backup/restore commands
│   └── interactive.go          # Interactive mode
├── docs/
│   ├── CONTRIBUTING.md         # Contribution guidelines
│   ├── EXAMPLES.md             # Detailed examples
│   ├── Project_Summary.md      # Project overview
│   └── Quickstart.md           # Quick start guide
├── internal/
│   ├── models/                 # Data models
│   │   └── task.go
│   ├── storage/                # Storage implementations
│   │   └── storage.go
│   ├── ui/                     # UI utilities
│   │   ├── ui.go
│   │   └── interactive.go
│   └── utils/                  # Utility functions
│       └── utils.go
├── pkg/
│   ├── filter/                 # Task filtering
│   │   └── filter.go
│   └── sort/                   # Task sorting
│       └── sort.go
├── main.go                     # Application entry point
├── go.mod                      # Go module definition
├── go.sum                      # Go dependencies checksums
├── Makefile                    # Build automation
├── LICENSE                     # MIT License
└── README.md                   # This file
```

## Data Storage

Tasks are stored in `~/.todo-cli/` directory:

- `tasks.json` or `tasks.yaml` - Main task data
- `backups/` - Backup files with timestamps

## Color Coding

- **Red** - High priority tasks, overdue tasks
- **Yellow** - Medium priority tasks
- **Blue** - Low priority tasks
- **Green** - Completed tasks

## Automated Releases & CI/CD

This project uses GitHub Actions for fully automated releases. The release workflow is triggered automatically when you push a version tag.

### What Happens on Release

When you push a tag (e.g., `v1.0.0`), the GitHub Action automatically:

1. ✅ **Runs Tests** - Ensures all tests pass before building
2. 🔨 **Builds Binaries** for all platforms:
   - Linux (amd64)
   - macOS Intel (amd64)
   - macOS Apple Silicon (arm64)
   - Windows (amd64)
3. 📦 **Creates Release Archives** - `.tar.gz` for Unix, `.zip` for Windows
4. 📝 **Generates Changelog** - From git commits since last release
5. 🚀 **Creates GitHub Release** - With all binaries attached
6. 🍺 **Updates Homebrew Tap** - Automatically updates the formula in [gourangadassamrat/homebrew-tap](https://github.com/gourangadassamrat/homebrew-tap)
7. 📤 **Uploads Artifacts** - Stores build artifacts for 7 days

### How to Create a Release

Simply push a tag following semantic versioning:

```bash
# Create a new version tag
git tag -a v1.0.0 -m "Release version 1.0.0"

# Push the tag to GitHub
git push origin v1.0.0
```

That's it! The CI/CD pipeline handles everything else automatically.

### Version Format

Tags must follow the format: `v*.*.*` (e.g., `v1.0.0`, `v1.2.3`, `v2.0.0-beta.1`)

### Release Artifacts

Each release includes:

- `todo-{version}-linux-amd64.tar.gz`
- `todo-{version}-darwin-amd64.tar.gz` (macOS Intel)
- `todo-{version}-darwin-arm64.tar.gz` (macOS Apple Silicon)
- `todo-{version}-windows-amd64.zip`
- Release notes with changelog

### Homebrew Formula Update

The workflow uses [mislav/bump-homebrew-formula-action](https://github.com/mislav/bump-homebrew-formula-action) to automatically:

- Calculate SHA256 checksum of the macOS binary
- Update the formula file in the tap repository
- Commit changes with a descriptive message

**Note**: Requires `HOMEBREW_TAP_TOKEN` secret to be set in repository settings.

### Setting Up Automated Releases

If you're forking this project, you'll need to:

1. **Create a Personal Access Token**:
   - Go to GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
   - Generate new token with `repo` scope
   - Copy the token

2. **Add Token to Repository Secrets**:
   - Go to Repository Settings → Secrets and variables → Actions
   - Create new secret: `HOMEBREW_TAP_TOKEN`
   - Paste your token

3. **Create Homebrew Tap Repository**:
   - Create a repository named `homebrew-tap`
   - Add a `Formula` directory
   - The workflow will automatically create/update the formula

For more details, see [`.github/workflows/release.yml`](.github/workflows/release.yml)

## Contributing

We welcome contributions! Please see [docs/CONTRIBUTING.md](docs/CONTRIBUTING.md) for guidelines.

### Quick Contribution Steps

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development

```bash
# Clone the repository
git clone https://github.com/GourangaDasSamrat/todo-cli-go.git
cd todo-cli-go

# Install dependencies
go mod download

# Run tests
go test -v ./...

# Build
go build -o todo

# Run
./todo --help
```

### Using Make

```bash
# Build the application
make build

# Run tests
make test

# Clean build artifacts
make clean

# Install locally
make install
```

### Release Guidelines

- Use semantic versioning (v1.0.0, v1.1.0, v2.0.0)
- Write clear, descriptive commit messages
- Update documentation for new features
- Add tests for new functionality
- Test locally before creating a release tag

## Documentation

- [Quick Start Guide](docs/Quickstart.md) - Get started in 5 minutes
- [Examples](docs/EXAMPLES.md) - Detailed usage examples
- [Contributing](docs/CONTRIBUTING.md) - How to contribute
- [Project Summary](docs/Project_Summary.md) - Project overview

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) file for details.

## Author

Created with ❤️ by [Gouranga Das Samrat](https://github.com/GourangaDasSamrat)

## Support

- 🐛 **Bug Reports**: [Open an issue](https://github.com/GourangaDasSamrat/todo-cli-go/issues)
- 💡 **Feature Requests**: [Open an issue](https://github.com/GourangaDasSamrat/todo-cli-go/issues)
- 💬 **Questions**: [Discussions](https://github.com/GourangaDasSamrat/todo-cli-go/discussions)

## Acknowledgments

Built with:

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework (if used)
- [Go](https://golang.org/) - Programming language

---

**Star ⭐ this repository if you find it helpful!**
