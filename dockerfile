FROM golang:alpine as build
WORKDIR /app
COPY . /app
RUN go build -o gb-api .
FROM alpine
WORKDIR /app
COPY --from=build /app/gb-api .
EXPOSE 1323
CMD ["gb-api"]
