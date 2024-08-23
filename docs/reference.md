# CLI & references

## `gookme init`

The main initialization command. It :

- **creates** a `hooks` folder at the root of your project where you can write **project-wide hooks** that will be triggered on every commit
- **writes** `.git/hooks` files

### Options for `gookme init`

| option  | short | description |
|---------|-------|-------|
| --types | -t    | The types of Git hooks to hook. Accepted values are: pre-commit,  prepare-commit-msg, commit-msg,  post-commit, post-merge, post-rewrite,  pre-rebase, post-checkout, pre-push |
| --all   | -a    | Initialize all available hooks. Has precedence over the --types flag |

### Examples for `gookme init`

```sh title="Initialize Gookme with pre-commit and commit-msg hooks"
gookme init --types pre-commit,commit-msg
```

```sh title="Initialize Gookme with all available hooks"
gookme init --all
```

## `gookme run`

Mainly used for debugging and dry run :

### Options for `gookme run`

| option  | short | description |
|---------|-------|-------|
| --type | -t    | The types of Git hook to run. Accepted values are: pre-commit,  prepare-commit-msg, commit-msg,  post-commit, post-merge, post-rewrite,  pre-rebase, post-checkout, pre-push |
| --from   | -f    | (optional) Starting git reference used to evaluate hooks to run. If set, `to` has to be set as well, otherwise this option is ignored. |
| --to   | -o    | (optional) Ending git reference used to evaluate hooks to run. If set, `from` has to be set as well, otherwise this option is ignored. |

### Examples for `gookme run`

```sh title="Run all hooks for the pre-commit type"
gookme run --type pre-commit <args>
```

## Hook files

### General description

See [Writing your hooks](getting-started.md#writing-your-hooks)

### Available options

- `steps`

The list of steps (commands) being executed by this hook. In a step you can define :

#### Step options

| Option        | Description           | Required  |
| ------------- | ------------- | ------|
| `name`      | The name that will be given to this step | yes |
| `command`      | The command invoked at this step |   yes |
| `onlyOn` | A shell wildcard (or a list of wildcard) conditioning the execution of the step based on modified files      |    no |
| `serial` | A boolean value describing if the package hook execution should await for the step to end |    no |
| `from` | Extend a shared step |    no |

!!! warning
    Gookme exits with a non-zero status code if any of the steps fails, as soon as one fails.

!!! warning
    The pattern provided in `onlyOn` will be matched agains the relative path of matched files of the execution, from the package folder, not from the repository root.

### Available arguments

A set of arguments is provided by Mookme, that can be directly used in the hooks command definitions using the following syntax in the step definition:

````json title="hooks/commit-msg.json"
{
    "command": "run-something $1" // Will be replaced with the value of `args`
}
````

- `$1`

The argument being passed by git to the hook file. See [the Git documentation on the hooks](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks) for more details about what it contains depending on the hook type being executed.
