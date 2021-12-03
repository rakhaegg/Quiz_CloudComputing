
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY model ./model

RUN go mod download

COPY *.go ./

RUN go build -o /session

EXPOSE 8081

CMD [ "/session" ]