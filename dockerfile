# # build stage
# FROM golang:alpine AS build-env
# ADD . /src
# RUN cd /src && go build -o app

# # final stage
# FROM alpine
# WORKDIR /play-with-lefthook
# COPY --from=build-env /src/app /play-with-lefthook/
# ENTRYPOINT ./app

FROM golang AS build
# RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/bot
ADD . .
RUN go get -d ./... 
# RUN dep ensure -vendor-only
RUN go install -v ./...
COPY . /go/src/play
RUN go build -o /bin/play

# final stage
FROM golang:alpine
# RUN apk add --no-cache bash
WORKDIR /play-with-lefthook
COPY --from=build /bin/play /bin/app
ENTRYPOINT ["/bin/app"]