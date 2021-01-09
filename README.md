# about 

cgrep is a grep variant, written in golang. It will use more cpu than normal grep, but therefor it will be faster. Useful, when to scan directories recursive and grep each file content.
But it also can read from stdin.

# install

```zsh
go install
```

Recommended: set an alias.

## usage

## list options

```zsh
cgrep
```


## reading from stdin 

For example, portscan and cgrep:

```zsh
portscan scan -i 127.0.0.1 -t 20 | cgrep ^open
```
Using more than one thread:

```zsh
portscan scan -i 127.0.0.1 -t 20 | cgrep --threads=20 ^open
```

## scan directory recursive and grep file content

```zsh
cgrep --path ./ "^open"
```

Using more threads

```zsh
cgrep --path ./ --threads 10 "^open"
```

## save and load expressions which you often use

save an expression:

```zsh
cgrep -save -name formAction "form(.?)action="
```

This will store a named expression "formAction" at /{your-homedir}/cgrep/templates.json

To use it, you can use the option and combine it with the other stuff like for example path grepping:

```zsh
cgrep -expression formAction -p ./
```

