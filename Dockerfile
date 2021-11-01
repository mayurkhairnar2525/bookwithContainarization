FROM golang:latest
LABEL maintainer ="Quique<mayurkhairnar325@gmail.com"
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY . .
ENV MYSQL_ROOT_PASSWORD: "12345678"
RUN go build
CMD ["./containerization"]



