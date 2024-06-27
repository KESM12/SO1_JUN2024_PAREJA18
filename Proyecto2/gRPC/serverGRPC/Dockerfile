# Etapa de construcción
FROM golang:alpine AS builder

# Establece el directorio de trabajo en /go/src/app
WORKDIR /go/src/app

# Copia el resto de la aplicación al directorio de trabajo
COPY . .

# Instala dependencias y compila la aplicación
RUN go get -d -v ./...
RUN go build -o /go/bin/app

# Etapa de producción
FROM alpine:latest

# Instala certificados CA para permitir conexiones HTTPS
RUN apk --no-cache add ca-certificates

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el binario de la aplicación desde la etapa de construcción
COPY --from=builder /go/bin/app /app

# Expone el puerto 3001 en el contenedor
EXPOSE 3001

# Define el comando por defecto para ejecutar la aplicación
ENTRYPOINT [ "/app/app" ]


# # Usa una imagen base de Go
# FROM golang:alpine AS builder


# # Establece el directorio de trabajo en /app
# WORKDIR /go/src/app


# # Copia el resto de la aplicación al directorio de trabajo
# COPY . .

# # Compila la aplicación
# RUN go get -d -v
# RUN go build -o /go/bin/app

# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# WORKDIR /app
# COPY --from=builder /go/bin/app /app

# ENTRYPOINT [ "/app/app" ]

# # Expone el puerto 3001
# EXPOSE 3001

