FROM golang:1.25.0-alpine

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY . .

RUN go build

CMD ./main