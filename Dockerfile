# Usar a imagem base do Go com versão desejada
FROM golang:latest

# Definir diretório de trabalho dentro do container
WORKDIR /app

# Definir mirror alternativo para o Go Modules proxy
ENV GOPROXY=https://goproxy.io,direct

# Instalar protoc (Protocol Buffers Compiler) e outras dependências
RUN apt-get update && apt-get install -y protobuf-compiler

# Instalar os plugins protoc-gen-go e protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Adicionar o diretório Go bin ao PATH
ENV PATH=$PATH:/root/go/bin

# Comando de entrada padrão (pode ser alterado conforme o uso)
CMD ["go", "version"]
