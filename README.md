# Para gerar os protobufers
docker build -t protobuf --network=host .

docker compose up


# Para gerar os clients de lambdas
Adicionar a entidade no main (cmd/main/main.go)
e executar o arquivo, logo depois mover os arquivos
generated para pasta respectiva