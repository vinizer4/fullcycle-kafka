FROM golang:1.16-buster
# Definir o diretório de trabalho
WORKDIR /go/src
# Atualizar pacotes e instalar dependências
RUN apt-get update && apt-get install -y build-essential
# Copiar o código-fonte da aplicação
COPY . .
# Baixar as dependências Go
RUN go mod tidy
# Rodar o comando padrão
CMD ["tail", "-f", "/dev/null"]