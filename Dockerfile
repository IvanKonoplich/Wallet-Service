FROM golang:1.18.1

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o avitoTest ./cmd/main.go

CMD ["./avitoTest"]