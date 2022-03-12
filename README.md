# co-author

A TUI that saves time from writing Co-authored-by lines to the commits in a fancy way.

## Demo

Watch the [record](https://terminalizer.com/view/d12a3f1a5606). Thanks to [terminalizer](https://terminalizer.com).

## Features

Thanks to [charmbracelet](https://github.com/charmbracelet) community for providing such a nice and easy TUI framework.

- [X] Filter
- [X] Selectable
- [X] Reset selections

## Usage

### Install it first

Go to the [releases](https://github.com/erkanzileli/co-author/releases) or just install the latest version.

```shell
go install github.com/erkanzileli/co-author@latest
```

### Add to your project hooks

It works with the `prepare-commit-msg` Git hook. So you can use predefined template on your hook.

If you already have a `prepare-commit-msg` then take necessary output from the `hook` command.

```shell
co-author hook
```

If you have not a `prepare-commit-msg` defined yet, run command below.

```shell
co-author hook > .git/hooks/prepare-commit-msg
chmod +x .git/hooks/prepare-commit-msg
```

## Contributing

Feel free to add anything useful or fix something.

## License

MIT