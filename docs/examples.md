# Examples

## `commitlint`

*This guide will help you in the setup of a hook a hook for linting your commits with [commitlint](https://github.com/conventional-changelog/commitlint)*

### Prerequisite for Commitlint

- You have installed `Gookme`
- You have installed and configured [commitlint](https://github.com/conventional-changelog/commitlint)
- You have setup `Gookme` using `gookme init` (see [get started](getting-started.md) if needed)

### Hook for Commitlint

In the global hooks folder of your project `hooks/commit-msg.json` add the following configuration :

```js title="hooks/commit-msg.json"
{
    "steps": [
        // ...
        // your other steps
        {
            "name": "Lint commit message",
            "command": "npx commitlint --edit $1"
        }
        // ...
        // your other steps
    ]
}
```

### Running multiple steps sequentially

It is possible to run multiple steps sequentially by setting the `serial` option to `true`. When a step is set to `serial`, the next step will only be executed when the step marked as `serial` has succeeded. Previous steps will be executed in parallel unless they are also set to `serial`.

```json title="hooks/commit-msg.json"
{
  "steps": [
    {
      "name": "go vet",
      "command": "go vet ./...",
      "serial": true
    },
    {
      "name": "format go code",
      "command": "go fmt ./...",
      "serial": true
    },
    {
      "name": "test go code",
      "command": "go test ./...",
      "serial": true
    },
    {
      "name": "lint go code",
      "command": "golangci-lint run ./...",
      "serial": true
    }
  ]
}
```
