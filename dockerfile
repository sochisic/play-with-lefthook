FROM golang:1.13 AS build

WORKDIR /go/src/play
ADD . .
RUN go get -d ./... 
RUN go install -v ./...
COPY . /go/src/play
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/play

# final stage
FROM golang:1.13-alpine
WORKDIR /play-with-lefthook
COPY --from=build /bin/play /bin/play
EXPOSE 80
EXPOSE 443
ENTRYPOINT ["/bin/play"]