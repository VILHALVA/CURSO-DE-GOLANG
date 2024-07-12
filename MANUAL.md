# MANUAL
## INSTALAÇÃO:
### INSTALANDO NO WINDOWS:
1. **Baixar o instalador:**
   - Acesse o site oficial do Go em [golang.org](https://golang.org/).
   - Clique no link "Download Go" na página inicial.
   - Escolha a versão adequada para o seu sistema operacional (Windows) e faça o download do instalador.

2. **Executar o instalador:**
   - Após o download, execute o instalador clicando duas vezes no arquivo baixado (com extensão `.msi`).
   - Siga as instruções do instalador. O instalador padrão normalmente instala o Go no diretório `C:\Go`.

3. **Configurar variáveis de ambiente:**
   - Adicione o diretório `bin` do Go ao seu PATH do sistema:
     - Abra o "Painel de Controle" -> "Sistema e Segurança" -> "Sistema" -> "Configurações avançadas do sistema".
     - Clique em "Variáveis de Ambiente".
     - Na seção "Variáveis do Sistema", encontre a variável PATH e clique em "Editar".
     - Adicione o caminho para o diretório `bin` do Go (normalmente `C:\Go\bin`) ao PATH e clique em "OK".

4. **Verificar a instalação:**
   - Abra o prompt de comando (CMD).
   - Digite `go version` e pressione Enter. Se a instalação for bem-sucedida, você verá a versão do Go instalada.

### INSTALANDO NO MACOS:
1. **Instalar via Homebrew (recomendado):**
   - Abra o Terminal.
   - Execute o seguinte comando para instalar o Go via Homebrew:
     ```
     brew install go
     ```

2. **Verificar a instalação:**
   - Após a instalação, digite `go version` no Terminal e pressione Enter. Você deverá ver a versão do Go instalada.

### INSTALANDO NO LINUX:
1. **Instalação via gerenciador de pacotes (Debian/Ubuntu):**
   - Abra o Terminal.
   - Execute o seguinte comando para instalar o Go:
     ```
     sudo apt-get update
     sudo apt-get install golang
     ```

2. **Instalação manual:**
   - Baixe a versão mais recente do Go no site oficial: [golang.org](https://golang.org/dl/).
   - Extraia o arquivo baixado para `/usr/local` ou em um diretório de sua escolha.
   - Adicione o caminho do diretório `bin` do Go ao seu PATH do sistema.

3. **Verificar a instalação:**
   - Abra o Terminal.
   - Digite `go version` e pressione Enter. Você deverá ver a versão do Go instalada.

## PROJETO: DE CRIAÇÃO DE APP ATÉ COPILAÇÃO
### CRIANDO UM PROJETO:
1. **Criar um diretório para o projeto:**
   - Abra o terminal ou prompt de comando.
   - Navegue até o diretório onde deseja criar o projeto.
   - Crie um novo diretório para o projeto, por exemplo:
     ```
     mkdir meu_projeto
     cd meu_projeto
     ```

2. **Inicializar um módulo Go:**
   - Para criar um novo módulo Go, execute o seguinte comando no terminal:
     ```
     go mod init nome_do_modulo
     ```
   - Substitua `nome_do_modulo` pelo nome que deseja dar ao seu módulo. Por exemplo:
     ```
     go mod init meu_modulo
     ```
   - Isso criará um arquivo `go.mod` no diretório do projeto, que registra as dependências do módulo. Esse arquivo será semelhante a esse:
   ```mod
   module main

   go 1.22.3
   ```

3. **Criar um arquivo Go:**
   - Dentro do diretório do projeto, crie um novo arquivo Go com a extensão `.go`, por exemplo:
     ```
     touch main.go
     ```

4. **Escrever código Go:**
   - Abra o arquivo `main.go` em um editor de texto ou IDE de sua escolha.
   - Escreva o código Go no arquivo conforme necessário para seu projeto. Exemplo:
   ```go
   package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }
   ```

   * **Este é o seu código Go. Neste código, você:**

    - Declare um mainpacote (um pacote é uma forma de agrupar funções e é composto por todos os arquivos no mesmo diretório).

    - Importe o popular fmtpacote , que contém funções para formatação de texto, incluindo impressão no console. Este pacote é um dos pacotes de biblioteca padrão que você obteve quando instalou o Go.

    - Implemente uma mainfunção para imprimir uma mensagem no console. Uma mainfunção é executada por padrão quando você executa o mainpacote.

### EXECUTANDO O PROJETO:
1. **Executar o código Go:**
   - Para executar o código Go, use o comando `go run` seguido do nome do arquivo Go que contém a função `main()`. Por exemplo:
     ```
     go run main.go
     ```

     OU:
     ```
     go run .
     ```

2. **Compilar o projeto:**
   - Se desejar compilar o projeto em um executável, use o comando `go build`. Por exemplo:
     ```
     go build
     ```
   - Isso criará um executável com o mesmo nome do diretório atual. Por exemplo, se o diretório do projeto for `meu_projeto`, o executável será `meu_projeto.exe` no Windows ou `meu_projeto` no Linux/macOS.

3. **Executar o executável compilado:**
   - Após compilar o projeto, você pode executar o executável gerado diretamente. Por exemplo:
     ```
     ./meu_projeto
     ```

     OU:
     ```
     meu_projeto
     ```

## PROJETO: SUBINDO O SERVIDOR:
O Go possui um pacote padrão chamado `net/http`, que fornece um servidor HTTP nativo. Com esse pacote, é possível criar rapidamente um servidor web para lidar com solicitações HTTP.

Aqui está um exemplo simples de como criar um servidor web básico em Go usando o pacote `net/http`:

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Olá, você acessou: %s\n", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Servidor web iniciado. Acesse http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
```

Neste exemplo:

- Definimos uma função `handler` que será chamada sempre que uma solicitação for feita para o servidor. Essa função responde com uma mensagem que inclui o caminho da solicitação (`r.URL.Path`).
- No `main()`, registramos a função `handler` para o caminho raiz ("/") usando `http.HandleFunc`.
- Em seguida, chamamos `http.ListenAndServe` para iniciar o servidor na porta 8080.

Quando você executa este programa e acessa `http://localhost:8080/` em seu navegador, verá a mensagem "Olá, você acessou: /". Este é apenas um exemplo simples, mas com o pacote `net/http`, você pode construir servidores web mais complexos e robustos.
