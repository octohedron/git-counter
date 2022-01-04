# git-counter

A basic commit time-of-day counter with support for multiple directories

---

## USAGE

1. `cd` to `$GOPATH/src/github.com/octohedron`
2. `git clone git@github.com:octohedron/git-counter.git && cd git-counter`
3. `go build`
4. `./git-counter -dir=/full/path/to/folder/with/git/repos -author="Your name"`

## Examples

```
$ go run main.go -dir=/full/path1... -dir=/full/pathN... -author='User'
$ go run main.go -dir=/home/user/go/src/github.com/user
```

## Biggest github repos

A folder with the biggest repos on github (run `download.sh` to try it)

- AdaCore/gnatstudio
- AdaCore/spark2014
- adonovan/gopl.io
- ansible/ansible
- ansible/awx
- aosp-mirror/platform_development
- aosp-mirror/platform_frameworks_base
- aosp-mirror/platform_frameworks_support
- aosp-mirror/platform_system_core
- apache/groovy
- apache/hadoop
- apache/httpd
- apache/maven
- apache/spark
- apple/foundationdb
- apple/swift
- atom/atom
- Autodesk/maya-usd
- blender/blender
- chromium/chromium
- dart-lang/sdk
- django/django
- dlang/dmd
- dlang/phobos
- docker/docker-ce
- docker/dockerhub.io
- dotnet/aspnetcore
- dotnet/fsharp
- dotnet/roslyn
- dotnet/runtime
- electron/electron
- elixir-lang/elixir
- emacs-mirror/emacs
- EpicGames/UnrealEngine
- erlang/otp
- ethereum/solidity
- facebook/react-native
- fish-shell/fish-shell
- flutter/flutter
- gcc-mirror/gcc
- ghc/ghc
- GNOME/gimp
- gohugoio/hugo
- golang/go
- google/jax
- Homebrew/homebrew-core
- JetBrains/intellij-community
- JetBrains/kotlin
- JuliaLang/julia
- jupyter/notebook
- kubernetes/kubernetes
- llvm/llvm-project
- llvm/phabricator
- magento/magento2
- mdaniel/virtualbox-org-svn-vbox-trunk
- meteor/meteor
- microsoft/TypeScript
- microsoft/vscode
- mirror/vbox
- mongodb/mongo
- mozilla/gecko-dev
- mysql/mysql-server
- NixOS/nixpkgs
- nodejs/node
- openjdk/jdk
- openjdk/jdk14u
- openjdk/loom
- openshift/installer
- openshift/openshift-ansible
- openshift/origin
- ornicar/lila
- Perl/perl5
- php/php-src
- postgres/postgres
- python/cpython
- pytorch/pytorch
- ruby/ruby
- rust-lang/cargo
- rust-lang/rust
- scala/scala
- sourcegraph/sourcegraph
- SWI-Prolog/swipl
- SWI-Prolog/swipl-devel
- tcltk/tcl
- tensorflow/tensorflow
- torvalds/linux
- wch/r-source

```
MAX 761,865 TOTAL 10,812,030
  0  367576 ***************************************
  1  273885 *****************************
  2  212318 ***********************
  3  177053 *******************
  4  156972 *****************
  5  155261 *****************
  6  171967 *******************
  7  212361 ***********************
  8  305061 *********************************
  9  446704 ***********************************************
 10  571736 *************************************************************
 11  629134 *******************************************************************
 12  584214 **************************************************************
 13  637317 *******************************************************************
 14  720788 ****************************************************************************
 15  759528 ********************************************************************************
 16  761865 ********************************************************************************
 17  713516 ***************************************************************************
 18  600517 ****************************************************************
 19  514272 *******************************************************
 20  488211 ****************************************************
 21  488404 ****************************************************
 22  457511 *************************************************
 23  405859 *******************************************
2022/01/04 02:28:09 13.260406042s
```

## Notes

- `-author` flag is optional and accepts regex (same as git), i.e.
  `-author="Name.*"`
- Allows for many paths, passing the same flag, i.e. `-dir=path1 -dir=path2 ...`
- Only supports folders (i.e. `-dir=path` flag) that have folders with `.git`
  directories on them

## Motivation

Other tools such as [knadh/git-bars](https://github.com/knadh/git-bars) don't
support multiple directories
