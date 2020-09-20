# Git prettier
This is a git prettifyer tool that formats the output of `git status` to a short format. The purpose of the tool is to use it in your `PS1` variable so that you allways are up to date with your git repo.

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
`gp` only serves two flags `-v` and `-status`. The latter yield the git status on a pre-defined format and will only print values if in a git folder.  

```sh
gp -status
>[master | +1 -2 U…]
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
