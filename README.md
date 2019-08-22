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
MAX 589 TOTAL 6607
  0  272 ****************************
  1  188 ********************
  2  121 *************
  3   17 **
  4    6 *
  5    5 *
  6    1 *
  7   14 **
  8   59 *******
  9  138 ***************
 10  380 ***************************************
 11  550 *********************************************************
 12  589 ************************************************************
 13  459 ***********************************************
 14  532 *******************************************************
 15  518 *****************************************************
 16  527 ******************************************************
 17  561 **********************************************************
 18  380 ***************************************
 19  246 **************************
 20  237 *************************
 21  163 *****************
 22  244 *************************
 23  400 *****************************************
```

## Notes

- `-author` flag is optional and accepts regex (same as git), i.e. `-author="Name.*"`
- Allows for many paths, passing the same flag, i.e. `-dir=path1 -dir=path2 ...`

## Motivation

Other tools such as [knadh/git-bars](https://github.com/knadh/git-bars) don't support multiple directories
