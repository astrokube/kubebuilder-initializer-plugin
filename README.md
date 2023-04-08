[![GitHub Release](https://img.shields.io/github/v/release/astrokube/kubebuilder-initializer-plugin)](https://github.com/astrokube/kubebuilder-initializer-plugin/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/astrokube/kubebuilder-initializer-plugin.svg)](https://pkg.go.dev/github.com/astrokube/kubebuilder-initializer-plugin)
[![go.mod](https://img.shields.io/github/go-mod/go-version/astrokube/kubebuilder-initializer-plugin)](go.mod)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://img.shields.io/github/license/astrokube/kubebuilder-initializer-plugin)
[![Build Status](https://img.shields.io/github/actions/workflow/status/astrokube/kubebuilder-initializer-plugin/build.yml?branch=main)](https://github.com/astrokube/kubebuilder-initializer-plugin/actions?query=workflow%3ABuild+branch%3Amain)
[![CodeQL](https://github.com/astrokube/kubebuilder-initializer-plugin/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/astrokube/kubebuilder-initializer-plugin/actions/workflows/codeql.yml)
---

# Kubebuilder Initializer Plugin

A powerful Kubebuilder plugin to initialize dynamically the structure of your kubebuilder operator project.

## Prerequisites

This is a plugin for the kubebuilder cli tool If you don't have the Kubebuilder cli installed in your computer, please
visit the official documentation, [kubebuilder documentation](https://github.com/kubernetes-sigs/kubebuilder).

## Installation

We provide you a variety of alternatives to install the plugin in your computer, take the onw that best fits your needs.

### Download the executable

1. Visit the latest release at [Release page](https://github.com/astrokube/kubebuilder-initializer-plugin/releases)
2. Download the version that works for you
3. Extract the tarball
4. Copy the executable file to the path used by Kubebuilder to read the external plugins
   - OSX:  ~/Library/Application\ Support/kubebuilder/plugins/kubebuilder-initializer/v1-alpha
   - Linux: $HOME/.config/kubebuilder/plugins/kubebuilder-initializer/v1-alpha

### From the code

To compile the code from your own computer, you just need to run the following commands

```bash
git clone https://github.com/astrokube/kubebuilder-initializer-plugin.git
cd kubebuilder-initializer-plugin
make build install
```

To check that the installation was success, verify that the executable was copied to the below path (depending on the
operating system)

- OSX:  ~/Library/Application\ Support/kubebuilder/plugins/kubebuilder-initializer/v1-alpha
- Linux: $HOME/.config/kubebuilder/plugins/kubebuilder-initializer/v1-alpha


### Homebrew

Additionally,  you could download the executable from our Astrokube Brew repository.

1. Add the Astrokube Brew repository
```bash
brew tap astrokube/tools
```

2. Install Kubebyuilder Initializer plugin:
```bash
brew install kubebuilder-initializer-plugin
```

3. You will need to create a symbolic link to the path used by Kubebuilder to read the external plugins.

```bash
mkdir -p ~/Library/Application\ Support/kubebuilder/plugins/kubebuilder-initializer-plugin/v1-alpha/
ln -s /usr/local/Cellar/kubebuilder-initializer-plugin/0.1.0/bin/kubebuilder-initializer-plugin \
  ~/Library/Application\ Support/kubebuilder/plugins/kubebuilder-initializer-plugin/v1-alpha/kubebuilder-initializer-plugin
```

### Verify the installation

Once the plugin is installed in your computer, a new plugin is available for you to be used when running the Kubebuilder
cli tool, you can run `kubebuilder help` to check it.

The `kubebuilder-initializer-plugin/v1-alpha` appears in the list of available plugins.

![Kubebuilder pLugins](docs/assets/plugins.png)

## Define your own template

The Kubebuilder Initializer plugin understand a template like a Git repository in which the name of the elements in the repository
(both folders and files) and the content of the files can contain variables. 

The plugin takes advantage of Go templates to process the templates;  that provides us with a very flexible way to define templates.
That implies that we could not only to define variables 

[go.mod](https://github.com/astrokube/kubebuilder-operator-template/blob/main/go.mod#L1)
```text
module {{.repository.server}}/{{.repository.owner}}/{{.repository.name}}
```

but also add some logic to our own templates

[OWNERS_ALIASES](https://github.com/astrokube/kubebuilder-operator-template/blob/main/OWNERS_ALIASES#L4-L9)
```text
# See the OWNERS docs: https://git.k8s.io/community/contributors/guide/owners.md

aliases:
{{- range .owners}}
  {{.alias}}:
{{- range .members}}
    - [{{.}}](https://github.com/{{.}})
{{- end}}
{{- end}}
```

You can find some examples of templates in [AWESOME_TEMPLATES.md](AWESOME_TEMPLATES.md), and we encourage you to contribute
with your own templates, so please, feel free to open a pull request with an entry in this file if you want to share a 
template with others.

**TODO**
We ask you to share an example files with the variables that need to be passed in order to customize the templates,
See an example [here](). Only yaml files are supported (JSON could be supported for futures release If this was required  
by the community )

*For upcoming releases, the plugin will inspect the files in the templates and It will be able to generate the yaml file for you*

## Getting started

This plugin is used exclusively in the initial scaffolding (`kubebuilder init`) and It's compatible with any other plugin.
When we run the `init` command, the Kubebuilder cli creates  the PROJECT file, this  is  the main piece for Kubebuilder 
to create consistency and being  able to inject code when we run  other commands such as `kubebuildfer create api` or 
`kubebuilder create webhook` 
On the other hand, the `Kubebuilder Initializer plugin` must be used in conjunction with other plugins that will
take the control once we need to create a Webhook or an API.

To take advantage of the Initializer plugin, we just need a repository, that will be used as a template, and the 
variables file  that will allow us to customize the template. By the default, the plugin read the variables from a 
named file `.kubebuilder-layout.yaml`, but this can be customized If required. 

To sum up, to initialize our project we just need to pass the argument `--from` and we could additionally pass the argument
`--vars` in case of we don't want to use the default `.kubebuilder-layout.yaml`.

In the below example, we would use our plugin in conjunction with the  `go.kubebuilder.io/v3` that help us to work with 
implementation of operators in Go.

```bash
kubebuilder init  --plugins go.kubebuilder.io/v3,kubebuilder-initializer-plugin/v1-alpha \
  --domain astrokube --owner astrokube --repo github.com/astrokube/k8s-testing-operator \
  --from "https://github.com/astrokube/kubebuilder-operator-template"
```

Be aware that, in the above example, the arguments `domain`, `repo` and `owner` are required by the plugin 
`go.kubebuilder.io/v3`.

**Non default branches**

By default, the plugin will fetch the code in the default branch, but we can specify another branch:

```bash
--from "github.com/astrokube/kubebuilder-operator-template#<branch>"
````

**With credentials**

For those repositories that require authentication we can provide the user credentials or a token as It's shown on
the below

```bash
--from "<user>:<password>@github.com/astrokube/kubebuilder-operator-template"
...
 --from "<token>@github.com/astrokube/kubebuilder-operator-template"
```


## Kubebuilder

To deep dive into how Kubebuilder deals with external plugins you can visit the following article
[Extensible CLI and Scaffolding Plugins - Phase 2](https://github.com/kubernetes-sigs/kubebuilder/blob/master/designs/extensible-cli-and-scaffolding-plugins-phase-2.md)



## Contributing

Visit the [CONTRIBUTING.md](CONTRIBUTING.md) file.