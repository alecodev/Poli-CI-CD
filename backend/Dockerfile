# Utiliza una imagen base de Go
FROM golang:1.22.3

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de tu backend al contenedor
COPY . .

# Compila tu aplicación
RUN go build -o my-backend main.go

# Expone el puerto en el que tu backend escucha
EXPOSE 8080

# Define el comando para ejecutar tu aplicación
CMD ["./my-backend"]