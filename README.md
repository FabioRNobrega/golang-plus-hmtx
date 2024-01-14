# Go Messenger Example | Farno

This is a simple webserver application written in Go with HTMX, inspired by the talk by [BugBytes](https://www.youtube.com/@bugbytes3923) in the video on [YouTube](https://www.youtube.com/watch?v=F9H6vYelYyU). I've incorporated some enhancements, including the layout, while retaining the core concepts presented in the video.

I use:

- Go version 1.21.6

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Tests](#tests)
- [Git Guideline](#git-guideline)

## Install

To run this application, you need to have Go installed. You can follow the [Go Installation Guide](https://golang.org/doc/install) for assistance. For a robust experience, I strongly recommend using ASDF for managing Go versions. To set up this application, ensure you have ASDF installed and configured. Follow the [ASDF Installation Guide](https://asdf-vm.com/#/core-manage-asdf) for detailed instructions.


## Usage

Clone this repository and navigate into it. Install dependencies with the following command:

```bash
go get -u ./...
```

Start the Go application by running:

```bash
go run ./main.go
```

Now you can visit localhost:3000 from your browser.

## Tests

Tests are written using the Go testing framework and can be found in the tests folder. To run the tests, use:

```bash
go test ./...
```

## Git Guideline

Create your branches and commits using the English language and following this guideline.

### Branches

- Feature:  `feat/branch-name`
- Hotfix: `hotfix/branch-name`
- POC: `poc/branch-name`
- Example `example/branch-name`

### Commits prefix

- Chore: `chore(context): message`
- Feat: `feat(context): message`
- Fix: `fix(context): message`
- Refactor: `refactor(context): message`
- Tests: `tests(context): message`
- Docs: `docs(context): message`
