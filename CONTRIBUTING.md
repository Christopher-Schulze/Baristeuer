# Contributing to Baristeuer

Thank you for considering a contribution!

## Branching Strategy
- Create a new branch for each change based on `main`.
- Use descriptive names such as `feature/<topic>` or `bugfix/<issue>`.
- Open a pull request targeting `main` when your work is ready for review.

## Coding Conventions
- Format Go code with `gofmt -w` before committing. The CI will fail if formatting changes are detected.
- Keep JavaScript and React files formatted using Prettier.
- Ensure the project builds without errors using `go build ./...` and `npm run build` for the UI.
- Write tests for new functionality when possible.

## Release Process
- Create an annotated tag with the new version (e.g. `git tag -a v1.0.0 -m "v1.0.0"`).
- Push the tag to GitHub.
- The [release workflow](.github/workflows/release.yml) runs `scripts/package.sh` which builds binaries for macOS, Windows and Linux.
- Generated ZIP files and installers are automatically attached to the GitHub release.

