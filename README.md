Projeto simples de um servidor web que aceita apenas requisições get na porta 8080

subir com go ruan main.go

para o endpoint teste-requisição

curl --location --request GET 'http://localhost:8080/calculadora'

para o endpoint calculadora

curl --location --request GET 'http://localhost:8080/calculadora' \
--header 'Content-Type: application/json' \
--data '{
    "val1": 4,
    "val2": 5,
    "operacao": "+"
    }'
