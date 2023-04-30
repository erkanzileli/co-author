# co-author

A TUI that saves time from writing Co-authored-by lines to the commits in a fancy way. It works as a Git hook.

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

It works with the `prepare-commit-msg` Git hook. You can configure it per project or globally.

**Caution:** Before running the commands below, consider checking your hook configurations for saving them instead of
overwriting.

### Installation

Go to the [releases](https://github.com/erkanzileli/co-author/releases) or just install the latest version.

```shell
go install github.com/erkanzileli/co-author@latest # Install to your GOBIN directory
```

Check it is accessible from your terminal

```shell
co-author version
```

If it's not, then check your **PATH** variable because it should include the binaries installed with Go.

It might be on some places like the following

- `~/go/bin`
- `ls $(go env GOPATH)/bin`

### Per project

You can enable it for a specific project. Open a terminal in your project directory and run the commands below.

```shell
co-author hook >.git/hooks/prepare-commit-msg # Create the hook
chmod +x .git/hooks/prepare-commit-msg        # Make it executable
```

### Global

You can also enable it globally. In this way, you don't have to enable it every project you have. Run the commands
below.

```shell
cd ~ # Go to your home directory
mkdir -pv .git/hooks # Create a global Git hooks directory if it doesn't exist
git config --global --add core.hooksPath ~/.git/hooks # Set the global Git hooks directory
co-author hook > .git/hooks/prepare-commit-msg # Create the hook
chmod +x .git/hooks/prepare-commit-msg # Make it executable
```

### Migrate with the existing hook

If you already have a `prepare-commit-msg` hook defined, then take necessary output from the command below. The rest is
up to you.

```shell
co-author hook # This will print the hook template
```

### Customizations

## Load Config

You usually see the previous committers. If you want to always see a couple of person, provide a configuration file.
It will only load the committers from the file.

Create a file named `.git-co-authors.yaml`. You can use the snippet below as a template.

```yaml
committers:
  - name: user
    email: user@example.com
```

## pre-commit

[pre-commit](https://pre-commit.com) helps you integrate various Git hooks to your project with a simple file. You use
co-author
with pre-commit.

Add the snippet below to your `.pre-commit-config.yaml`

```yaml
default_install_hook_types:
  - prepare-commit-msg # it must exist

repos:
  - repo: https://github.com/erkanzileli/co-author.git
    rev: v0.0.2 # use the latest version
    hooks:
      - id: co-author
        stages:
          - prepare-commit-msg
```

After you configured the `.pre-commit-config.yaml` file, run command below to install the hooks.

```shell
pre-commit install
```

Now you can start using co-author with pre-commit.

## Contributing

Feel free to add anything useful or fix something.

## License

MIT