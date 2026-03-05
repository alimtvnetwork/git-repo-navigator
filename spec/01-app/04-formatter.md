# Formatter

## Responsibility

Render a list of `ScanRecord` into one of three formats.

## Formats

### Terminal

Print a table or list of clone instructions to stdout.  
Each line: `git clone -b <branch> <url> <relative-path>`

### CSV

Write a CSV file with headers:

```
repoName,httpsUrl,sshUrl,branch,relativePath,absolutePath,cloneInstruction,notes
```

### JSON

Write a JSON array of `ScanRecord` objects.

## Output Location

- Terminal: stdout.
- CSV/JSON: path from `--out-file` flag, or `config.outputDir` + default name.

## Multiple Outputs

The `--output` flag may accept comma-separated values in a future version.  
Phase 1: single output format per invocation.
