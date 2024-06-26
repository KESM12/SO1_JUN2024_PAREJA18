# Usa una imagen base de Go
FROM golang:alpine AS builder


# Establece el directorio de trabajo en /app
WORKDIR /go/src/app


# Copia el resto de la aplicación al directorio de trabajo
COPY . .

# Compila la aplicación
RUN go get -d -v
RUN go build -o /go/bin/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/bin/app /app

ENTRYPOINT [ "/app/app" ]

# Expone el puerto 3001
EXPOSE 3001

