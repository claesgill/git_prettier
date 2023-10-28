# Git prettier
This is a git prettifyer tool that formats the output of `git status` to a short format. The purpose of the tool is to use it in your `PS1` variable so that you allways are up to date with your git repo.

## Content
1. [Install](#install)
1. [Usage](#usage)
    1. [Examples](#examples)

## Install
To get access to the `gp`-binary globally, you need to copy it into the bin folder:  
```sh
cp bin/git_prettier /usr/local/bin/
```

To verify that it works:  
```sh
gp -v
> Git Prettier version v0.0.1
```

## Usage
`gp` serves a couple of flags you may want to use. See the list underneath.

| flag | type | default | description |
| :--- | :--- | :------ | :---------- |
| -v   | bool | false   | get the version |
| -status   | bool | false   | formatting git status if any |
| -c   | bool | false   | color the status output |
| -b   | bool | false   | get active branch |
| -m   | bool | false   | get number of tracked modified files |
| -d   | bool | false   | get number of tracked deleted files |
| -u   | bool | false   | get number of untracked files |

### Examples
#### `-status` 
This flag yield the git status on a pre-defined format and will only print values if in a git folder.  
```sh
gp -status
>[master | +1 -2 U…]
```
You can also use the color flag `-c` in combination with `-status`.

#### `-b` `-m` `-d` `-u`
```sh
gp -b
>master
gp -m
>+1
````
You can use flags in combination, but they will each print on a separate line.
```sh
gp -b -m -d -u
>master
>+1
>-1
>U…
```

### PS1
> :information_source: This is an example of how you can set `PS1`. Feel free to do whatever you want.

To use `gp` inside you `PS1`varible, edit your `.bash_profile` or `.bashrc` and put the following in the bottom:  
```sh
# \u = username
# \W = folder
PS1="\u in \W \[\$(gp -status)\] "
```
Your terminal should look something like this (if in a git repo):  
```sh
Claes in git_prettier [master | +1 U…] 
```
