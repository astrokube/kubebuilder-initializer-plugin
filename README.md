[![GitHub Release](https://img.shields.io/github/v/release/astrokube/kubebuilder-initializer-plugin)](https://github.com/astrokube/kubebuilder-initializer-plugin/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/astrokube/kubebuilder-initializer-plugin.svg)](https://pkg.go.dev/github.com/astrokube/kubebuilder-initializer-plugin)
[![go.mod](https://img.shields.io/github/go-mod/go-version/astrokube/kubebuilder-initializer-plugin)](go.mod)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://img.shields.io/github/license/astrokube/kubebuilder-initializer-plugin)
[![Build Status](https://img.shields.io/github/actions/workflow/status/astrokube/kubebuilder-initializer-plugin/build.yml?branch=main)](https://github.com/astrokube/kubebuilder-initializer-plugin/actions?query=workflow%3ABuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/astrokube/kubebuilder-initializer-plugin)](https://goreportcard.com/report/github.com/astrokube/kubebuilder-initializer-plugin)
[![CodeQL](https://github.com/astrokube/kubebuilder-initializer-plugin/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/astrokube/kubebuilder-initializer-plugin/actions/workflows/codeql.yml)
---

# Kubebuilder Initializer Plugin

A powerful Kubebuilder plugin to initialize dynamically the structure of your kubebuilder operator repositories.

## Getting started


## Installation

### Prerequisites

This is a plugin for the kubebuilder tool. In case of you haven't installed the tool yet, please visit the 
[kubebuilder documentation](https://github.com/kubernetes-sigs/kubebuilder) and follow the instructions to get 
kuberbuilder properly installed in your computer.

### Homebrew

### Download the executable files

### Build from the code

```bash
git clone https://github.com/astrokube/kubebuilder-initializer-plugin.git
cd kubebuilder-initializer-plugin
make build install
```

The executable will be copied to the following directory *~/Library/Application\ Support/kubebuilder/plugins/kubebuilder-initializer/v1-alpha*
If your os is OSX or *$HOME/.config/kubebuilder/plugins/kubebuilder-initializer/v1-alpha* for linux users.

## How this plugin work

To deep dive into how Kubebuilder deals with external plugins you can visit the following article 
[Extensible CLI and Scaffolding Plugins - Phase 2](https://github.com/kubernetes-sigs/kubebuilder/blob/master/designs/extensible-cli-and-scaffolding-plugins-phase-2.md) 

## Contributing

Visit the [CONTRIBUTING.md](CONTRIBUTING.md) file.