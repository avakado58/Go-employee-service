FROM golang:latest

COPY ./ ./
RUN go install -mod vendor
RUN go build -o go-employee
CMD ./go-employee