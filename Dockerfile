FROM golang:latest 
COPY . /go/src/github.com/gomaglev/microshop
WORKDIR /go/src/github.com/gomaglev/microshop
RUN go get ./cmd/microshop
RUN go build -ldflags "-w -s" -o ./cmd/microshop/microshop ./cmd/microshop
RUN touch .env
CMD ["microshop", "start"]