FROM golang:alpine AS build

COPY . /workspace
WORKDIR /workspace
RUN go build -o /workspace/bin/helloworld --ldflags '-s -w -extldflags "-static"' .

FROM scratch

COPY --from=build /workspace/bin/helloworld /usr/bin/helloworld

ENTRYPOINT ["/usr/bin/helloworld"]
