FROM golang:latest

WORKDIR /src
COPY ./create-middle-name/src /src

RUN apt update -y && apt upgrade -y
RUN apt install xz-utils patch file -y


ENV HOST 0.0.0.0

CMD [ "go", "run", "main.go" ]
