# Development

## local run
- `sudo /usr/loca/go/bin/go run . <command> <args>`

## Cleaning/updating dependencies
- `go mod tidy`

## Cobra

### Initialization
- `cobra init --pkg-name github.com/thetillhoff/eac`

### Adding new commands
- `cobra add <command>`
- `cobra add <subcommand> -p '<parentcommand>Cmd'`
Make sure to rename the created files to camelCase. That way it is clearer what is a command and what is a subcommand and what subcommand belongs to what command.
PascalCase is not optimal, since the variables from cobra are already automatically named camelCase, so it's easier to stay consistent that way.

## Tests

### Unit tests

The goal is to stay over 80% unit test coverage.

Run tests with `sudo /usr/local/go/bin/go test ./pkg/eac`.

Get more details with `-v`.

Detect coverage with `-cover`.

Get even more details with `-covermode=atomic`.

Save the result to a file with `-coverprofile=cover.out`.

Detect race conditions with `-race`.

Then get per-function coverage with `go tool cover -func=cover.out`.

View in-detail coverage with `go tool cover -html=cover.out`.

In total:
- `sudo /usr/local/go/bin/go test ./pkg/eac -v -cover -covermode=atomic -coverprofile=cover.out -race`
- `go tool cover -func=cover.out` or `go tool cover -html=cover.out`

### `commander-cli`
Uses https://github.com/commander-cli/commander for testing the cli as a whole.

Adding a new test case: `commander add --file=./.commander.yaml <command>`

Running tests: `commander test --file=./.commander.yaml`

## Notes

### Project structure
> Let Go pkgs in your internal directory be support packages that
> - will only return error in case of problems
> - won't log (console or otherwise)
> - won't panic
> Let your application (packages in your cmd directory) decide what the appropriate behavior in case of an error (log / graceful shutdown / recover to 100% integrity)
