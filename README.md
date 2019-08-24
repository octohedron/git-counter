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

A folder with the biggest repos on github

+ ansible
+ azure-docs
+ DefinitelyTyped
+ git
+ homebrew-core
+ kubernetes
+ linux
+ nixpkgs
+ symfony
+ tensorflow
+ vscode

```
MAX 173651 TOTAL 2065043
  0   50031 ************************
  1   37851 ******************
  2   27721 *************
  3   20679 **********
  4   17839 *********
  5   17821 *********
  6   22586 ***********
  7   35476 *****************
  8   65494 *******************************
  9  108253 **************************************************
 10  140042 *****************************************************************
 11  157182 *************************************************************************
 12  134611 ***************************************************************
 13  143373 *******************************************************************
 14  166292 *****************************************************************************
 15  173651 ********************************************************************************
 16  165089 *****************************************************************************
 17  129491 ************************************************************
 18   92490 *******************************************
 19   76414 ************************************
 20   70725 *********************************
 21   74684 ***********************************
 22   73787 **********************************
 23   63461 ******************************
2019/08/24 17:49:28 6.771680574s
```

## Notes

- `-author` flag is optional and accepts regex (same as git), i.e. `-author="Name.*"`
- Allows for many paths, passing the same flag, i.e. `-dir=path1 -dir=path2 ...`
- Only supports folders (i.e. `-dir=path` flag) that have folders with `.git` directories on them

## Motivation

Other tools such as [knadh/git-bars](https://github.com/knadh/git-bars) don't support multiple directories
