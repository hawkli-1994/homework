FROM golang:1.17-alpine AS build

WORKDIR /app
COPY go.mod *.go ./
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /server

FROM scratch
COPY --from=build /server /server
EXPOSE 8000
ENTRYPOINT ["/server"]