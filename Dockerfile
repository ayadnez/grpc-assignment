FROM golang:1.22.5

WORKDIR /app

COPY go.mod ./
COPY go.sum ./


COPY . ./ 

RUN go build -o grpc-server . && go build -o grpc-client ./client

# expose ports for grpc-server and client

EXPOSE 3000
EXPOSE 4001

# ENV variable to switch between server and client 

ENV SERVICE_TYPE=server



CMD ["/bin/sh", "-c", "if [ \"$SERVICE_TYPE\" = \"server\" ]; then ./grpc-server; else ./grpc-client; fi"]