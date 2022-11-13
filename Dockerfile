FROM golang:alpine as build
WORKDIR /go/src/golang-united-homework  
COPY . .
RUN go mod download
RUN go build -o ./server ./cmd/main.go

FROM alpine
WORKDIR .
COPY --from=build /go/src/golang-united-homework/server ./server
CMD ["./server"]