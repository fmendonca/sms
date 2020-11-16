package models

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/fmendonca/sms/db"
)

type EstrutDBSMS struct {
	Id          int
	Nome        string
	NTelefone   string
	Status      string
	SMSMensagem string
}

func BuscaItem() []EstrutDBSMS {
	db := db.Conexaodb()

	allcad, err := db.Query("select * from tb_ntelefone")
	if err != nil {
		panic(err.Error())
	}

	d := EstrutDBSMS{}

	estructdbsms := []EstrutDBSMS{}

	for allcad.Next() {
		var id int
		var nome string
		var telefone string
		var status string
		log.SetOutput(os.Stdout)

		err = allcad.Scan(&id, &nome, &telefone, &status)
		if err != nil {
			panic(err.Error())
		}

		d.Id = id
		d.Nome = nome
		d.NTelefone = telefone
		d.Status = status

		estructdbsms = append(estructdbsms, d)
	}
	defer db.Close()
	return estructdbsms
}

func CriarItem(nome, telefone, status string) {
	db := db.Conexaodb()
	log.SetOutput(os.Stdout)
	insertNoBanco, err := db.Prepare("insert into tb_ntelefone(nome, telefone, status) values(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insertNoBanco.Exec(nome, telefone, status)
	defer db.Close()
}

func DeletaItem(id string) {

	db := db.Conexaodb()
	deletarItem, err := db.Prepare("DELETE FROM tb_ntelefone WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	deletarItem.Exec(id)
	log.Println("DELETE")
	defer db.Close()
}

func EditItem(id string) EstrutDBSMS {
	db := db.Conexaodb()
	itemDoBanco, err := db.Query("select * from tb_ntelefone where id=?", id)
	if err != nil {
		panic(err.Error())
	}

	editarItem := EstrutDBSMS{}

	for itemDoBanco.Next() {
		var id int
		var nome string
		var telefone string
		var status string

		err = itemDoBanco.Scan(&id, &nome, &telefone, &status)
		if err != nil {
			panic(err.Error())
		}
		editarItem.Id = id
		editarItem.Nome = nome
		editarItem.NTelefone = telefone
		editarItem.Status = status
	}

	defer db.Close()
	return editarItem

}

func AtualizarItem(id int, nome, telefone, status string) {
	//funcao para atualizar os itens no banco
	db := db.Conexaodb()

	atualizaItem, err := db.Prepare("update tb_ntelefone set nome=?, telefone=?, status=? where id=?")
	if err != nil {
		panic(err.Error())
	}

	atualizaItem.Exec(nome, telefone, status, id)
	defer db.Close()

}

func CadastrarMSG(mensagem string) {

	structsms := []string{}

	db := db.Conexaodb()
	//insertNoBanco, err := db.Prepare("insert into tb_sms(smsmensagem) values(?)")
	insertNoBanco, err := db.Prepare("update tb_sms set smsmensagem=? where id=1")
	if err != nil {
		panic(err.Error())
	}
	insertNoBanco.Exec(mensagem)

	ntel, err := db.Query("select telefone from tb_ntelefone where status='Habilitado'")
	if err != nil {
		panic(err.Error())
	}

	for ntel.Next() {
		var telefone string
		err = ntel.Scan(&telefone)
		if err != nil {
			panic(err.Error())
		}

		structsms = append(structsms, telefone)
	}

	for _, value := range structsms {
		err := SendSMS(value, mensagem)
		//fmt.Println(value, mensagem)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer db.Close()
}

func SendSMS(phoneNumber string, message string) error {

	AccessKeyID := os.Getenv("AWSACCESSKEY")
	SecretAccessKey := os.Getenv("AWSSECRETKEY")
	AwsRegion := os.Getenv("AWSREGION")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AwsRegion),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
	},
	)

	// Create SNS service
	svc := sns.New(sess)

	// Pass the phone number and message.
	params := &sns.PublishInput{
		PhoneNumber: aws.String(phoneNumber),
		Message:     aws.String(message),
	}

	// sends a text message (SMS message) directly to a phone number.
	resp, err := svc.Publish(params)

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(resp) // print the response data.

	return nil
}
