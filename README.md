# Go-Learning

A personal learning repository for mastering the Go programming language, built with the long-term goal of becoming a **Software Engineer (SWE)** and **DevOps** professional.

This is not a tutorial site or a finished application. It is a structured workspace where concepts are learned, practiced, and reinforced through hands-on code — chapter by chapter.

---

## Purpose

Go is one of the most practical languages to learn for both software engineering and DevOps careers. It is used to build:

- **Backend services and APIs** (SWE)
- **CLI tools** used daily in development workflows
- **Cloud-native infrastructure** — Kubernetes, Docker, Terraform, and many observability tools are written in or heavily integrated with Go
- **Reliable, concurrent systems** that handle production workloads at scale

This repository tracks that journey from the ground up: understanding syntax and fundamentals first, then applying them to real problem-solving exercises that mirror how Go is used in industry.

---

## Career Goals

### Software Engineering (SWE)

Go teaches strong fundamentals that transfer across languages and roles:

- Explicit typing and clear program structure
- Error handling as a first-class concern
- Building small, focused programs that grow into services
- Reading and writing code that is easy to maintain in teams

The `Basics/` folder builds the foundation. The `AI_Challenges/` folder forces you to combine concepts the way you would on a real task — not just copy examples, but solve a problem.

### DevOps

DevOps engineers work at the intersection of development and operations. Go fits that path directly:

- Many DevOps tools (`kubectl`, `helm`, `terraform` providers, `prometheus`, `consul`) are Go-based or Go-friendly
- CLI programs with `os.Args`, flags, and structured output are core DevOps patterns
- Go's fast compile times and single-binary deployments make it ideal for automation, agents, and internal tooling
- Understanding Go makes it easier to read, debug, and contribute to the infrastructure stack

Command-line exercises in this repo (e.g. parsing `os.Args`, handling errors, formatted output) are early practice for the kind of tooling DevOps engineers build and operate every day.

---

## Repository Structure

```
Go-Learning/
├── Basics/
│   ├── Chapter_Notes/              # Notes and examples per topic
│   └── Chapter_Practice_Projects/  # Small exercises to reinforce lessons
├── AI_Challenges/                  # Real-world problem-solving projects
└── README.md
```

### `Basics/`

The core curriculum. Each subfolder under `Chapter_Notes/` maps to a specific Go concept learned from coursework or self-study.

| Folder | Topic |
|--------|-------|
| `Hello World!/` | Entry point — `package main`, `import`, `fmt.Println` |
| `Curly_Braces/` | Go syntax rules and program structure |
| `Variables_Understaind/` | Variable declaration (`var`, `:=`), types, zero values |
| `If_and_Else_Statements/` | Conditionals, `switch`, `os.Args`, `strconv` |
| `Loops_and_range/` | `for` loops, `break`, slices, `range` |

`Chapter_Practice_Projects/` holds standalone exercises that drill variables, arithmetic, and formatting without tying them to a single chapter file.

**Why separate Notes from Practice?**
- **Chapter Notes** = learn and document a concept in isolation
- **Practice Projects** = apply it without step-by-step guidance

---

## What I've Learned

A running log of concepts covered in this repo and where they live in the code.

| Topic | What I learned | File(s) |
|-------|----------------|---------|
| Program structure | `package main`, `import`, `func main()`, `fmt.Println` | `Basics/Chapter_Notes/Hello World!/hw.go` |
| Syntax rules | Go requires specific brace placement; strict formatting | `Basics/Chapter_Notes/Curly_Braces/curly.go` |
| Variables | `var`, `:=`, zero values, global vs local scope | `Basics/Chapter_Notes/Variables_Understaind/var.go` |
| Arithmetic & formatting | Math operations, `fmt.Printf` with `%d` and `%.2f`, `math.Abs` | `Basics/Chapter_Practice_Projects/variables.go`, `variables_2.go` |
| Conditionals | `if` / `else if` / `else`, early `return` | `Basics/Chapter_Notes/If_and_Else_Statements/control.go` |
| Switch statements | `switch` with an expression and without one (boolean cases) | `Basics/Chapter_Notes/If_and_Else_Statements/control.go` |
| CLI input | Reading command-line args with `os.Args`, validating `len(os.Args)` | `Basics/Chapter_Notes/If_and_Else_Statements/control.go` |
| Parsing strings | Converting user input with `strconv.Atoi` and error handling | `Basics/Chapter_Notes/If_and_Else_Statements/control.go` |
| Loops | Traditional `for`, while-style loops, `break` | `Basics/Chapter_Notes/Loops_and_range/forloops.go` |
| Slices & range | Creating slices, iterating with `for i, v := range` | `Basics/Chapter_Notes/Loops_and_range/forloops.go` |

### Topics checklist

- [x] Hello World and program structure
- [x] Variables (`var`, `:=`, globals, types)
- [x] Conditionals (`if` / `else if` / `else`)
- [x] `switch` statements
- [x] Command-line arguments (`os.Args`)
- [x] String-to-number conversion (`strconv`)
- [x] `for` loops and `range` over slices
- [x] Formatted output (`fmt.Printf`, `fmt.Print`)

### Upcoming (planned path toward SWE / DevOps)

