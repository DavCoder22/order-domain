# Usa una imagen base de Go
FROM golang:1.23.6 AS builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del microservicio
COPY src/ ./src/
COPY go.mod go.sum ./

# Descarga las dependencias y construye el binario
WORKDIR /app/src
RUN go mod tidy && go build -o /app/main

# Imagen final más liviana
FROM alpine:latest
WORKDIR /root/

# Copia el binario desde la imagen anterior
COPY --from=builder /app/main .

# Verifica que el binario tiene permisos de ejecución
RUN chmod +x /root/main

# Expone el puerto que usa el servicio
EXPOSE 8080

# Comando para ejecutar el microservicio
CMD ["./main"]