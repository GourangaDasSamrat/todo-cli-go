# Todo CLI - Usage Examples

This document provides practical examples for common use cases.

## Basic Task Management

### Creating Tasks

```bash
# Simple task
todo add -t "Buy groceries"

# Task with priority
todo add -t "Fix critical bug" -p high

# Detailed task
todo add \
  -t "Prepare quarterly report" \
  -d "Include sales metrics, customer feedback, and growth projections" \
  -p medium \
  --project "Q1 Reports" \
  --tags "finance,reports" \
  --due "2024-03-31 17:00"
```

### Viewing Tasks

```bash
# All tasks
todo list

# Only pending tasks
todo list --status pending

# High priority tasks only
todo list --priority high

# Tasks sorted by due date
todo list --sort date

# Ascending order (oldest first)
todo list --sort created --asc

# Project-specific tasks
todo list --project "Website Redesign"

# Tasks with specific tag
todo list --tags "urgent"
```

## Advanced Filtering

### Combining Filters

```bash
# High priority pending tasks
todo list --status pending --priority high

# Overdue tasks sorted by priority
todo list --status overdue --sort priority

# Project tasks sorted by due date
todo list --project "Mobile App" --sort date
```

## Task Updates

### Editing Tasks

```bash
# Change title
todo edit -i <task-id> -t "New title"

# Update priority
todo edit -i <task-id> -p high

# Add due date
todo edit -i <task-id> --due "2024-04-01 09:00"

# Change project
todo edit -i <task-id> --project "New Project"

# Update multiple fields
todo edit -i <task-id> \
  -t "Updated title" \
  -p medium \
  --tags "feature,backend" \
  --due "2024-04-15"
```

### Task Completion

```bash
# Mark as complete
todo complete -i <task-id>

# Mark as incomplete (reopen)
todo complete -i <task-id> --incomplete
```

## Search and Discovery

### Keyword Search

```bash
# Search in title and description
todo search "meeting"

# Search for bug-related tasks
todo search "bug"

# Search for specific terms
todo search "design review"
```

## Backup and Recovery

### Creating Backups

```bash
# Create a backup
todo backup

# List available backups
todo restore
```

### Restoring Data

```bash
# Restore from specific backup
todo restore tasks_backup_2024-03-10_14-30-00.json
```

## Interactive Mode

### Starting Interactive Mode

```bash
# Launch interactive interface
todo interactive

# Or using alias
todo i
```

**Interactive Menu Options:**
1. View All Tasks
2. Add New Task
3. Edit Task
4. Delete Task
5. Mark Complete/Incomplete
6. Filter Tasks
7. Search Tasks
8. Backup Data
9. Restore Data
10. Exit

## Workflow Examples

### Morning Routine

```bash
# Check what's on the agenda
todo list --status pending --sort date

# Review overdue items
todo list --status overdue
```

### Project Planning

```bash
# Add all project tasks
todo add -t "Research competitors" --project "Market Analysis" -p high
todo add -t "Gather user feedback" --project "Market Analysis" -p medium
todo add -t "Create presentation" --project "Market Analysis" -p low

# View project roadmap
todo list --project "Market Analysis" --sort priority
```

### End of Day

```bash
# Review completed tasks
todo list --status completed

# Create daily backup
todo backup

# Check tomorrow's tasks
todo list --status pending --sort date
```

### Weekly Review

```bash
# All high priority items
todo list --priority high

# Overdue tasks requiring attention
todo list --status overdue --sort date

# Tasks by project
todo list --project "Website" --sort priority
todo list --project "Mobile" --sort priority
```

## Using Command Aliases

All commands have short aliases for faster typing:

```bash
# Add
todo a -t "Quick task"          # instead of 'add'

# List
todo l                           # instead of 'list'
todo ls --status pending         # another alias

# Edit
todo e -i 123 -t "New title"    # instead of 'edit'

# Delete
todo rm -i 123                   # instead of 'delete'
todo del -i 123                  # another alias

# Complete
todo c -i 123                    # instead of 'complete'

# Search
todo s "keyword"                 # instead of 'search'

# Interactive
todo i                           # instead of 'interactive'
```

## Storage Options

### Using JSON (default)

```bash
todo add -t "Task in JSON" --storage json
```

### Using YAML

```bash
todo add -t "Task in YAML" --storage yaml
todo list --storage yaml
```

## Pro Tips

### 1. Quick Task Addition
```bash
# Use short flags for speed
todo a -t "Quick task" -p h
```

### 2. Default Sorting
```bash
# List always shows newest first by default
todo list

# For oldest first
todo list --asc
```

### 3. Tab Completion
After installing bash completion:
```bash
# Type and press TAB
todo add --<TAB>
todo list --priority <TAB>
```

### 4. Combining with Other Tools
```bash
# Export task count
task_count=$(todo list | grep "Total:" | awk '{print $2}')
echo "You have $task_count tasks"

# Find tasks and process
todo list --status pending | grep "urgent"
```

### 5. Regular Backups
```bash
# Add to crontab for automatic daily backups
0 0 * * * /usr/local/bin/todo backup
```

## Date Formats

Supported date formats:
- `2024-03-15` (date only)
- `2024-03-15 14:30` (date and time)
- `2024-03-15 14:30:00` (full timestamp)
- `03/15/2024` (US format)
- `03/15/2024 14:30` (US format with time)

## Color Guide

When viewing tasks:
- **Red/Bold**: High priority tasks
- **Yellow**: Medium priority tasks
- **Blue**: Low priority tasks
- **Green**: Completed tasks
- **Red**: Overdue tasks (with ⚠ symbol)

## Getting Help

```bash
# General help
todo --help

# Command-specific help
todo add --help
todo list --help
todo edit --help
```