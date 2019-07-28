# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o app

# final stage
FROM alpine
WORKDIR /play-with-lefthook
COPY --from=build-env /src/app /play-with-lefthook/
ENTRYPOINT ./app