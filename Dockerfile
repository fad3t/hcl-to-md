FROM ubuntu:latest
WORKDIR /app
COPY . .
RUN go build .
CMD ["/app/hcl-to-md"]
