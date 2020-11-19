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

To run `write` command:
```
go run . write 'echo "This is the command"'
```

## Install

- Install to your local go path with `go install .`
- Add the local go path to your `.zshrc`: 
```bash
export PATH="$HOME/go/bin:$PATH"
```
- Source the key-bindings:
```
source shell/key-bindings.zsh
```
 