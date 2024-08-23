# Features

## Reusable steps

`Gookme` provides you with step-sharing features, allowing you to declare shared step example, and to use them in your steps.

Given a project directory such as this:

``` title="Project structure"
project-root
|--- hooks
    |--- shared
        |--- flake8.json
|--- packages
    |--- some-package
        |--- hooks
            |--- pre-commit.json
```

!!! tip
    The `hooks/shared` folder is automatically generated with a `.gitkeep` file by `gookme init`.

You can declare a step in `hooks/shared/flake8.json` ...

```json title="hooks/shared/flake8.json"
{
  "name": "Ensure Style (flake8)",
  "command": "flake8 $(python-module) --ignore E203,W503 --max-line-length 90",
  "onlyOn": "**/*.py"
}
```

... and then re-use it in `some-package/hooks/pre-commit.json` with the `from` keyword:

```json title="some-package/hooks/pre-commit.json"
{
  "steps": [
    {
      "from": "flake8"
    },
    ... // some other steps
  ]
}
```

## Writing and using utils scripts

It is possible to declare some scripts in the project root `Gookme` hooks folder, and then use them directly in the commands invoked by the steps.

Given a project directory such as this:

```sh title="Project structure"
project-root
|--- .mookme.json
|--- hooks
    |--- partials
        |--- pylint-changed-files
 |--- packages
    |--- some-package
        |--- .hooks
            |--- pre-commit.json
```

*Here is how the `python-changed-files` script looks like*

```bash title="project-root/hooks/partials/pylint-changed-files"
#!/usr/bin/env bash
git --no-pager diff --cached --name-only --diff-filter=AM --relative -- "***.py" | tr '\n' '\0' | xargs -0 "$@"
```

!!! tip
    The `hooks/partials` is automatically generated with a `.gitkeep` file by `gookme init`.

One can declare a script in flake8 (don't forget to `chmod+x`) and then re-use it in `some-package/hooks/pre-commit.json` by directly invoking the script's name:

```json title="some-package/hooks/pre-commit.json"
{
  "steps": [
    {
      "name": "Run pylint but only on changed files",
      "command": "python-changed-files pylint"
    },
    ... // some other steps
  ]
}
```

## Use a range of commits

Using the Gookme CLI, it is possible to invoke a set of hooks and steps selected using the files changed between two git references.

````bash
gookme run -t pre-commit --from HEAD~1 --to f9ff43
gookme run -t pre-commit --from HEAD~25 --to d58688dd611ef01079f61ebae36df0ce8c380ddb
````

You can find more details about these options on the [gookme run reference](reference.md#gookme-run) page.
