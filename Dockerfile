FROM golang:latest
WORKDIR /go/src/Go-employee-service
COPY . .

RUN go install -mod vendor
RUN go build
CMD ["./EnployeeService"]