#
# Build stage
#

FROM golang:1.22-alpine as compiler
WORKDIR /todo
COPY . .
RUN go build -o service main.go

#
# Run stage
#
FROM alpine

COPY --from=compiler /todo /todo
WORKDIR /todo
CMD ./service
EXPOSE 8080