[![GitHub Release](https://img.shields.io/github/v/release/astrokube/kubebuilder-initializer-plugin)](https://github.com/astrokube/kubebuilder-initializer-plugin/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/astrokube/kubebuilder-initializer-plugin.svg)](https://pkg.go.dev/github.com/astrokube/kubebuilder-initializer-plugin)
[![go.mod](https://img.shields.io/github/go-mod/go-version/astrokube/kubebuilder-initializer-plugin)](go.mod)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://img.shields.io/github/license/astrokube/kubebuilder-initializer-plugin)
[![Build Status](https://img.shields.io/github/actions/workflow/status/astrokube/kubebuilder-initializer-plugin/build.yml?branch=main)](https://github.com/astrokube/kubebuilder-initializer-plugin/actions?query=workflow%3ABuild+branch%3Amain)
[![CodeQL](https://github.com/astrokube/kubebuilder-initializer-plugin/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/astrokube/kubebuilder-initializer-plugin/actions/workflows/codeql.yml)
---

# Kubebuilder Initializer Plugin

A powerful Kubebuilder plugin to initialize dynamically the structure of your kubebuilder operator project.

### Prerequisites

This is a plugin for the kubebuilder cli tool If you don't have the Kubebuilder cli installed in your computer, please
visit the official documentation, [kubebuilder documentation](https://github.com/kubernetes-sigs/kubebuilder).

### Installation

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


#### Homebrew

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

## Getting started

Once the plugin is installed in your computer, a new plugin is available for you to be used when running the Kubebuilder
cli tool, you can run `kubebuilder help` to check it.

The `kubebuilder-initializer-plugin/v1-alpha` appears in the list of available plugins.

![Kubebuilder pLugins](docs/assets/plugins.png)

This plu

2. Choose the template for scaffolding the initial structure of our Kubebuilder operator. You can 
create your own template as described (here]() or alternatively you could take advantage of some of the well-known templates
that you can find in [AWESOME_TEMPLATES.md](AWESOME_TEMPLATES.md)

3. Once we have chosen the template that we want to use, we just need to write the yaml file  that contains the values that 
will make us to customize the template. By default, the plugin will take a file named `.kubebuilder-layout.yaml`, otherwise 
you will need to pass an extra argument with the path to the file.

4. Initialize your project. Keep in mind that this plugin is used exclusively to initialize our project structure, so we should
use also a plugin that supports the APIs and webhooks creation,for instance the `go.kubebuilder.io/v3` that is prpvided out  of the box by Kubbebuilder.

```bash
kubebuilder init  --plugins go.kubebuilder.io/v3,kubebuilder-layout/v1-alpha \
  --from "github.com/astrokube/kubebuilder-operator-template" \
  --domain astrokube \
  --owner astrokube \
  --repo github.com/astrokube/k8s-testing-operator
```

The only argument that needs to be passed is the `from`

*In the above example, the args `domain`, `repo` and `owner` are required by the plugin `go.kubebuilder.io/v3`.*

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