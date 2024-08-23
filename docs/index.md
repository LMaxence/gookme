# Welcome on Gookme

*A simple and easy-to-use, yet powerful and language agnostic git hook for monorepos.*

## What is Gookme ?

Gookme is a Git hook manager. It's sole purpose is to execute some scripts when you want to commit or do other git stuff. It could be a
linter, tests, your favorite commit message checker.

!!! note "What can I run with Gookme ?"

    Everything that is invoked through a cli can be used with Gookme!

You are welcome to use it and enjoy it's simplicity.
**If you encounter any bug or weird behavior, don't be afraid to open an [issue](https://github.com/LMaxence/gookme/issues/new/choose) :)**

## How does it work ?

1. Gookme is invoked by a git hook script
2. Gookme looks for `hooks` folders across your repository
3. For each detected folders, Gookme detects if there are any hooks defined for the git hook currently being executed
4. For each detected folders, Gookme detects if there are git-staged changes in this folder
5. If both conditions above are valid, Gookme runs concurrently (or not) the different commands provided in the hook files.

## Why not ... ?

### `bash` scripts

**`bash` scripts directly written in my `.git/hooks` folder**

- Even if it is true that, in the end, `Gookme` will do nothing more than invoking commands the exact same way a bash
script would, the `.git/hooks` folder is a not a versioned one, *hence it will prevent you from sharing configuration*.
- `Gookme` provides you with a way to version these hooks, and to share repository configuration among the rest of your team.
- The hook setup is a one liner for the new developers landing in your team. It won't download anything, just write a
small line in your `.git/hooks` files

### [`pre-commit`](https://pre-commit.com/)

!!! info
    `pre-commit` was the tool I used before developing Mookme, then Gookme.

There were several issues with `pre-commit`, that led me to develop my own tool :

- pre-commit is not designed for monorepos, hence most of the hooks are some sort of hacks
- per-package environment is not easy to manage, because `pre-commit` has it's own global environment and we have to create global dependency to run a particular hook for one package.

!!! warning
    This led us to one of the guideline used by `Gookme` to work:
    If we run a hook on a package in your monorepo it means:

    - There are changes in the folder of this package
    - The dev environment of this package is properly installed (Gookme does not proceed to any additional setup)