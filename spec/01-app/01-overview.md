# GitMap — Overview

## Purpose

GitMap is a CLI tool that scans a directory tree for Git repositories,
extracts clone URLs and branch information, and outputs structured data
(terminal, CSV, JSON). It can also re-clone repositories from that
structured data, preserving the original folder hierarchy.

## Working Name

`gitmap`

## Code Style Constraints

| Constraint          | Rule                              |
|---------------------|-----------------------------------|
| `if` conditions     | Always positive — no `!`, no `!=` |
| Function length     | 8–15 lines                        |
| File length         | 100–200 lines max                 |
| Package granularity | One responsibility per package    |

## High-Level Components

1. **Config loader** — reads JSON config, merges with CLI flags.
2. **Scanner** — walks directories, detects `.git` folders.
3. **Mapper** — converts raw Git data into output records.
4. **Formatter** — renders records to terminal, CSV, or JSON.
5. **Cloner** — re-clones repos from a previously generated file.

## Assumptions

- Remote URL is extracted from `origin` remote only.
- Symlinked directories are not followed.
- "Text file" input means one clone command per line.
