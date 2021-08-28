
# XXX = `pwd`
SRC = ${HOME}/src/Docker/docker-playpen/examples/using_alpine/share

build:
	docker build -t my-alpine -f Dockerfile.alpine .

run:
	# docker run -it --rm -d --name alpine my-alpine
	ls -l $(SRC)
	docker run -it --rm --name alpine -d -v $(SRC):/share my-alpine

