# GoPload

GoPload é um serviço de upload de arquivos construído com Go e Gin. Suporta o upload de arquivos, download por ID e exclusão de arquivos.

## Descrição

GoPload permite o upload, download e exclusão de arquivos de maneira eficiente e segura. Utiliza tecnologias como Go, Gin e Gorm para gerenciamento de rotas e banco de dados, e Docker para facilitar a implantação em diferentes ambientes.

## Rotas

- **GET /healthcheck**: Verifica o status do serviço.
- **POST /upload**: Rota para upload de arquivos.
  - **Parâmetros**:
    - **files**: Campo obrigatório. Um ou mais arquivos para upload (multipart/form-data).
    - **path** (opcional): Define o caminho para salvar o arquivo em uma pasta específica.
- **GET /file/:idArquivo**: Rota para baixar arquivos pelo ID.
- **DELETE /file/:id**: Rota para deletar um arquivo pelo ID. Atualiza o campo `is_active` para `false` e define a data de exclusão.

## Ambientes

A aplicação dispõe de dois ambientes: desenvolvimento (dev) e produção (prod). Para iniciar a aplicação, especifique o perfil desejado no Docker Compose.

### Configuração de Ambiente

A aplicação requer um arquivo `.env` com as seguintes chaves para funcionar:

```shell
DB_CONNECTION_STRING=
BASIC_AUTH_USERNAME=
BASIC_AUTH_PASSWORD=
```

## Como usar

#### Pré-requisitos
Certifique-se de ter o Docker instalado em sua máquina. Você pode baixá-lo em Docker's website.

#### Clonando o Repositório
Clone este repositório para o seu ambiente local usando o seguinte comando:

```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
cd gopload
```
#### Configuração do Arquivo de Ambiente
Renomeie o arquivo .env.example para .env e preencha as chaves com as informações necessárias, como as configurações do banco de dados e as credenciais de autenticação básica.

#### Subindo a Aplicação
Navegue até o diretório raiz do projeto e execute o seguinte comando para subir a aplicação usando o Docker Compose. Lembre-se de especificar o perfil desejado (prod ou dev):

```bash
docker-compose --profile prod up -d
```
#### Acessando a Aplicação
Após a execução bem-sucedida do comando docker-compose up, a aplicação estará disponível em http://localhost:3000 (ou na porta especificada no arquivo .env).

## Utilizando a Aplicação
#### Upload de Arquivos:
Para fazer upload de arquivos, envie uma requisição POST para /upload contendo um ou mais arquivos no campo files (multipart/form-data). Você também pode opcionalmente fornecer um campo path para especificar o local onde o arquivo será salvo.

#### Baixar Arquivos:
Acesse a rota /file/:idArquivo para baixar arquivos pelo ID.

#### Deletar Arquivos:
Envie uma requisição DELETE para /file/:id para deletar um arquivo pelo ID. Isso atualizará o campo is_active para false e definirá a data de exclusão.

## Estrutura do Projeto

- `cmd`: Contém o arquivo principal para iniciar a aplicação.
- `internal`: Contém a lógica da aplicação, incluindo módulos, entidades e banco de dados.
- `pkg`: Contém pacotes reutilizáveis.
- `scripts`: Contém scripts úteis para desenvolvimento e implantação.

## Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nome-da-feature`)
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Faça o push para a branch (`git push origin feature/nome-da-feature`)
5. Crie um novo Pull Request

## Licença

Este projeto está licenciado sob a [`Licença MIT`](LICENSE).