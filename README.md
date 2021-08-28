# Docker Playpen

Check in the bin folder for some scripts to run and connect to alpine and ubuntu images. I add this to the end of my PATH with something like.

```
export PATH=$PATH:${SRC}/docker/docker_playpen/bin
```
where SRC points to the root of my source tree.

Here are the SHELL aliases I use for docker:

```
# ---- docker --------------------------------------------------------

# export TERM=cygwin

alias do=docker
alias d-c=docker-compose
alias dps='docker ps'
alias did='docker images --digests'
alias dil='docker image ls'
alias dila='docker image ls -a'
alias dip='docker image prune'
alias dcl='docker container ls'
alias dcla='docker container ls-a '
alias dcp='docker container prune'
```


# Docker Getting STarted

```
docker run -d -p 80:80 docker/getting-started
```


