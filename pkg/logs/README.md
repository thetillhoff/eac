# Logs

This package is designed to give neatly formatted logging output.
There are four types:
```
Info("Additional information for debugging purposes.", obj...)
Success("A success message.", obj...)
Warn("A warning,", obj...)
Error("An error message that will be output to stderr.", obj...)
```

- Info text is prefixed with "INF" in normal text color.
- Success text is prefixed with "<span style="color:green">SUC</span>" in green color.
- Warnings are prefixed with "<span style="color:yellow">WRN</span>" in yellow color.
- Error text is prefixed with "<span style="color:red">ERR</span>" in red color

## Additional options
This package has two public options.

One is `ContinueOnError` (default: false). It determines whether the application should exit when `Error()` is called.

The other is `Verbose` (default: false). It determines whether `Info()` messages are output or not.
