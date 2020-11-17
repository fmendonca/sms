# SMS
SMS Project OpenSource GoLang and AWS

Projeto desenvolvido em GoLang com Interface HTML para disparo de SMS em massa através da AWS

### Variaveis de Ambiente

Especificar a variavel para fazer o bind na porta correta ex: 8000
```sh
export  PORT=8080
```
#### Conexão para o Banco de Dados.
#### Exemplo:
```sh
export	DBUSER="root"
export	DBPASS="mariadb"
export	DBHOST="tcp(192.168.86.190:3306)"
export	DBNAME="smsdb"
```

#### Access key e region da AWS
#### Exemplo:
```sh
export  AWSACCESSKEY="AKIA57XXXXXXXXXXX"
export  AWSSECRETKEY="vICCNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
export  AWSREGION="us-east-2"
```

Baixe as dependencias do go para que o projeto funcione.
```sh
cat requirements.txt | xargs go get
```

Para iniciar o Projeto Especifique as variáveis e depois rode o seguinte comando.
```sh
go run main.go
```

### RoadMap

- Novos Provedores
- Autenticação


License
---
MIT

***Compartilhe o Conhecimento!***
