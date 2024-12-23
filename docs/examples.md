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
