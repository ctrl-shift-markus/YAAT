<div align="center">
  <img src="../assets/Logo.png" alt="The acronym YAAT hand drawn with a grassy texture" width="250">
  <br>
  <a href="https://github.com/ctrl-shift-markus/YAAT/releases">
    <img src="https://img.shields.io/github/v/release/ctrl-shift-markus/YAAT?color=41D800&amp;style=flat" alt="Version">
  </a>
  <img src="https://img.shields.io/badge/language-Go-00ADD8?style=flat" alt="Language">
  <img src="https://img.shields.io/badge/platforms-Windows%2C_macOS%2C_Linux-9700D8?style=flat" alt="Platforms">
  <a href="https://github.com/ctrl-shift-markus/YAAT/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/ctrl-shift-markus/YAAT?color=D82B00&amp;style=flat" alt="License">
  </a>
  <p>Contributing Guidelines</p>
</div>

## Pull Requests

PRs are always welcome!

### Code

Before you submit any code, make sure it: 
- Builds without errors (`go build ./cmd/yaat`)
- Passes all tests (`go test ./...`)
    - If you are implementing a core feature, make sure you add tests for it as well.
- Passes the linter (`golangci-lint`, which you can install [here](https://golangci-lint.run/docs/welcome/install/local/))

If it won't pass these requirements locally, when you try to open a PR the status checks will block it.

### How to Submit

To submit a PR:

1. Create a new issue (if required) describing your change and label it, making sure it's not a duplicate of an existing one
2. Fork YAAT, create a new branch (following the [conventional branch](https://conventionalbranch.org/) specification) and add your code
3. Open a PR that closes your issue (by adding `closes #<issue-id>` in the description) and apply the same label
4. Fix any issues flagged by status checks
5. Wait for your PR to be reviewed and merged

## Issues

If you've found a bug or have an idea for a new feature but don't want to contribute yourself, you can simply open a [new issue](https://github.com/ctrl-shift-markus/yaat/issues).

## License

By contributing to YAAT, you agree that your contribution(s) will be licensed under the [GNU General Public License v3.0](LICENSE).