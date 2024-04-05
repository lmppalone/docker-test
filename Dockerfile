# base image
FROM golang:1.22.2-alpine3.19

# ENV
ENV APP_PORT=8000

# working directory for container
WORKDIR /go/src/app

# copy everything to app
# COPY source dest
# source: current location on host machine
# dest: location in docker container
COPY . .

# download all dependencies listed in go.mod
RUN go get -d -v

# build the executable
RUN go build -v -o app main.go

# run the executable
CMD [ "sh", "-c", "./app -p $APP_PORT" ]