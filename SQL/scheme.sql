CREATE DATABASE smsdb;

USE smsdb;

CREATE TABLE `tb_ntelefone` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `nome` varchar(30) NOT NULL,
  `telefone` varchar(30) NOT NULL,
  `status` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE `tb_sms` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `smsmensagem` varchar(500) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO tb_sms(smsmensagem) values('first value');