- [ ] Functions and packages
- [ ] Structs and methods
- [ ] Error handling patterns (`error` type, wrapping)
- [ ] File I/O and environment variables
- [ ] Go modules and project layout
- [ ] Testing (`go test`)
- [ ] HTTP servers and REST APIs
- [ ] Concurrency (goroutines, channels)
- [ ] Building and shipping CLI tools as binaries
- [ ] Dockerizing a Go service
- [ ] CI/CD basics with Go projects

---

## How I Completed It

This repo follows a three-step learning loop. Each step builds on the last — notes teach the concept, practice drills it, and AI challenges prove you can combine everything to solve a real problem.

### Step 1: Learn (`Basics/Chapter_Notes/`)

Read and write small example files for one topic at a time. Each file focuses on a single idea — e.g. how `switch` works, or how to loop over a slice.

### Step 2: Practice (`Basics/Chapter_Practice_Projects/`)

Apply the topic without copying from notes. Practice projects use variables, math, and formatting in short standalone programs.

### Step 3: Apply (`AI_Challenges/`)

Solve a scenario-based problem that combines multiple topics from Steps 1 and 2. This is where lesson knowledge becomes problem-solving skill.

### Completion log

| Challenge | Status | How it was completed |
|-----------|--------|----------------------|
| **Challenge 01: Snack Shop Checkout** | Completed | Learned variables and conditionals in `Basics/`, then loops and slices in `Loops_and_range/`. Built a CLI checkout tool that parses args, calculates totals with tiered discounts, and handles payment outcomes. |

**Challenge 01 — what I built:**
- A command-line snack shop that accepts order quantities and cash via `os.Args`
- Menu stored in slices, receipt built with `range` loops
- Tiered discounts using `switch` (0% / 5% / 10% based on item count)
- Payment logic for exact payment, change, or amount still owed
- Input validation and error handling for bad args and negative quantities

---

## AI Challenges

### What they are

`AI_Challenges/` holds **scenario-based problem-solving projects** tailored to your current skill level. They are not copy-paste exercises — each one presents a small real-world situation and asks you to design and build a solution using only what you've learned so far.

### Why they exist — real-world problem solving

Chapter notes teach **syntax**. AI challenges teach **thinking like a developer**:

- Break a problem into steps (input → process → output)
- Combine multiple concepts in one program (not one concept per file)
- Handle bad input the way production tools do (validate, error, return)
- Build CLI tools — the same pattern used in DevOps scripts, internal utilities, and automation

Every challenge is modeled after situations you'd encounter as an SWE or DevOps engineer: parsing user input, applying business rules, formatting output, and shipping a working program from the terminal.

### How to work through a challenge

1. **Open the challenge file** — the full problem, rules, examples, and hints are in the comment block at the top of the `.go` file.
2. **Read the scenario** — understand what the program should do before writing code.
3. **Plan first** — list the steps: what to parse, what to calculate, what to print.
4. **Use your notes** — refer back to `Basics/Chapter_Notes/` when you forget syntax.
5. **Try it yourself first** — attempt the solution before asking for help.
6. **Run and test** — use the example commands in the challenge header; try edge cases (wrong args, not enough cash, etc.).
7. **Move on** — when a challenge is done, new ones are added as you learn more topics.

### How to run a challenge

```bash
cd AI_Challenges/Challenge_01_Snack_Shop
go run snack_shop.go <chips> <sodas> <candies> <cash>
```

**Example:**
```bash
go run snack_shop.go 2 3 1 15.00
```

### Challenge tracker

| # | Challenge | Status | Concepts | Real-world parallel |
|---|-----------|--------|----------|---------------------|
| 01 | Snack Shop Checkout | Completed | slices, `range`, `switch`, `if/else`, `os.Args`, `strconv`, `fmt.Printf` | Point-of-sale / billing CLI — pricing rules, discounts, payment handling |
| 02 | *(coming soon)* | — | functions, structs | TBD as new topics are covered |

New challenges are added as more topics are learned. Each one scales in difficulty and maps to skills used in SWE and DevOps work.

---

## How to Run Code

From any `.go` file directory:

```bash
go run <filename>.go
```

For programs that take command-line arguments:

```bash
go run control.go 10
go run snack_shop.go 2 3 1 15.00
```

**Requirements:** [Go](https://go.dev/dl/) installed (`go version` to verify).

---

## Learning Workflow (summary)

```
Chapter Notes  →  Practice Projects  →  AI Challenges
    (learn)          (drill)              (solve)
```

1. **Learn** — work through a topic in `Basics/Chapter_Notes/`
2. **Practice** — complete exercises in `Basics/Chapter_Practice_Projects/`
3. **Apply** — solve the matching challenge in `AI_Challenges/`
4. **Repeat** — move to the next topic; challenges grow with your skills

---

## Why Go for This Path

| Area | Why Go matters |
|------|----------------|
| SWE | Clean syntax, strong stdlib, widely used for backends and microservices |
| DevOps | Native to the cloud-native ecosystem; ideal for CLIs, agents, and automation |
| Both | Fast to compile, easy to deploy as a single binary, great for production systems |

---

*This repository grows with the learning process. Every file here is a step toward writing reliable software and operating it in production.*
