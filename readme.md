# Nome do Projeto

Um breve resumo sobre o que é o projeto.

## Descrição

Descrição mais detalhada do projeto, incluindo seu propósito, funcionalidades e tecnologias utilizadas.

## Rotas

- **GET /healthcheck**: Verifica o status do serviço.
- **POST /upload**: Rota para upload de arquivos.

  - **Parâmetros**:
    - **files**: Campo obrigatório. Um ou mais arquivos para upload (multipart/form-data).
    - **path** (opcional): Define o caminho para salvar o arquivo em uma pasta específica.
- **GET /file/:idArquivo**: Rota para baixar arquivos pelo ID.

## Ambientes

A aplicação dispõe de dois ambientes: desenvolvimento (dev) e produção (prod). Para iniciar a aplicação, especifique o perfil desejado no Docker Compose.

### Configuração de Ambiente

A aplicação requer um arquivo `.env` com as seguintes chaves para funcionar:

```shell
DB_CONNECTION_STRING=
BASIC_AUTH_USERNAME=
BASIC_AUTH_PASSWORD=
```

Dentro do repositório, há um arquivo `.env.example` que pode ser renomeado para `.env` e preenchido com as informações necessárias.

## Como Usar

- Pré-requisitos:
Certifique-se de ter o Docker instalado em sua máquina. Você pode baixá-lo em Docker's website.
- Clonando o Repositório:
Clone este repositório para o seu ambiente local usando o seguinte comando:
```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
```
- Configuração do Arquivo de Ambiente:
Renomeie o arquivo .env.example para .env e preencha as chaves com as informações necessárias, como as configurações do banco de dados e as credenciais de autenticação básica.
- Subindo a Aplicação:
Lembrando que a aplicação tem dois ambientes prod e dev para mudar qual ambiente quer subir é só alterar na flag --profile onde atualmente está prod. Dito isso navegue até o diretório raiz do projeto e execute o seguinte comando para subir a aplicação usando o Docker Compose:
```bash
 docker-compose --profile prod up -d
```
- Acessando a Aplicação:
Após a execução bem-sucedida do comando docker-compose up, a aplicação estará disponível em http://localhost:3000 (ou na porta especificada no arquivo .env).
- Utilizando a Aplicação:
Para fazer upload de arquivos, envie uma requisição POST para /upload contendo um ou mais arquivos no campo files (multipart/form-data). Você também pode opcionalmente fornecer um campo path para especificar o local onde o arquivo será salvo.
- Outras Operações:
A aplicação também oferece uma rota /file/:idArquivo para baixar arquivos pelo ID

## Contribuição

Se deseja contribuir com este projeto, por favor siga as instruções do arquivo CONTRIBUTING.md.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).