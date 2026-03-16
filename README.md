# gh-codeowners

A [GitHub CLI](https://cli.github.com/) extension for GitHub's [CODEOWNERS file](https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/about-code-owners#codeowners-syntax).

The `codeowners` [GitHub CLI](https://cli.github.com/) extension identifies the owners for files in a local repository or directory. This extension uses the [codeowners go module from hmarr](https://github.com/hmarr/codeowners).

## Installation
```bash
gh extension install robandpdx/gh-codeowners
```

### Usage

By default, the command line tool will walk the directory tree, printing the code owners of any files that are found.

```console
$ gh codeowners --help
usage: codeowners <path>...
  -f, --file string           CODEOWNERS file path
  -h, --help                  show this help message
  -i, --ignore-file strings   gitignore-style file of patterns to exclude (repeatable)
  -o, --owner strings         filter results by owner
  -u, --unowned               only show unowned files (can be combined with -o)

$ ls
CODEOWNERS       DOCUMENTATION.md README.md        example.go       example_test.go

$ cat CODEOWNERS
*.go       @example/go-engineers
*.md       @example/docs-writers
README.md  product-manager@example.com

$ gh codeowners
CODEOWNERS                           (unowned)
README.md                            product-manager@example.com
example_test.go                      @example/go-engineers
example.go                           @example/go-engineers
DOCUMENTATION.md                     @example/docs-writers
```

To limit the files the tool looks at, provide one or more paths as arguments.

```console
$ gh codeowners *.md
README.md                            product-manager@example.com
DOCUMENTATION.md                     @example/docs-writers
```

Pass the `--owner` flag to filter results by a specific owner.

```console
$ gh codeowners -o @example/go-engineers
example_test.go                      @example/go-engineers
example.go                           @example/go-engineers
```

Pass the `--unowned` flag to only show unowned files.

```console
$ gh codeowners -u
CODEOWNERS                           (unowned)
```

Pass the `--ignore-file` flag to exclude files matching patterns in a [gitignore-style](https://git-scm.com/docs/gitignore) file. The flag can be repeated to load multiple ignore files.

```console
$ cat .buildignore
*.go

$ gh codeowners -i .buildignore
CODEOWNERS                           (unowned)
README.md                            product-manager@example.com
DOCUMENTATION.md                     @example/docs-writers
```
