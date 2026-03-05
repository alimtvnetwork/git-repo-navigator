# Data Model

## ScanRecord

| Field            | Type   | Required | Default | Notes                          |
|------------------|--------|----------|---------|--------------------------------|
| id               | string | yes      | UUID    | Unique record identifier       |
| repoName         | string | yes      | —       | Derived from remote URL         |
| httpsUrl         | string | yes      | —       | HTTPS clone URL                 |
| sshUrl           | string | yes      | —       | SSH clone URL                   |
| branch           | string | yes      | "main"  | Current checked-out branch      |
| relativePath     | string | yes      | —       | Path relative to scan root      |
| absolutePath     | string | yes      | —       | Full filesystem path            |
| cloneInstruction | string | yes      | —       | Full `git clone` command        |
| notes            | string | no       | ""      | User or system notes            |

## Config

See [06-config.md](./06-config.md).
