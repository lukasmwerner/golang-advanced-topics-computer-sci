FROM golang:alpine

# Install Deps
RUN apk add git
RUN go get github.com/labstack/echo
RUN go get github.com/dgrijalva/jwt-go

COPY . /src
WORKDIR /src

RUN go build .

RUN chmod +x ./src

ENTRYPOINT /src/src