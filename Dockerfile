# Etapa de construção
FROM golang:1.23-alpine AS builder

# Instala dependências necessárias para o Go
RUN apk add --no-cache git

# Define o diretório de trabalho
WORKDIR /go/src/app

# Copia os arquivos do projeto para o contêiner
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante dos arquivos do projeto
COPY . .

# Compila o binário da aplicação
RUN go build -o main cmd/main.go

# Etapa final: imagem leve
FROM alpine:latest

# Define o diretório de trabalho
WORKDIR /app

# Copia o binário da etapa de construção
COPY --from=builder /go/src/app/main .

# Expõe a porta da aplicação
EXPOSE 8080

# Executa o binário
CMD ["./main"]
