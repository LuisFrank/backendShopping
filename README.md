The most straightforward way to use this image is to use a Go container as both the build and runtime environment. In your Dockerfile, writing something along the lines of the following will compile and run your project:

FROM golang:1.17

WORKDIR /go/src/app
COPY . .


CMD ["app"]

You can then build and run the Docker image:
 Usar estos para crear contenedor y compilar
$ docker build -t my-golang-app .
$ docker run -it --rm --name my-running-app my-golang-app

Compile your app inside the Docker container
There may be occasions where it is not appropriate to run your app inside a container. To compile, but not run your app inside the Docker instance, you can write something like:

$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.17 go build -v
This will add your current directory as a volume to the container, set the working directory to the volume, and run the command go build which will tell go to compile the project in the working directory and output the executable to myapp. Alternatively, if you have a Makefile, you can run the make command inside your container.

$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.17 make



## https://blog.friendsofgo.tech/posts/dockerizando-tu-aplicacion-en-go/
## https://hub.docker.com/_/golang

