# Context

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zpatrick/context/blob/master/LICENSE)

Context is a tool that aims to make it easier to manage shell configurations.
I often find myself tediously setting and/or updating environment variables when switching contexts or working on multiple projects at once. Context provides a light wrapper that allows you to view, list, edit, and delete different bash scripts meant to be `source`'d into your current shell. 

# Installation
Download the binary for your OS from the [release](https://github.com/zpatrick/context/releases) page and name it **context**. 

In your `~/.bash_profile`, add the following function:
```
switch() {
  source ~/.context/"$@"
}
```

Once configured, you can easily switch between contexts using:
```
$ switch CONTEXT
```

# Usage

**Create/Edit**: to create or edit a context, use the `edit` command:
```
$ context edit CONTEXT
```

This will open up a text editor (default: `vim`) for the specified `CONTEXT`. 
You can choose which editor to use by setting `CTX_EDITOR` environment variable, 
or by using the `--editor` flag. 

**List**: to list your contexts, use the `list` command:
```
$ context list
dev
prod
test
```

**Delete**: to delete a context, use the `delete` command:
```
$ context delete test
remove context 'test'?: y
```

**Switch**: the switch command will print the contents of a context to stdout. 
This is meant to be used in conjunction with the `source` function (see [installation](#installation)):
```
$ context switch dev
export X_VAR_A=123
export X_VAR_B=456

$ source <(context switch dev)
```

# License
This work is published under the MIT license.

Please see the `LICENSE` file for details.
