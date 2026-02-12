# Todo CLI - Architecture Documentation

## Table of Contents

1. [Overview](#overview)
2. [High-Level Architecture](#high-level-architecture)
3. [Component Design](#component-design)
4. [Data Flow](#data-flow)
5. [Storage Architecture](#storage-architecture)
6. [Package Organization](#package-organization)
7. [Design Patterns](#design-patterns)
8. [Technology Stack](#technology-stack)

---

## Overview

Todo CLI is built following **Clean Architecture** principles with clear separation of concerns, dependency inversion, and modular design. The application is structured into distinct layers that promote maintainability, testability, and scalability.

### Key Architectural Principles

- **Separation of Concerns**: Each package has a single, well-defined responsibility
- **Dependency Inversion**: High-level modules don't depend on low-level modules
- **Interface-Based Design**: Abstractions over concrete implementations
- **Single Responsibility**: Each component does one thing well
- **DRY (Don't Repeat Yourself)**: Shared functionality is extracted into reusable packages

---

## High-Level Architecture

```mermaid
graph TB
    subgraph "User Interface Layer"
        CLI[CLI Commands<br/>Cobra Framework]
        Interactive[Interactive Mode<br/>promptui]
        UI[UI Utilities<br/>Color Output]
    end

    subgraph "Business Logic Layer"
        Models[Task Models<br/>Data Structures]
        Filter[Filter Logic<br/>Task Filtering]
        Sort[Sort Logic<br/>Task Sorting]
        Utils[Utilities<br/>ID Gen, Parsing]
    end

    subgraph "Data Access Layer"
        Storage[Storage Interface]
        JSON[JSON Storage<br/>Implementation]
        YAML[YAML Storage<br/>Implementation]
    end

    subgraph "External"
        FS[File System<br/>~/.todo-cli/]
    end

    CLI --> Models
    CLI --> Filter
    CLI --> Sort
    CLI --> Storage
    CLI --> UI
    CLI --> Utils

    Interactive --> Models
    Interactive --> Storage
    Interactive --> UI
    Interactive --> Filter
    Interactive --> Sort

    Filter --> Models
    Sort --> Models

    Storage --> JSON
    Storage --> YAML

    JSON --> FS
    YAML --> FS

    style CLI fill:#e1f5ff
    style Interactive fill:#e1f5ff
    style UI fill:#e1f5ff
    style Models fill:#fff4e1
    style Filter fill:#fff4e1
    style Sort fill:#fff4e1
    style Utils fill:#fff4e1
    style Storage fill:#e8f5e9
    style JSON fill:#e8f5e9
    style YAML fill:#e8f5e9
    style FS fill:#f3e5f5
```

---

## Component Design

### 1. Command Layer (`cmd/`)

The command layer implements CLI commands using the Cobra framework. Each command is self-contained and orchestrates the necessary components to fulfill its responsibility.

```mermaid
graph LR
    subgraph "Command Layer"
        Root[root.go<br/>Root Command]
        Add[add.go<br/>Add Task]
        List[list.go<br/>List Tasks]
        Edit[edit.go<br/>Edit Task]
        Delete[delete.go<br/>Delete Task]
        Complete[complete.go<br/>Mark Complete]
        Search[search.go<br/>Search Tasks]
        Backup[backup.go<br/>Backup/Restore]
        Interactive[interactive.go<br/>Interactive Mode]
    end

    Root --> Add
    Root --> List
    Root --> Edit
    Root --> Delete
    Root --> Complete
    Root --> Search
    Root --> Backup
    Root --> Interactive

    style Root fill:#4fc3f7
    style Add fill:#81c784
    style List fill:#81c784
    style Edit fill:#81c784
    style Delete fill:#e57373
    style Complete fill:#81c784
    style Search fill:#fff176
    style Backup fill:#ba68c8
    style Interactive fill:#ff8a65
```

**Responsibilities:**

- Parse command-line arguments and flags
- Validate user input
- Orchestrate business logic components
- Handle errors and display results
- Manage user interaction flow

### 2. Model Layer (`internal/models/`)

Core data structures and business logic for tasks.

```mermaid
classDiagram
    class Task {
        +string ID
        +string Title
        +string Description
        +Priority Priority
        +Status Status
        +[]string Tags
        +string Project
        +time.Time DueDate
        +time.Time CreatedAt
        +time.Time UpdatedAt
        +time.Time CompletedAt
        +IsOverdue() bool
        +UpdateStatus()
        +MarkComplete()
        +MarkIncomplete()
    }

    class Priority {
        <<enumeration>>
        PriorityLow
        PriorityMedium
        PriorityHigh
        +String() string
    }

    class Status {
        <<enumeration>>
        StatusPending
        StatusCompleted
        StatusOverdue
        +String() string
    }

    class TaskList {
        +[]*Task Tasks
        +Add(task *Task)
        +Remove(id string) bool
        +GetByID(id string) *Task
        +UpdateAllStatuses()
    }

    Task --> Priority
    Task --> Status
    TaskList o-- Task
```

**Responsibilities:**

- Define task data structure
- Implement task-related business rules
- Handle task state transitions
- Provide task collection management

### 3. Storage Layer (`internal/storage/`)

Abstraction for data persistence with multiple implementations.

```mermaid
classDiagram
    class Storage {
        <<interface>>
        +Load() (*TaskList, error)
        +Save(taskList *TaskList) error
        +Backup() error
        +Restore(backupFile string) error
        +ListBackups() ([]string, error)
    }

    class JSONStorage {
        -string dataPath
        -string backupPath
        +NewJSONStorage() (*JSONStorage, error)
        +Load() (*TaskList, error)
        +Save(taskList *TaskList) error
        +Backup() error
        +Restore(backupFile string) error
        +ListBackups() ([]string, error)
    }

    class YAMLStorage {
        -string dataPath
        -string backupPath
        +NewYAMLStorage() (*YAMLStorage, error)
        +Load() (*TaskList, error)
        +Save(taskList *TaskList) error
        +Backup() error
        +Restore(backupFile string) error
        +ListBackups() ([]string, error)
    }

    Storage <|.. JSONStorage
    Storage <|.. YAMLStorage
```

**Responsibilities:**

- Abstract persistence implementation details
- Provide multiple storage backends (JSON, YAML)
- Handle file I/O operations
- Manage backups and restoration
- Ensure data integrity

**Design Benefits:**

- Easy to add new storage backends (e.g., SQLite, PostgreSQL)
- Testable through interface mocking
- Swappable implementations at runtime

### 4. Filter Package (`pkg/filter/`)

Reusable filtering logic for task collections.

```mermaid
graph TD
    A[Task Collection] --> B[Filter]
    B --> C{Status Filter?}
    B --> D{Priority Filter?}
    B --> E{Project Filter?}
    B --> F{Tags Filter?}
    B --> G{Keyword Filter?}
    B --> H{Date Range Filter?}

    C --> I[Filtered Tasks]
    D --> I
    E --> I
    F --> I
    G --> I
    H --> I

    style A fill:#e3f2fd
    style B fill:#fff9c4
    style I fill:#c8e6c9
```

**Filter Structure:**

```go
type Filter struct {
    Status   *models.Status
    Priority *models.Priority
    Project  string
    Tags     []string
    Keyword  string
    DateFrom time.Time
    DateTo   time.Time
}
```

**Responsibilities:**

- Apply multiple filter criteria
- Support combinatorial filtering
- Provide factory methods for common filters
- Efficient in-memory filtering

### 5. Sort Package (`pkg/sort/`)

Task sorting algorithms.

```mermaid
graph LR
    A[Unsorted Tasks] --> B{Sort By?}
    B -->|Priority| C[High → Low<br/>or Low → High]
    B -->|Due Date| D[Earliest → Latest<br/>or Latest → Earliest]
    B -->|Created At| E[Newest → Oldest<br/>or Oldest → Newest]
    B -->|Title| F[A → Z<br/>or Z → A]

    C --> G[Sorted Tasks]
    D --> G
    E --> G
    F --> G

    style A fill:#ffebee
    style B fill:#fff9c4
    style G fill:#e8f5e9
```

**Responsibilities:**

- Multiple sorting criteria
- Ascending/descending order support
- Efficient in-place sorting
- Handle edge cases (nil dates, empty fields)

### 6. UI Layer (`internal/ui/`)

User interface utilities for display and interaction.

```mermaid
graph TB
    subgraph "UI Components"
        Colors[Color Utilities<br/>Priority/Status Colors]
        Display[Display Functions<br/>PrintTask, PrintTaskList]
        Messages[Message Functions<br/>Success, Error, Warning]
        Interactive[Interactive Prompts<br/>Menu, Input, Select]
    end

    subgraph "External Libraries"
        FatihColor[fatih/color<br/>ANSI Colors]
        PromptUI[promptui<br/>Interactive Prompts]
    end

    Colors --> FatihColor
    Interactive --> PromptUI

    style Colors fill:#ffccbc
    style Display fill:#c5cae9
    style Messages fill:#b2dfdb
    style Interactive fill:#f8bbd0
```

**Responsibilities:**

- Color-coded output
- Formatted task display
- Interactive menus and prompts
- User feedback messages
- Error presentation

---

## Data Flow

### Add Task Flow

```mermaid
sequenceDiagram
    participant User
    participant CLI as CLI Command
    participant Utils
    participant Models
    participant Storage
    participant FS as File System

    User->>CLI: todo add -t "Task" -p high
    CLI->>Utils: GenerateID()
    Utils-->>CLI: Unique ID
    CLI->>Models: Create Task struct
    Models-->>CLI: Task instance
    CLI->>Storage: Load()
    Storage->>FS: Read tasks.json
    FS-->>Storage: Task data
    Storage-->>CLI: TaskList
    CLI->>Models: TaskList.Add(task)
    CLI->>Storage: Save(taskList)
    Storage->>Models: UpdateAllStatuses()
    Storage->>FS: Write tasks.json
    FS-->>Storage: Success
    Storage-->>CLI: nil (success)
    CLI->>User: Display success message
```

### List Tasks Flow

```mermaid
sequenceDiagram
    participant User
    participant CLI as CLI Command
    participant Storage
    participant Filter
    participant Sort
    participant UI
    participant FS as File System

    User->>CLI: todo list --status pending --sort priority
    CLI->>Storage: Load()
    Storage->>FS: Read tasks.json
    FS-->>Storage: Task data
    Storage->>Storage: UpdateAllStatuses()
    Storage-->>CLI: TaskList
    CLI->>Filter: Apply(tasks, status=pending)
    Filter-->>CLI: Filtered tasks
    CLI->>Sort: Sort(tasks, by=priority)
    Sort-->>CLI: Sorted tasks
    CLI->>UI: PrintTaskList(tasks)
    UI->>User: Colored table output
```

### Interactive Mode Flow

```mermaid
sequenceDiagram
    participant User
    participant Interactive
    participant UI
    participant Storage
    participant Models

    User->>Interactive: todo interactive

    loop Until Exit
        Interactive->>UI: Show menu
        UI->>User: Display options
        User->>UI: Select option
        UI-->>Interactive: Choice

        alt View Tasks
            Interactive->>Storage: Load()
            Storage-->>Interactive: TaskList
            Interactive->>UI: PrintTaskList()
            UI->>User: Display tasks

        else Add Task
            Interactive->>UI: PromptTaskInput()
            UI->>User: Prompt for details
            User->>UI: Enter task info
            UI-->>Interactive: Task data
            Interactive->>Models: Create Task
            Interactive->>Storage: Save()
            Interactive->>UI: PrintSuccess()

        else Edit/Delete/Complete
            Interactive->>UI: SelectTask()
            UI->>User: Show task list
            User->>UI: Select task
            Interactive->>Storage: Update/Delete
            Interactive->>UI: Show result
        end
    end
```

---

## Storage Architecture

### File System Structure

```mermaid
graph TB
    subgraph "~/.todo-cli/"
        Tasks[tasks.json<br/>Main task data]
        TasksYAML[tasks.yaml<br/>YAML format]

        subgraph "backups/"
            B1[tasks_backup_2024-03-10_14-30-00.json]
            B2[tasks_backup_2024-03-11_09-15-30.json]
            B3[tasks_backup_2024-03-12_18-45-12.json]
        end
    end

    style Tasks fill:#4caf50
    style TasksYAML fill:#2196f3
    style B1 fill:#ff9800
    style B2 fill:#ff9800
    style B3 fill:#ff9800
```

### Storage Implementation Strategy

```mermaid
graph TD
    A[Application Start] --> B{Check Storage Type Flag}
    B -->|--storage json| C[Initialize JSONStorage]
    B -->|--storage yaml| D[Initialize YAMLStorage]

    C --> E[Check Data Directory]
    D --> E

    E --> F{Directory Exists?}
    F -->|No| G[Create ~/.todo-cli/]
    F -->|Yes| H[Continue]
    G --> H

    H --> I{Data File Exists?}
    I -->|No| J[Create Empty TaskList]
    I -->|Yes| K[Load Existing Data]

    J --> L[Ready for Operations]
    K --> L

    style C fill:#81c784
    style D fill:#64b5f6
    style G fill:#ffb74d
    style J fill:#fff176
    style K fill:#fff176
    style L fill:#4db6ac
```

### Backup Strategy

```mermaid
flowchart LR
    A[Backup Command] --> B[Read Current Data]
    B --> C[Generate Timestamp]
    C --> D[Create Backup Filename<br/>tasks_backup_YYYY-MM-DD_HH-MM-SS.json]
    D --> E[Copy Data to Backup File]
    E --> F[Store in backups/ Directory]
    F --> G[Backup Complete]

    style A fill:#e1bee7
    style G fill:#c5e1a5
```

---

## Package Organization

### Directory Structure Philosophy

```mermaid
graph TB
    subgraph "Public API - pkg/"
        Filter[filter/<br/>Reusable filtering]
        Sort[sort/<br/>Reusable sorting]
    end

    subgraph "Private Implementation - internal/"
        Models[models/<br/>Core data structures]
        Storage[storage/<br/>Persistence layer]
        UI[ui/<br/>Display & interaction]
        Utils[utils/<br/>Helper functions]
    end

    subgraph "CLI Layer - cmd/"
        Commands[Individual commands<br/>add, list, edit, etc.]
    end

    subgraph "Entry Point"
        Main[main.go]
    end

    Main --> Commands
    Commands --> Models
    Commands --> Storage
    Commands --> UI
    Commands --> Filter
    Commands --> Sort
    Commands --> Utils

    Storage --> Models
    Filter --> Models
    Sort --> Models

    style Filter fill:#b2ebf2
    style Sort fill:#b2ebf2
    style Models fill:#fff9c4
    style Storage fill:#c8e6c9
    style UI fill:#f8bbd0
    style Utils fill:#ffccbc
    style Commands fill:#ce93d8
    style Main fill:#90caf9
```

### Package Dependencies

```mermaid
graph LR
    subgraph "Layer 1: Foundation"
        Models[models]
    end

    subgraph "Layer 2: Business Logic"
        Filter[filter]
        Sort[sort]
        Utils[utils]
    end

    subgraph "Layer 3: Infrastructure"
        Storage[storage]
        UI[ui]
    end

    subgraph "Layer 4: Application"
        Commands[cmd]
    end

    Filter --> Models
    Sort --> Models
    Storage --> Models

    Commands --> Models
    Commands --> Filter
    Commands --> Sort
    Commands --> Storage
    Commands --> UI
    Commands --> Utils

    style Models fill:#4caf50
    style Filter fill:#2196f3
    style Sort fill:#2196f3
    style Utils fill:#2196f3
    style Storage fill:#ff9800
    style UI fill:#ff9800
    style Commands fill:#9c27b0
```

**Dependency Rules:**

1. **Layer 1 (Models)**: No dependencies on other layers
2. **Layer 2**: Can depend on Layer 1 only
3. **Layer 3**: Can depend on Layers 1 and 2
4. **Layer 4**: Can depend on all lower layers

---

## Design Patterns

### 1. Repository Pattern (Storage)

```mermaid
classDiagram
    class Repository {
        <<interface>>
        +Load()
        +Save()
        +Backup()
        +Restore()
    }

    class JSONRepository {
        +Load()
        +Save()
        +Backup()
        +Restore()
    }

    class YAMLRepository {
        +Load()
        +Save()
        +Backup()
        +Restore()
    }

    Repository <|.. JSONRepository
    Repository <|.. YAMLRepository

    class Command {
        -Repository repo
    }

    Command --> Repository
```

**Benefits:**

- Abstracts data access
- Easy to swap implementations
- Testable with mocks

### 2. Strategy Pattern (Filtering & Sorting)

```mermaid
classDiagram
    class FilterStrategy {
        <<interface>>
        +Apply(tasks) []Task
    }

    class StatusFilter {
        +Apply(tasks) []Task
    }

    class PriorityFilter {
        +Apply(tasks) []Task
    }

    class KeywordFilter {
        +Apply(tasks) []Task
    }

    FilterStrategy <|.. StatusFilter
    FilterStrategy <|.. PriorityFilter
    FilterStrategy <|.. KeywordFilter
```

**Benefits:**

- Flexible filtering logic
- Composable filters
- Easy to add new filter types

### 3. Command Pattern (CLI Commands)

```mermaid
classDiagram
    class Command {
        <<interface>>
        +Execute()
    }

    class AddCommand {
        +Execute()
    }

    class ListCommand {
        +Execute()
    }

    class EditCommand {
        +Execute()
    }

    Command <|.. AddCommand
    Command <|.. ListCommand
    Command <|.. EditCommand

    class CLI {
        +RegisterCommand(cmd Command)
    }

    CLI --> Command
```

**Benefits:**

- Encapsulates actions
- Easy to add new commands
- Supports undo/redo (future enhancement)

### 4. Factory Pattern (Filter & Storage Creation)

```go
// Filter Factory
func NewStatusFilter(status Status) *Filter
func NewPriorityFilter(priority Priority) *Filter
func NewKeywordFilter(keyword string) *Filter

// Storage Factory
func NewJSONStorage() (*JSONStorage, error)
func NewYAMLStorage() (*YAMLStorage, error)
```

---

## Technology Stack

### Core Technologies

```mermaid
graph TB
    subgraph "Language & Runtime"
        Go[Go 1.21+<br/>Compiled Language]
    end

    subgraph "CLI Framework"
        Cobra[spf13/cobra<br/>Command Framework]
        Pflag[spf13/pflag<br/>Flag Parsing]
    end

    subgraph "UI Libraries"
        Color[fatih/color<br/>ANSI Colors]
        PromptUI[manifoldco/promptui<br/>Interactive Prompts]
        Readline[chzyer/readline<br/>Line Editing]
    end

    subgraph "Data Handling"
        JSON[encoding/json<br/>JSON Marshaling]
        YAML[gopkg.in/yaml.v3<br/>YAML Parsing]
    end

    Go --> Cobra
    Cobra --> Pflag
    Go --> Color
    Go --> PromptUI
    PromptUI --> Readline
    Go --> JSON
    Go --> YAML

    style Go fill:#00add8
    style Cobra fill:#326ce5
    style Color fill:#ff6f00
    style PromptUI fill:#7b1fa2
    style JSON fill:#ffa726
    style YAML fill:#26a69a
```

### Library Selection Rationale

| Library         | Purpose             | Why Chosen                                           |
| --------------- | ------------------- | ---------------------------------------------------- |
| **Cobra**       | CLI framework       | Industry standard, used by kubectl, Hugo, GitHub CLI |
| **promptui**    | Interactive prompts | Rich terminal UI, arrow key navigation               |
| **fatih/color** | Terminal colors     | Cross-platform, easy API, widely used                |
| **yaml.v3**     | YAML parsing        | Latest version, full YAML 1.2 support                |

---

## Error Handling Architecture

```mermaid
graph TD
    A[User Action] --> B{Validation}
    B -->|Invalid| C[Return Error]
    C --> D[Display Error Message]
    D --> E[Return to User]

    B -->|Valid| F[Execute Operation]
    F --> G{Success?}
    G -->|No| H[Wrap Error with Context]
    H --> D

    G -->|Yes| I[Display Success]
    I --> E

    style C fill:#ef5350
    style H fill:#ef5350
    style I fill:#66bb6a
```

**Error Handling Principles:**

1. **Early Validation**: Check inputs before processing
2. **Error Wrapping**: Add context to errors as they bubble up
3. **User-Friendly Messages**: Translate technical errors to clear messages
4. **Graceful Degradation**: Continue operation where possible
5. **Logging**: Record errors for debugging (future enhancement)

---

## Scalability Considerations

### Current Scale

- **Target**: Individual users managing 100-10,000 tasks
- **Performance**: In-memory operations on entire task list
- **Storage**: Single JSON/YAML file

### Future Scaling Paths

```mermaid
graph LR
    A[Current<br/>File-based] --> B[SQLite<br/>Local DB]
    B --> C[PostgreSQL<br/>Shared DB]
    C --> D[Cloud Storage<br/>Multi-device]

    style A fill:#4caf50
    style B fill:#2196f3
    style C fill:#ff9800
    style D fill:#9c27b0
```

**Scaling Strategy:**

1. Add SQLite storage implementation (10K+ tasks)
2. Implement pagination for large lists
3. Add caching layer
4. Consider client-server architecture for collaboration

---

## Testing Strategy

```mermaid
graph TB
    subgraph "Unit Tests"
        UT1[Models Tests<br/>Task logic]
        UT2[Filter Tests<br/>Filter algorithms]
        UT3[Sort Tests<br/>Sort algorithms]
        UT4[Storage Tests<br/>Mocked I/O]
    end

    subgraph "Integration Tests"
        IT1[Command Tests<br/>End-to-end flows]
        IT2[Storage Tests<br/>Real file I/O]
    end

    subgraph "Manual Tests"
        MT1[Interactive Mode<br/>User experience]
        MT2[CLI Usage<br/>Real scenarios]
    end

    style UT1 fill:#81c784
    style UT2 fill:#81c784
    style UT3 fill:#81c784
    style UT4 fill:#81c784
    style IT1 fill:#64b5f6
    style IT2 fill:#64b5f6
    style MT1 fill:#ffb74d
    style MT2 fill:#ffb74d
```

---

## Security Considerations

### Current Security Measures

```mermaid
graph TD
    A[User Input] --> B[Input Validation]
    B --> C[Sanitization]
    C --> D[Processing]

    E[File Operations] --> F[Permission Checks]
    F --> G[Safe File Paths]
    G --> H[Write to Disk]

    I[Data Storage] --> J[Local Storage Only]
    J --> K[User Directory Permissions]

    style B fill:#66bb6a
    style C fill:#66bb6a
    style F fill:#66bb6a
    style G fill:#66bb6a
    style K fill:#66bb6a
```

**Security Principles:**

1. **Input Validation**: Sanitize all user inputs
2. **File Permissions**: Use appropriate file permissions (0644 for files, 0755 for directories)
3. **Path Safety**: Prevent path traversal attacks
4. **No Network**: No external communication (air-gapped by design)
5. **Local Only**: Data stays on user's machine

### Future Security Enhancements

- [ ] Encryption at rest for sensitive tasks
- [ ] Optional password protection
- [ ] Secure cloud sync with E2E encryption

---

## Performance Characteristics

### Time Complexity

| Operation    | Best Case  | Average Case | Worst Case |
| ------------ | ---------- | ------------ | ---------- |
| Add Task     | O(1)       | O(1)         | O(n)\*     |
| List Tasks   | O(n)       | O(n)         | O(n)       |
| Filter Tasks | O(n)       | O(n)         | O(n)       |
| Sort Tasks   | O(n log n) | O(n log n)   | O(n log n) |
| Search Tasks | O(n)       | O(n)         | O(n)       |
| Edit Task    | O(n)       | O(n)         | O(n)       |
| Delete Task  | O(n)       | O(n)         | O(n)       |
| Save to Disk | O(n)       | O(n)         | O(n)       |

\*O(n) when saving to disk

### Space Complexity

- **Memory**: O(n) where n = number of tasks
- **Disk**: O(n) for main file + O(m) for m backups

### Performance Optimizations

- **In-memory operations**: All filtering/sorting done in RAM
- **Lazy loading**: Load data only when needed
- **Efficient sorting**: Uses Go's optimized sort.Slice
- **Minimal allocations**: Reuse structures where possible

---

## Extension Points

The architecture provides clear extension points for future enhancements:

```mermaid
graph TB
    subgraph "Current Features"
        A[Task Management]
        B[Filtering]
        C[Sorting]
        D[Storage]
    end

    subgraph "Easy Extensions"
        E[New Storage Backend<br/>Implement Storage interface]
        F[New Filter Type<br/>Add to Filter struct]
        G[New Sort Criteria<br/>Add to SortBy enum]
        H[New Command<br/>Add cmd/*.go file]
    end

    subgraph "Future Features"
        I[Recurring Tasks<br/>Add RecurrenceRule to Task]
        J[Subtasks<br/>Add ParentID to Task]
        K[Collaboration<br/>Add User to Task]
        L[Sync Service<br/>New sync package]
    end

    A --> E
    B --> F
    C --> G
    D --> H

    E --> I
    F --> J
    G --> K
    H --> L

    style E fill:#81c784
    style F fill:#81c784
    style G fill:#81c784
    style H fill:#81c784
    style I fill:#64b5f6
    style J fill:#64b5f6
    style K fill:#64b5f6
    style L fill:#64b5f6
```

---

## Conclusion

The Todo CLI architecture is designed to be:

✅ **Maintainable**: Clear separation of concerns, well-documented
✅ **Extensible**: Easy to add new features and storage backends
✅ **Testable**: Interface-based design enables easy mocking
✅ **Performant**: Efficient algorithms and data structures
✅ **User-Friendly**: Intuitive CLI and interactive mode
✅ **Professional**: Follows Go best practices and industry standards

The architecture balances simplicity for current needs with flexibility for future growth, making it a solid foundation for a production-grade CLI application.
