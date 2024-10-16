# create-commit-status-action

## Dev

Releases can be built with:

```bash
make release
```

Note: windows-arm64 and darwin-arm64 are not compressed due to non-support and
compatibility issues respectively.

Binaries are in the `out` dir

The action release branch is `main` and is automated upon pushes to
`release`.
