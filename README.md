
Some first tests


```zsh
grep --color=auto --exclude-dir={.bzr,CVS,.git,.hg,.svn} -Hnri "action=" \* 6,84s user 0,40s system 84% cpu 8,535 total
```

```zsh
cgrep scan -p ./ -e "action=" -t 1 2,60s user 1,03s system 70% cpu 5,168 total

cgrep scan -p ./ -e "action=" -t 10 1,59s user 1,09s system 61% cpu 4,369 total

cgrep scan -p ./ -e "action=" -t 100 1,36s user 1,06s system 58% cpu 4,144 total

cgrep scan -p ./ -e "action=" -t 1000 0,70s user 0,58s system 33% cpu 3,809 total

cgrep scan -p ./ -e "action=" -t 2000 0,44s user 0,17s system 242% cpu 0,251 total

```