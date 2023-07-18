## 🚀 Começando:

Essas instruções permitirão que você obtenha uma cópia e consiga rodar o projeto em sua máquina local.

  ####  🖨️ Realize o clone do projeto:
    git clone https://github.com/LeonardoGregoriocs/transaction_routine.git

### 📋 Pré-requisitos & Instalação:

   Instalar o [Docker](https://docs.docker.com/get-docker)
   
   Instalar o [Golang](https://go.dev/doc/install)
   
   Instalar o [Postman](https://www.postman.com/downloads/) ou outra plataforma para realizar as operações da API. 

### ⬆️ Subindo o Banco de dados:

Após realizar a instalação do docker, basta abrir um terminal dentro da pasta em que colou o projeto e rodar o comando abaixo:

    docker-compose up -d --build

O comando acima irá construir e subir 2 containeres do docker: 

  1 - O primeiro contendo o PostgreSQL.
  
  2 - O segundo irá subir uma interface amigável para você acessar ao banco de dados, se achar melhor, pode usar um SGBD de sua preferência. 

Logo após, é somente acessar o seguinte endereço em seu navegador: 
    
    localhost:7000

Para acessar o banco de dados, basta colocar as credenciais abaixo: 

    Sistema --> PostgresSQL
    Servidor --> database
    Usuário --> user
    Senha --> user123
    Base de dados --> dataclient

![image](https://github.com/LeonardoGregoriocs/transaction_routine/assets/83976271/84d45001-bb14-4d57-9d2d-0f77efd55570)

### ⬆️ Subindo a aplicação:

Para subir a aplicação, basta abrir um terminal em sua IDE, e digitar o comando abaixo para entrar na pasta cmd: 

    cd cmd

Logo após, digitar o comando abaixo: 

    go run main.go

Pronto, sua aplicação já está no ar! 

Agora para consultar os endpoinst, basta utilizar as rotas abaixo: 

    - POST - http://localhost:8000/accounts - Criar um novo cliente. 
    - GET - http://localhost:8000/accounts/id - Buscar um cliente através do ID. 
    - POST - http://localhost:8000/transactions - Criar uma nova transação.
    
