# Getting started

## Installation

=== "Linux"

    ``` bash title="Install Gookme"
    curl -sSL https://raw.githubusercontent.com/LMaxence/gookme/main/scripts/install.sh | bash
    ```

=== "MacOS"

    ``` bash title="Install Gookme"
    curl -sSL https://raw.githubusercontent.com/LMaxence/gookme/main/scripts/install.sh | bash
    ```
=== "Windows"

    ``` bash title="Install Gookme"
    Invoke-WebRequest -Uri https://raw.githubusercontent.com/LMaxence/gookme/main/scripts/install.ps1 -OutFile install.ps1
    .\install.ps1
    ```

## Configuration

Gookme will automatically detect the `hooks` folder across your repository and trigger the command related to your current VCS state.

Hence it only requires a very minimal configuration, as most of this is defined by where you place your `hooks` folders, and what you put in them.

### Setup Git to run Gookme

To setup Gookme, you need to create scripts in the `.git/hooks` folder, or add a few lines in the existing ones. **This is done automatically by running the `init` command**

``` sh title="Initialize Gookme with all Git hooks"
gookme init --all
```

### Selectively setup Gookme

If you want to only setup the hooks for a specific hook type, you can use the `--types` option of the `gookme init` command.

```bash title="Initialize Gookme with pre-commit and commit-msg hooks"
gookme init --types pre-commit,commit-msg
```

### Removing Gookme from your Git hooks scripts

If you want to remove Gookme from your Git hooks scripts, you can use the `clean` command.

```bash title="Remove Gookme from your Git hooks scripts"
gookme clean
```

This will remove all the lines added by Gookme in your Git hooks scripts.

## Writing your hooks

### Global structure of your project hooks

`Gookme` is designed for monorepos, hence it assumes your project has a root folder where global hooks can be defined, and multiple packages where you can define per-package hook.

!!! tip
    Hooks are written in a folder `hooks` located at the root of your project and at the root of your packages' folders.


    **When using `Gookme` in a monorepo, you will have a project structure following this :**

    ```bash
    <root> # where your .git is located
    |- hooks # will always be executed when you commit
    |  |- pre-commit.json # will be executed with the pre-commit git hook
    |  |- commit-msg.json  # will be executed with the commit-msg git hook
    |  |- prepare-commit-msg.json
    |  |- post-commit.json
    |- packages
    |  |- package A
    |  |  |- hooks # will be executed if you commit changes on package A
    |  |  |  |- pre-commit.json
    |  |  |  |- post-commit.json
    |  |- package A
    |  |  |- hooks # will be executed if you commit changes on package B
    |  |  |  |- pre-commit.json
    ```

With `Gookme`, your hooks are stored in JSON files called `{hook-type}.json` where the hook type is one of the available git hooks, eg :

- `pre-commit`
- `prepare-commit-msg`
- `commit-msg`
- `post-commit`
- `post-merge`
- `post-rewrite`
- `pre-rebase`
- `post-checkout`
- `pre-push`

!!! warning
    If the command executed by a hook fails, it will prevent the git command to be executed. We recommend you to use the pre-receive hooks carefully, with relatively safe commands, otherwise you might prevent your team for doign stuff like `git pull` or `git fetch`.


### How will Gookme decide which hooks to run ?

1. Gookme will pick up files concerned by an execution. The result of this step is a list of relative paths from the root of the repository.

2. For each folder where a `hooks` folder is found, Gookme will assess if there are file paths in the previous list matching the relative path to this folder from the root of the repository.

3. Other selections (`onlyOn` for instance) are applied on each step of each package, based on the list of paths attached to the package and it's steps.

!!! tip
    Depending on the hook type being executed, or the arguments passed to the command, the list of paths can be different.

    - The hook type `pre-push` will consider the list of files to be pushed to the remote server
    - The hook type `post-commit` will consider the list of files included in the last commit
    - If the `run` command argument --from and --to are used, the list of paths will be the list of files changed between the two commits
    - All other hook types will use the list of files changed in the current commit, and staged

### Example of hook file

Your hooks are defined in simple json files.

- For complete reference, see the [JSON hooks reference](reference.md#hook-files)
- For specific hook examples, see the [recipes](examples.md)

A hook defines a list of `steps`, which are basically commands to run, with a name for proper display. A few configuration options are available, but the minimal requirement is `name` and `command`.

Here is an example that will run your commit message using `commitlint`.

``` js title="hooks/commit-msg.json"
{
    "steps": [{
        "name": "Lint commit message",
        "command": "commitlint lint --message $1"
    }]
}
```

!!! tip
    When writing package-scoped hooks, the current working directory assumed by `Gookme` is the folder where this package's `hooks'` folder is located

!!! warning
    The token `$1` in the hook command is replaced with the hook arguments when the command is executed. See [the  Git documentation on hooks](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks)

**More examples to [get you started](examples.md)**!
