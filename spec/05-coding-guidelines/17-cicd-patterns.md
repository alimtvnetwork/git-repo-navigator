# CI/CD Patterns

## Pipeline Structure

### Stage Ordering

Every pipeline follows a fixed stage sequence:

| Stage | Purpose | Fails Fast |
|-------|---------|------------|
| Lint | Static analysis, formatting | Yes |
| Test | Unit + integration tests | Yes |
| Build | Compile artifacts | Yes |
| Package | Archive, compress, checksum | Yes |
| Deploy | Push to target environment | Yes |
| Verify | Smoke tests post-deploy | Yes |

### Rules

- Each stage must pass before the next begins
- Lint and test stages run in parallel when independent
- Build artifacts are produced exactly once and promoted across environments
- Never rebuild between staging and production — promote the same artifact

### Pipeline-as-Code

- Define pipelines in version-controlled YAML (e.g., `.github/workflows/*.yml`)
- No manual configuration in CI provider UIs
- Pin action versions to full SHA or exact tag — never use `@latest` or `@main`

```yaml
# Good
- uses: actions/checkout@v4
- uses: actions/setup-go@v5

# Bad
- uses: actions/checkout@main
```

## Build Caching

### Cache Layers

| Layer | What to Cache | Key Strategy |
|-------|---------------|--------------|
| Dependencies | `node_modules`, Go module cache | Hash of lock file |
| Build output | Compiled binaries, intermediate objects | Hash of source + config |
| Docker layers | Base images, dependency layers | Content-addressable |
| Test results | Coverage reports | Branch + commit SHA |

### Cache Invalidation Rules

- Always include the lock file hash in the cache key
- Prefix keys with runner OS and tool version
- Set `restore-keys` for partial cache hits
- Never cache secrets, credentials, or environment-specific config

```yaml
- uses: actions/cache@v4
  with:
    path: ~/go/pkg/mod
    key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
    restore-keys: ${{ runner.os }}-go-
```

### Dependency Resolution

- Use `go mod download` or `npm ci` (not `npm install`) for reproducible installs
- Frozen lock files in CI — fail if lock file is out of date
- Vendor dependencies for hermetic builds when network reliability is a concern

## Deployment Gates

### Environment Promotion

```
dev → staging → production
```

| Gate | Requirement |
|------|-------------|
| dev → staging | All tests pass, lint clean, build succeeds |
| staging → prod | Smoke tests pass, manual approval (for critical services) |
| hotfix → prod | Abbreviated pipeline, post-deploy verification mandatory |

### Approval Rules

- Production deployments require at least one explicit approval for critical services
- Automated deployments are acceptable for non-critical services with full test coverage
- Approval requests must include: diff summary, test results, rollback plan
- Time-boxed approvals — auto-expire after 24 hours

### Feature Flags

- Use feature flags to decouple deployment from release
- Deploy dark features behind flags; enable progressively
- Clean up stale flags within one sprint of full rollout

## Rollback Strategies

### Rollback Decision Tree

```
Deploy fails?
├── Pre-traffic: cancel deployment, no rollback needed
├── Post-traffic, < 5 min: automatic rollback to previous version
└── Post-traffic, > 5 min: manual assessment, then rollback or forward-fix
```

### Rollback Mechanisms

| Strategy | When to Use | Recovery Time |
|----------|-------------|---------------|
| Revert commit | Simple code changes | Minutes |
| Re-deploy previous artifact | Binary/container deployments | Minutes |
| Blue-green switch | Zero-downtime services | Seconds |
| Database rollback | Schema changes with backward compat | Varies |

### Rules

- Every deployment must have a documented rollback procedure
- Keep the previous version's artifact available for at least 7 days
- Database migrations must be backward-compatible (expand-contract pattern)
- Rollback must not require a full rebuild — use pre-built artifacts
- Test rollback procedures periodically in staging

### Backward-Compatible Migrations

```
Phase 1 (expand):  Add new column, keep old column
Phase 2 (migrate): Backfill data, deploy code using new column
Phase 3 (contract): Remove old column after verification
```

## Artifact Management

### Naming Convention

```
<binary>-<os>-<arch>.<ext>
```

Examples: `gitmap-linux-amd64.tar.gz`, `gitmap-windows-arm64.zip`

### Checksums

- Generate SHA-256 checksums for all release artifacts
- Publish checksums alongside artifacts
- Verify checksums before deployment

### Retention

| Artifact Type | Retention |
|---------------|-----------|
| Release binaries | Permanent |
| CI build artifacts | 30 days |
| Test reports | 90 days |
| Coverage data | 90 days |

## Constraints

- No secrets in pipeline definitions — use CI provider secret storage
- No interactive prompts in CI scripts
- All scripts must exit non-zero on failure (`set -euo pipefail` in Bash)
- Pipeline duration target: < 10 minutes for the full cycle
- Flaky tests must be quarantined immediately, not retried silently

## Contributors

- [**Md. Alim Ul Karim**](https://www.linkedin.com/in/alimkarim) — Creator & Lead Architect. System architect with 20+ years of professional software engineering experience across enterprise, fintech, and distributed systems. Recognized as one of the top software architects globally. Alim's architectural philosophy — consistency over cleverness, convention over configuration — is the driving force behind every design decision in this framework.
  - [Google Profile](https://www.google.com/search?q=Alim+Ul+Karim)
- [Riseup Asia LLC (Top Leading Software Company in WY)](https://riseup-asia.com) (2026)
  - [Facebook](https://www.facebook.com/riseupasia.talent/)
  - [LinkedIn](https://www.linkedin.com/company/105304484/)
  - [YouTube](https://www.youtube.com/@riseup-asia)
