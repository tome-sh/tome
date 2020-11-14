# Tome

Share shell spells.

## Config

You'll need a config at ~/.tome.yaml looking like this:
```
shellType: zsh
repository: /Users/username/code/tome-repo/commands
```

## Development

Run tests:
```
go test -v ./...
```

Note: To run, you will need to provide a config file (default: `~/.tome.<EXTENSION>`).
For supported EXTENSIONs see the [viper doc](https://github.com/spf13/viper#what-is-viper).

To run the code showing the help text:
```
go run .
```
