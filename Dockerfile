# Utiliza una imagen de Golang como base
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos necesarios (incluido tu código) al directorio de trabajo del contenedor
COPY . .

# Instala las dependencias de tu aplicación
RUN go mod download

# Compila tu aplicación
RUN go build -o main .

# Expone el puerto en el que se ejecutará tu aplicación
EXPOSE 8000

# Ejecuta tu aplicación cuando se inicie el contenedor
CMD ["./main"]
