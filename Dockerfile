# Build stage
FROM golang:alpine AS build-env
COPY . $GOPATH/src/github/quickstar/wally/
RUN apk update && apk add git
WORKDIR $GOPATH/src/github/quickstar/wally/
RUN go get && go build -o /wally

# Final stage
FROM alpine
RUN apk update && apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /wally /app/
VOLUME ["/root/Pictures/wally/"]
ENTRYPOINT ["./wally"]
