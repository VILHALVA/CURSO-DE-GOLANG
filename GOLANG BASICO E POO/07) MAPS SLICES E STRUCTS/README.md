# MAPS SLICES E STRUCTS
Vamos explorar Maps, Slices e Structs em Go. São estruturas de dados fundamentais que desempenham papéis importantes na linguagem. Vou explicar cada um deles e fornecer exemplos de código.

### Maps (Mapas)

Mapas em Go são estruturas de dados que armazenam pares de chave-valor, onde cada chave é única. Eles são semelhantes a dicionários ou tabelas hash em outras linguagens de programação. Para declarar e usar um mapa em Go:

```go
// Declarar um mapa (map[keyType]valueType)
meuMapa := make(map[string]int)

// Adicionar valores ao mapa
meuMapa["um"] = 1
meuMapa["dois"] = 2

// Acessar valores
valor := meuMapa["um"]

// Verificar se uma chave existe
if _, existe := meuMapa["três"]; existe {
    fmt.Println("A chave 'três' existe no mapa.")
}

// Remover uma chave
delete(meuMapa, "dois")
```

### Slices (Fatias)

Slices são estruturas de dados semelhantes a arrays, mas são mais flexíveis e dinâmicos. Eles permitem que você trabalhe com sequências de elementos. Você pode criar um slice de várias maneiras:

```go
// Criar um slice vazio com capacidade inicial
meuSlice := make([]int, 0, 10)

// Criar um slice usando uma parte de um array existente
meuArray := [5]int{1, 2, 3, 4, 5}
meuSlice := meuArray[1:4] // Contém [2, 3, 4]

// Criar um slice literal
meuSlice := []int{1, 2, 3, 4, 5}

// Adicionar elementos a um slice
meuSlice = append(meuSlice, 6)

// Acessar elementos de um slice
valor := meuSlice[2]

// Tamanho e capacidade de um slice
tamanho := len(meuSlice)
capacidade := cap(meuSlice)
```

### Structs (Estruturas)

Structs são tipos de dados personalizados que permitem que você agrupe diferentes campos de dados sob uma única estrutura. Eles são úteis para modelar objetos complexos ou representar dados com várias propriedades. Aqui está como você pode declarar e usar uma struct em Go:

```go
// Definir uma struct
type Pessoa struct {
    Nome  string
    Idade int
}

// Criar uma instância da struct
pessoa1 := Pessoa{"Alice", 30}
pessoa2 := Pessoa{Nome: "Bob", Idade: 25}

// Acessar campos da struct
nome := pessoa1.Nome
idade := pessoa2.Idade

// Modificar campos da struct
pessoa1.Idade = 31
```

Estes são conceitos fundamentais em Go que são amplamente utilizados para trabalhar com dados e estruturas de controle. Você pode criar mapas para armazenar associações chave-valor, slices para lidar com listas dinâmicas de elementos e structs para representar estruturas de dados mais complexas. 