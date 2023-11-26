FROM golang:alpine as build
WORKDIR $GOPATH/src/github.com/jamesgawn/gohelloworld

# Copy the rest of the project and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./helloworld ./hello.go

# Reset to scratch to drop all of the above layers and only copy over the final binary
FROM scratch
ENV HOME=/home
COPY --from=build /helloworld /bin/helloworld

ENTRYPOINT ["/bin/helloworld"]

EXPOSE 8080