# SMS
SMS Project OpenSource GoLang and AWS

Projeto desenvolvido em GoLang com Interface HTML para disparo de SMS em massa através da AWS

### Variaveis de Ambiente

#### especificar a variavel para fazer o bind na porta correta ex: 8000
```sh
export  PORT=8080
```

# conexão para o Banco de Dados.
# Exemplo:
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

#### Dependencias GoLang
```sh
cat requirements.txt | xargs go get
```

#### RoadMap

- Novos Provedores
- Autenticação
- Disparo de e-mail em Massa
- Disparo pelo WhatsApp


100% OpenSource.
Compartilhe o Conhecimento!