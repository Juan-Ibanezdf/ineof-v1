📂 Estrutura de Pastas
/api
 ├── /docs                 # Documentação gerada pelo Swagger
 ├── /sqlc                 # Arquivos gerados pelo SQLC (queries e structs)
 ├── /cmd                  # Pontos de entrada da aplicação
 │    ├── server.go        # Arquivo de inicialização do servidor
 ├── /config               # Configurações gerais do sistema (env, conexões, etc.)
 │    ├── /db              # Configurações do banco de dados usando o migration
 |    ├──├──/migration 
 |    ├──├──/query 
 |    ├──├──/sqlc 
 ├── /internal             # Código interno da aplicação
 │    ├── /models          # Estruturas de dados (entidades do banco)
 │    ├── /repository      # Repositórios de acesso ao banco (queries SQLC)
 │    ├── /services        # Lógica de negócio (chama os repositories)
 │    ├── /handlers        # Handlers HTTP (chama os services)
 │    ├── /middleware      # Middlewares (autenticação, logs, etc.)
 │    ├── /routes          # Definição das rotas
 ├── /pkg                  # Pacotes reutilizáveis
 ├── /tests                # Testes unitários e de integração
 ├── main.go               # Entrada principal da aplicação
 ├── go.mod                # Gerenciador de dependências
 ├── go.sum                # Hash das dependências
 ├── Dockerfile            # Arquivo para containerização
 ├── .env                  # Arquivo de variáveis de ambiente
 ├── .gitignore            # Arquivo para ignorar arquivos no Git
 ├── .Makefile             # Arquivo de configurações para Makefile 
 ├── .sqlc.yaml            # Arquivo para configurações do sqlc 
 ├── README.md             # Arquivo de README.md


📂 api/

    Contém a documentação gerada pelo Swagger.
    Arquivos OpenAPI ou JSON gerado pelo Swagger.


📂 sqlc/

    Contém queries SQL geradas pelo SQLC e os arquivos .json com a configuração do SQLC.

📂 cmd/

    Contém o ponto de entrada da aplicação. Se houver múltiplas execuções (ex: API, Worker), você pode criar:
        /cmd/api/main.go
        /cmd/worker/main.go

📂 config/

    Gerencia variáveis de ambiente e configurações da aplicação.

📂 internal/

    Código interno, não exportado para fora do projeto.

📂 internal/models/

    Estruturas das tabelas do banco geradas pelo SQLC.

📂 internal/repository/

    Implementa repositórios para acessar o banco de dados via SQLC.