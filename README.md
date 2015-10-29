# gitt

git tools - tools for multiple git repos handling

- recursive `git status`
- recursive `git pull`
- ...

# build from sources

```
GOPATH = /home/work

mkdir -p $GOPATH/src

cd $GOPATH
go get github.com/constabulary/gb/...
sudo cp bin/gb /usr/bin/
sudo cp bin/gb-vendor /usr/bin/
rm -rf src/github.com

cd $GOPATH/src
git clone git@github.com:mguzelevich/gitt.git

cd $GOPATH
gb build
```

# Examples

```
$GOPATH/bin/gitt [--root=<root path>] [-h] [-v] [command] [args]
```

run `git pull` for all repositories

full:
```
$GOPATH/bin/gitt --root=. pull
```

short:
```
$GOPATH/bin/gitt pull
```

run `git status` for all repositories

full:
```
$GOPATH/bin/gitt --root=./src/bla-bla-bla status
```

short:
```
$GOPATH/bin/gitt status
```

# output legend

sample output:

```
work % bin/gitt -git status .
started in [git status] mode

=== src/environment [master] == []
=== src/github.com/mguzelevich/gitt [master] == []
  changes not staged for commit:
    * git_status.go
=== src/github.com/mguzelevich/log4go [master] == []
  changes not staged for commit:
    * examples/example.xml
=== vendor [master] == []
```

# output legend

```
=== src/github.com/mguzelevich/gitt [master] == []
  changes not staged for commit:
    * git_status.go
```

first line:

```
=== src/github.com/mguzelevich/gitt [master+1] ~~ [origin/master+2]
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~  ~~~~~~ ~      ~~~~~~~~~~~~~ ~
 repo root folder  -|                   |   |             |      |
                                        |   |             |      |
                          local branch -|   |             |      |
                                            |             |      |
                      diff - local commits -|             |      |
                                                          |      |
                                           remote branch -|      |
                                                                 |
                                          diff - remote commits -|
```

files sections (`changes to be committed`, `changes not staged for commit`, `untracked files`)

```
  changes to be committed:
    + git_status.go
    ~ ~~~~~~~~~~~~~
    |      |
    |      |- file name
    |
    |- file operation:
        `+` - added;
        `*` - modified;
        `-` - deleted;
        `>` - renamed
```

```
  untracked files:
    ? samples.txt
    ~ ~~~~~~~~~~~
    |      |
    |      |- file name
    |
    |- `?` - untracked file
```
