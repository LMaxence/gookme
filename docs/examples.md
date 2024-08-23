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
        "name": "commit lint",
        "command": "cat {args} | ./node_modules/@commitlint/cli/cli.js"
    }
    // ...
    // your other steps
    ]
}
```

## `eslint`

*This guide will help you in the setup of a hook for linting your code with [eslint](https://eslint.org/)*

### Prerequisite for Eslint

- You have installed `mookme`
- You have installed and configured [eslint](https://eslint.org/)
- You have setup `mookme` using `npx mookme init` (see [getting started](getting-started.md) if needed)

### Hook for Eslint

- In the hooks folder of the package you want to lint with `eslint` `<project-root>/<package>/.hooks/commit-msg` add
the following configuration :

```js title="hooks/commit-msg.json"
{
    "steps": [
    // ...
    // your other steps
    {
        "name": "eslint",
        "command": "./node_modules/eslint/bin/eslint.js ."
    }
    // ...
    // your other steps
    ]
}
```

*Alternative: setup a `npm` script and directly invoke `eslint` in the command field :*

```js title="hooks/commit-msg.json"
{
    "steps": [
    // ...
    // your other steps
    {
        "name": "eslint",
        "command": "npm run lint"
    }
    // ...
    // your other steps
    ]
}
```
