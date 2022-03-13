# co-author

A TUI that saves time from writing Co-authored-by lines to the commits in a fancy way.

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

## Usage

### Install

Go to the [releases](https://github.com/erkanzileli/co-author/releases) or just install the latest version.

```shell
go install github.com/erkanzileli/co-author@latest
```

### Add as Git Hook

It works with the `prepare-commit-msg` Git hook.

If you have not a `prepare-commit-msg` defined yet, run command below.

```shell
co-author hook > .git/hooks/prepare-commit-msg
chmod +x .git/hooks/prepare-commit-msg
```

If you already have a `prepare-commit-msg` hook defined, then take necessary output from the command below.

```shell
co-author hook
```

## Contributing

Feel free to add anything useful or fix something.

## License

MIT