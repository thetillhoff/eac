# Development

## local run
- `go run . <command> <args>`

## Cleaning/updating dependencies
- `go mod tidy`

## Cobra

### Initialization
- `cobra init --pkg-name github.com/thetillhoff/eac`

### Adding new commands
- `cobra add <command>`
- `cobra add <subcommand> -p '<parentcommand>Cmd'`
Make sure to rename the created files to camelCase. That way it is clearer what is a command and what is a subcommand and what subcommand belongs to what command.

## Notes

### Project structure
> Let Go pkgs in your internal directory be support packages that
> - will only return error in case of problems
> - won't log (console or otherwise)
> - won't panic
> Let your application (packages in your cmd directory) decide what the appropriate behavior in case of an error (log / graceful shutdown / recover to 100% integrity)
