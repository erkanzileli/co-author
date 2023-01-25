# co-author

A TUI that saves time from writing Co-authored-by lines to the commits in a fancy way. Works as Git hook.

## Motivation

When doing pair programming, it's a good practice that add your pair's contact information to the commit message. You do
that by the following format

```
Co-authored-by: Your Pair <yourpair@mail.com>
```

So your pair appear as a contributor of the commit. But sometimes writing this line over and over feels like a hard
thing to do. At this point, this tool helps you to search, select and add the contributor to the commit message easily.

## Demo

Watch the record below. Thanks to [terminalizer](https://terminalizer.com).

![demo](demo.gif)

## Features

Thanks to [charmbracelet](https://github.com/charmbracelet) community for providing such a nice and easy TUI framework.

- [X] Filter
- [X] Selectable
- [X] Reset selections
- [X] Git hook template

## Usage

### Install

Go to the [releases](https://github.com/erkanzileli/co-author/releases) or just install the latest version.

```shell
go install github.com/erkanzileli/co-author@latest
```

Check it is accessible from your terminal

```shell
co-author
```

If it's not, then check your PATH variable because it should include the binaries installed with Go.

It might be on some places like the following

- `~/go/bin`
- `ls $(go env GOPATH)/bin`

### Add as Git Hook

It works with the `prepare-commit-msg` Git hook. You can configure it per project or globally.

**Caution:** Before running the commands below, consider checking your hook configurations for saving them instead of
overwriting.

#### Per project

If you have not a `prepare-commit-msg` defined yet, run the commands below.

```shell
co-author hook > .git/hooks/prepare-commit-msg
chmod +x .git/hooks/prepare-commit-msg
```

#### Global

If you want to use this globally, then run the commands below.

```shell
cd ~
mkdir -pv .git/hooks
git config --global --add core.hooksPath ~/.git/hooks
co-author hook > .git/hooks/prepare-commit-msg
chmod +x .git/hooks/prepare-commit-msg
```

If you already have a `prepare-commit-msg` hook defined, then take necessary output from the command below. The rest is
up to you.

```shell
co-author hook
```

## Load Config
Make it refer to the config.When there is a config, only refer to the config.
This makes it possible to target only those who participate in mob programming.

this file name only `.git-co-authors.yaml`
```yaml
committers:
  - name: user
    email: user@example.com
...
```

## Contributing

Feel free to add anything useful or fix something.

## License

MIT