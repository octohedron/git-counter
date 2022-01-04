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

- gohugoio/hugo
- gethugothemes/navigator-hugo
- adonovan/gopl.io
- sourcegraph/sourcegraph\
- EpicGames/UnrealEngine
- torvalds/linux
- kubernetes/kubernetes
- tensorflow/tensorflow
- Homebrew/homebrew-core
- microsoft/vscode
- ornicar/lila
- ansible/ansible
- ansible/awx
- django/django
- flutter/flutter
- rust-lang/cargo
- rust-lang/rust
- facebook/react-native
- electron/electron
- apple/swift
- nodejs/node
- gcc-mirror/gcc
- llvm/llvm-project
- llvm/phabricator
- apache/httpd
- apache/spark
- pytorch/pytorch
- mongodb/mongo
- google/jax
- microsoft/TypeScript
- php/php-src
- golang/go
- python/cpython
- atom/atom
- meteor/meteor
- mirror/vbox
- aosp-mirror/platform_frameworks_base
- aosp-mirror/platform_frameworks_support
- aosp-mirror/platform_development
- aosp-mirror/platform_system_core
- dart-lang/sdk
- dlang/dmd
- dlang/phobos
- mozilla/gecko-dev
- chromium/chromium
- apple/foundationdb
- mdaniel/virtualbox-org-svn-vbox-trunk
- openjdk/jdk14u
- openjdk/jdk
- openjdk/loom
- blender/blender
- Autodesk/maya-usd
- ethereum/solidity
- dotnet/runtime
- dotnet/roslyn
- dotnet/aspnetcore
- wch/r-source
- JetBrains/intellij-community
- GNOME/gimp
- apache/maven
- apache/hadoop
- magento/magento2
- erlang/otp
- elixir-lang/elixir
- ghc/ghc
- ruby/ruby
- Perl/perl5
- JetBrains/kotlin
- AdaCore/gnatstudio
- AdaCore/spark2014
- scala/scala
- apache/groovy
- docker/docker-ce
- docker/dockerhub.io
- emacs-mirror/emacs
- tcltk/tcl
- SWI-Prolog/swipl-devel
- SWI-Prolog/swipl
- dotnet/fsharp
- JuliaLang/julia
- fish-shell/fish-shell
- NixOS/nixpkgs
- jupyter/notebook
- openshift/origin
- openshift/openshift-ansible
- openshift/installer
- mysql/mysql-server
- postgres/postgres

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
