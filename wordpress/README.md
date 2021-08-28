
create your own docker volume mapped with the host directory before we mention in the docker-compose.yml as external

1.Create volume named share

```
docker volume create --driver local \
    --opt type=none \
    --opt device=/u/docker/share \
    --opt o=bind share
```

Windows:

```
docker volume create --driver local  --opt type=none   --opt device=C:\\dock
er\\share  --opt o=bind share
```

Use it in your docker-compose

```
version: "3"

volumes:
  share:
    external: true

services:
  workstation:
    container_name: "workstation"
    image: "ubuntu"
    stdin_open: true
    tty: true
    volumes:
      - share:/share:consistent
      - ./source:/source:consistent
    working_dir: /source
    ipc: host
    privileged: true
    shm_size: '2gb'
  db:
    container_name: "db"
    image: "ubuntu"
    stdin_open: true
    tty: true
    volumes:
      - share:/share:consistent
    working_dir: /source
    ipc: host
```

This way we can share the same directory with many services running in different containers




 # create a reusable volume
  $ docker volume create --driver local \
      --opt type=nfs \
      --opt o=nfsvers=4,addr=nfs.example.com,rw \
      --opt device=:/path/to/dir \
      foo

  # or from the docker run command
  $ docker run -it --rm \
    --mount type=volume,dst=/container/path,volume-driver=local,volume-opt=type=nfs,\"volume-opt=o=nfsvers=4,addr=nfs.example.com\",volume-opt=device=:/host/path \
    foo

  # or to create a service
  $ docker service create \
    --mount type=volume,dst=/container/path,volume-driver=local,volume-opt=type=nfs,\"volume-opt=o=nfsvers=4,addr=nfs.example.com\",volume-opt=device=:/host/path \
    foo

  # inside a docker-compose file
  ...
  volumes:
    nfs-data:
      driver: local
      driver_opts:
        type: nfs
        o: nfsvers=4,addr=nfs.example.com,rw
        device: ":/path/to/dir"




#!/bin/sh

# Download the file if it doesn't exist
if [ ! -f lib/huge_file.mmdb ]; then
  curl https://test.com/huge_file.mmdb --output  lib/huge_file.mmdb
fi

# Switch to the container command
exec "$@"






To solve this, you can create your own network:

docker network create dbnet
docker network connect dbnet mysql

Note, this assumes you have a container named mysql up and running.


Then configure your docker-compose.yml with:

version: '2'

networks:
  dbnet:
    external:
      name: dbnet

services:
    wordpress:
        image: wordpress
        ports:
            - 80:80
        environment:
            WORDPRESS_DB_USER: root
            WORDPRESS_DB_PASSWORD: password
        volumes:
            - /var/www/somesite.com:/var/www/html
        networks:
          - dbnet





Also note that 1.12 added the ability to change the escape character
from the backslash that Linux/Unix uses to a backtick. This is defined
in a special escape directive at the top of your Dockerfile, prefixed
with a #. I wouldn't recommend doing this with images you're building
for Linux hosts:

```
# escape=`

# Example Dockerfile with escape character changed

FROM windowsservercore
COPY testfile.txt c:\
RUN dir c:\
```

\\wsl$\Ubuntu\mnt\wsl
C:\Users\plh\AppData\Local\Docker




1) Place your folder under drive "C"

2) Open "Settings" in Docker Desktop -> "Shared Drives" -> "Reset Credentials" -> select drive "C" -> "Apply"

3) Open terminal and run (as proposed by Docker Desktop):
docker run --rm -v c:/Users:/data alpine ls /data

4) Open your docker-compose.yml and update path in -volumes:

volumes:
  - /data/YOUR_USERNAME/projects/my_project/jssecacerts:/usr/lib/jvm/java-1.8-openjdk/jre/lib/security/jssecacerts/
5) restart docker container

Share
Improve this answer
Follow
answered Apr 19 '19 at 15:29

Leonid Dashko
2,72911 gold badge1414 silver badges2121 bronze badges
Add a comment

3

Use:

volumes:
  - "C:/Users/Joey/Desktop/backend:/var/www/html"
Putting the whole thing in double quotes and using forward slashes worked for me. I was on windows 10 in windows 10 using Linux containers through WSL2




Look in:

/mnt/wsl/docker-desktop-data/tarcache/entries/docker.tar/b0963a124968e3a42af33f861a6ef36c5d80cb03d6d3348f8e0fc0c6f172b457/containers/services/docker/rootfs/mnt/c/docker/share


/mnt/wsl/docker-desktop-data/tarcache/entries/docker.tar/b0963a124968e3a42af33f861a6ef36c5d80cb03d6d3348f8e0fc0c6f172b457/containers/services/docker/rootfs/mnt/c/docker/share# ls -l
total 16
-rw-r--r-- 1 root root 29 Aug 28 12:12 New
-rw-r--r-- 1 root root 22 Aug 28 12:07 Testing
-rw-r--r-- 1 root root 12 Aug 28 11:47 XXXX
-rw-r--r-- 1 root root 29 Aug 28 12:18 Xyzzy


root@Hades:/mnt/wsl/docker-desktop-data/tarcache/entries/docker.tar/b0963a124968e3a42af33f861a6ef36c5d80cb03d6d3348f8e0fc0c6f172b457/containers/services/docker/rootfs/mnt/c/docker/share#

