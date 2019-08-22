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

## Example output

```
==> Executing: sh -c git --git-dir=/home/user/go/src/github.com/octohedron/.../.git log --author='Gus.*' --format='%ad' --date='format:%H'
...
MAX 621 TOTAL 6925
  0  274 ***************************
  1  188 *******************
  2  121 ************
  3   17 **
  4    6 *
  5    5 *
  6    1 *
  7   14 **
  8   60 ******
  9  145 ***************
 10  404 ****************************************
 11  592 **********************************************************
 12  621 ************************************************************
 13  491 ************************************************
 14  567 *******************************************************
 15  540 *****************************************************
 16  561 *******************************************************
 17  581 *********************************************************
 18  388 **************************************
 19  249 *************************
 20  243 ************************
 21  185 ******************
 22  263 **************************
 23  409 ****************************************
```

## Notes

- `-author` flag is optional and accepts regex (same as git), i.e. `-author="Name.*"`
- Allows for many paths, passing the same flag, i.e. `-dir=path1 -dir=path2 ...`
- Only supports folders (i.e. `-dir=path` flag) that have folders with `.git` directories on them

## Motivation

Other tools such as [knadh/git-bars](https://github.com/knadh/git-bars) don't support multiple directories
