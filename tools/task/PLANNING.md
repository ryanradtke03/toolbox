# td — Planning

## What it is

A personal CLI task manager. Fast to type, no ceremony. Published to npm as a global install.

---

## Commands & Flags

### `td add <title>`

Create a new task.

```bash
td add "fix auth bug"
td add "fix auth bug" -p high -t backend -t work -g finapse
td add "buy groceries" -p l -t errands
```

| Flag         | Short | Values                   | Default |
| ------------ | ----- | ------------------------ | ------- |
| `--priority` | `-p`  | `h/high` `m/med` `l/low` | `m`     |
| `--tag`      | `-t`  | any string, repeatable   | —       |
| `--project`  | `-g`  | any string               | —       |

---

### `td list`

Show all open tasks. Default sort: newest first.

```bash
td list
td list -p high
td list -t work
td list -g finapse
td list --sort oldest
td list --sort alpha
td list --all
td list -t work -p h --sort alpha
```

| Flag         | Short | Values                                       | Default  |
| ------------ | ----- | -------------------------------------------- | -------- |
| `--priority` | `-p`  | `h/high` `m/med` `l/low`                     | —        |
| `--tag`      | `-t`  | any string, repeatable                       | —        |
| `--project`  | `-g`  | any string                                   | —        |
| `--sort`     | `-s`  | `n/newest` `o/oldest` `a/alpha` `p/priority` | `newest` |
| `--all`      | `-a`  | boolean                                      | false    |

---

### `td done <id>`

Mark a task complete.

```bash
td done a3f9c2b1
```

No flags.

---

### `td rm <id>`

Delete a task permanently.

```bash
td rm a3f9c2b1
td rm a3f9c2b1 --force
```

| Flag      | Short | Values  | Default |
| --------- | ----- | ------- | ------- |
| `--force` | —     | boolean | false   |

Prompts "are you sure?" unless `--force` is passed.

---

### `td edit <id>`

Update a task's metadata inline, or open in `$EDITOR` with no flags.

```bash
td edit a3f9c2b1              # opens $EDITOR
td edit a3f9c2b1 -p low
td edit a3f9c2b1 -t backend
td edit a3f9c2b1 -g finapse
```

| Flag         | Short | Values                   | Default |
| ------------ | ----- | ------------------------ | ------- |
| `--priority` | `-p`  | `h/high` `m/med` `l/low` | —       |
| `--tag`      | `-t`  | any string, repeatable   | —       |
| `--project`  | `-g`  | any string               | —       |

---

## Data Model

```ts
type Priority = "low" | "med" | "high";
type Status = "open" | "done";

type Task = {
  id: string; // nanoid(8) e.g. 'a3f9c2b1'
  title: string;
  status: Status;
  priority: Priority; // default: 'med'
  tags: string[];
  project?: string;
  createdAt: string; // ISO 8601
  updatedAt: string; // ISO 8601
};
```

---

## Flag Value Normalization

Both short and long forms accepted everywhere:

```ts
// Priority
'h' | 'high' → 'high'
'm' | 'med'  → 'med'
'l' | 'low'  → 'low'

// Sort
'n' | 'newest'   → newest first (default)
'o' | 'oldest'   → oldest first
'a' | 'alpha'    → A → Z by title
'p' | 'priority' → high → med → low
```

---

## Storage

```
~/.local/share/td/tasks.json
```

Plain JSON array of Task objects. Migrate to SQLite if filtering gets slow (unlikely for personal use).

---

## Project Structure

```
td/
├── cmd/
│   └── td/
│       └── main.go           # entrypoint, cobra root + subcommands
├── internal/
│   ├── task/
│   │   └── task.go           # Task struct, parsePriority(), parseSort()
│   ├── db/
│   │   └── db.go             # ReadTasks() / WriteTasks()
│   └── commands/
│       ├── add.go
│       ├── list.go
│       ├── done.go
│       ├── rm.go
│       └── edit.go
├── go.mod
├── go.sum
├── PLANNING.md
└── README.md
```

---

## Dependencies

| Package     | Purpose                       |
| ----------- | ----------------------------- |
| `commander` | arg parsing, flag definitions |
| `nanoid`    | ID generation (`nanoid(8)`)   |
| `chalk`     | colored terminal output       |
| `tsx`       | dev runner (no build step)    |

---

## MVP Checklist

- [ ] `types.ts` — Task type
- [ ] `db.ts` — readTasks / writeTasks
- [ ] `parse.ts` — parsePriority, parseSort
- [ ] `add` command
- [ ] `list` command
- [ ] `done` command
- [ ] `rm` command
- [ ] `edit` command
- [ ] Colored `list` output
- [ ] Human-readable errors
- [ ] `README.md` with install + usage
- [ ] `npm publish`

---

## Later (not MVP)

- Due dates (`--due tomorrow`, `--due 2026-06-10`)
- `td clear` — archive all done tasks
- Export to markdown (`td export`)
- SQLite migration
- Shell completions (zsh, bash, fish)
- Recurring tasks
- `td log` — activity history
