FROM golang:1.22.3

WORKDIR /ascii-docker

COPY . .

EXPOSE 8050

LABEL container="dockercontainer"

CMD [ "go","run","." ]