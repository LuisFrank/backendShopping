# FROM golang:1.17

# WORKDIR /go/src/app
# COPY . .

# # RUN go get -d -v ./...
# # RUN go install -v ./...

# # CMD ["app"]
# CMD ["go","run","main.go"]


## https://blog.friendsofgo.tech/posts/dockerizando-tu-aplicacion-en-go/
FROM golang:1.17 AS build
WORKDIR /go/src/myapp
COPY . .
RUN go build -o /go/bin/myapp main.go

FROM scratch
COPY --from=build /go/bin/myapp /go/bin/myapp
ENTRYPOINT ["/go/bin/myapp"]