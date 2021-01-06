# install

```zsh
go install
```

Recommended: set an alias.

## usage

### reading from stdin 

For example, portscan and cgrep:

```zsh
portscan scan -i 127.0.0.1 -t 20 | cgrep ^open
```
Using more than one thread:

```zsh
portscan scan -i 127.0.0.1 -t 20 | cgrep --threads=20 ^open
```

### scan directory recursive and grep file content

```zsh
cgrep --path ./ "^open"
```

Using more threads

```zsh
cgrep --path ./ --threads 10 "^open"
```
