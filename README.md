# gocat

> Generate AI-ready project context from any codebase.

`gocat` is a fast, lightweight CLI that scans your project, respects ignore files, lets you interactively select files and directories, and generates a clean Markdown document optimized for AI coding assistants such as ChatGPT, Claude, Gemini, Cursor, GitHub Copilot, Continue, Cline, Roo Code, and more.

---

## Why?

Large projects contain thousands of files.

Instead of copying files manually, `gocat` lets you:

- browse your project
- select exactly the files you want
- generate a single AI-friendly Markdown document

The result can be pasted directly into your favorite AI assistant.

---

## Features

- 📁 Interactive project tree
- 📄 Interactive file & directory selection
- 🚀 Fast recursive scanning
- 🙈 Respects ignore rules
    - `.gitignore`
    - `.dockerignore`
    - `.gocatignore`
    - `.aiignore`
- 🐳 Automatically includes important project files
- 📝 Markdown output optimized for LLMs
- ⚡ Lightweight single binary
- 🌍 Language agnostic

---

## Installation

### Go

```bash
go install github.com/rahmatwaisi/gocat/cmd/gocat@latest
```

---

## Usage

Run inside any project.

```bash
gocat
```

The tool scans your project and displays an interactive tree.

Example:

```text
[d1] 📁 cmd/
└── [d2] 📁 gocat/
    └── [1] 📄 main.go

[d3] 📁 internal/
├── [d4] 📁 formatter/
│   ├── [2] 📄 formatter.go
│   └── [3] 📄 markdown.go

[4] 📄 README.md
```

Select files or directories.

```text
1 d3 4
```

A Markdown document is generated.

---

## Generated Output

````markdown
# cmd/gocat/main.go

```go
package main

func main() {
    ...
}
```

---

# README.md

```md
# gocat
...
```
