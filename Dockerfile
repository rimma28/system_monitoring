FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY client/go.mod ./client/
COPY client/go.sum ./client/
COPY protobuf/go.mod ./protobuf/
COPY protobuf/go.sum ./protobuf/
COPY os/linux/go.mod ./os/linux/
COPY os/linux/go.sum ./os/linux/
RUN go mod download

COPY . ./
#RUN cd client && go build -o client && cd ..
RUN go build -o system-monitor
ENTRYPOINT [ "./system-monitor" ]
CMD []
