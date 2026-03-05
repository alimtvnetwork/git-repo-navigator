# Project Repository

This repository contains two main components:

## 1. GitMap CLI (Go)

A command-line tool that scans directory trees for Git repositories, extracts clone URLs and branch info, and outputs structured data (terminal, CSV, JSON). It can re-clone repositories from that data, preserving folder hierarchy.

**→ [GitMap Documentation](./gitmap/README.md)**  
**→ [Specifications](./spec/01-app/)**

### Quick Start

```powershell
cd gitmap
.\build.ps1           # Pull, build, deploy to E:\bin-run
```

Or manually:

```bash
cd gitmap
go build -o ../bin/gitmap.exe .
```

### Usage

```bash
gitmap scan ./projects --mode ssh --output csv
gitmap clone ./gitmap-output/gitmap.csv --target-dir ./restored
```

---

## 2. Web Frontend (React + Vite)

A React + TypeScript + Tailwind CSS web application scaffold.

### Setup

```sh
npm i
npm run dev
```

### Tech Stack

- Vite
- TypeScript
- React
- shadcn-ui
- Tailwind CSS
