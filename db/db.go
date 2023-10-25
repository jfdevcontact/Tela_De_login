package db

import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"fmt"

)


type Usuarios struct{
	ID string
	Email string
	Nome string
}

func DB(){

	
	U := Usuarios{
		ID: uuid.New().String(),
		//Email: "joao@gmail.com",
		//Nome: "Joao",

	}

	Conexao_DB, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/usuarios")
	if err != nil{
		log.Println("Erro ao se conectar com o banco", err)
	}

	//log.Println("Conexão bem feita com o banco")

	defer Conexao_DB.Close()


	

	Inserir_dados_no_banco, err := Conexao_DB.Prepare("insert into usuarios (id, nome, email) values(?,?,?)")
	
	_, err = Inserir_dados_no_banco.Exec(U.ID, U.Email, U.Nome)
	if err != nil{
		log.Println("Erro ao inserir os dados no bano de dados", err)
	}

	//log.Println("Dados Inserindo no banco")
	defer Inserir_dados_no_banco.Close()

	



	 
	var email string
	var username string 

	fmt.Printf("Digite seu  email\n")
	fmt.Scan(&email)

	fmt.Printf("Digite seu usuario\n")
	fmt.Scan(&username)


	


	var count1 int
	err = Conexao_DB.QueryRow("SELECT count(*) FROM usuarios WHERE email = ? and nome = ?", email, username).Scan(&count1)
	if err != nil {
		log.Fatal(err)
	}

	
	switch {
	case count1 > 0:
		fmt.Printf("Seja bem vindo %s essa são suas infomações que tava no banco\n", username)
		fmt.Printf("Seu EMAIL %s, seu USUARIO %s\n", email, username)
		break
	default:
		fmt.Printf("Suas infomaões não esta no banco\n") 
			break
	}

	

}	








//create table usuarios (id varchar(255), email varchar(80), nome varchar(80), primary key (id));