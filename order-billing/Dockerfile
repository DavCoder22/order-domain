# Usar una imagen base oficial de Go
FROM golang:1.23.6-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo go.mod y go.sum
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el resto de la aplicación
COPY . .

# Construir la aplicación
RUN go build -o order-billing ./cmd/order-billing

# Exponer el puerto en el que la aplicación correrá
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./order-billing"]
