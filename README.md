# toolbox

A personal collection of CLI tools built with TypeScript.

# Tools

| Tool     | Description                                            |
| -------- | ------------------------------------------------------ |
| `devlog` | Log and review developer notes and progress entries    |
| `focus`  | Manage focus sessions and track what you're working on |
| `habit`  | Track daily habits and streaks                         |
| `jot`    | Quickly capture thoughts and short notes               |

## Development

**Prerequisites:** Node.js, npm

**Install dependencies:**

```bash
npm install
```

**Run a tool without building:**

```bash
npm run dev:<tool>
```

**Build:**

```bash
npm run build
```

**Install globally:**

```bash
npm install -g .
```

## Stack

- [TypeScript](https://www.typescriptlang.org/)
- [Commander.js](https://github.com/tj/commander.js) — CLI argument parsing
- [Chalk](https://github.com/chalk/chalk) — terminal styling
