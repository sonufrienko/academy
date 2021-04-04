FROM golang:alpine AS build

ENV HOME=/src
WORKDIR $HOME
COPY . $HOME
RUN GOOS=linux go build -o server


# start new image for lower size
FROM alpine

ENV HOME=/app
WORKDIR $HOME
COPY --from=build /src/server $HOME

ENV USER app
ENV GROUP app

# create group and user
RUN addgroup -g 1000 -S $GROUP && \
    adduser -u 1000 -S -G $GROUP app
RUN chown -R $USER:$GROUP $HOME

USER $USER

EXPOSE 3000

ENTRYPOINT ./server