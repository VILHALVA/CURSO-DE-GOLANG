# TIPOS
As diferenças entre `int32`, `int64`, `float32` e `float64` em Go estão relacionadas ao tamanho e à precisão dos valores que podem ser representados por esses tipos de dados. Aqui estão as principais distinções:

## Inteiros (`int32` e `int64`):
1. **`int32` (32 bits):**
   - É um tipo inteiro de 32 bits que pode representar valores inteiros na faixa de aproximadamente -2 bilhões a +2 bilhões.
   - Adequado para economizar espaço de memória quando a precisão de 32 bits é suficiente.

2. **`int64` (64 bits):**
   - É um tipo inteiro de 64 bits que pode representar valores inteiros em uma faixa muito mais ampla, de aproximadamente -9 quintilhões a +9 quintilhões.
   - Usado quando você precisa de uma faixa maior de valores inteiros ou quando a precisão de 64 bits é necessária.

Em resumo, a principal diferença entre `int32` e `int64` é a quantidade de bits usados para representar valores inteiros, o que afeta a faixa de valores que cada tipo pode armazenar.

## Ponto Flutuante (`float32` e `float64`):

1. **`float32` (32 bits):**
   - É um tipo de ponto flutuante de precisão simples de 32 bits.
   - Oferece menor precisão em comparação com `float64`.
   - É adequado para economizar espaço de memória quando a precisão de 32 bits é suficiente.

2. **`float64` (64 bits):**
   - É um tipo de ponto flutuante de precisão dupla de 64 bits.
   - Oferece maior precisão e faixa em comparação com `float32`.
   - Usado quando alta precisão é necessária ou quando a faixa de valores é grande.

Em resumo, a principal diferença entre `float32` e `float64` é a precisão dos números de ponto flutuante que eles podem representar. `float64` tem mais bits para representação, o que resulta em maior precisão, mas também requer mais espaço de memória. A escolha entre `float32` e `float64` depende das necessidades do seu programa em termos de precisão e economia de memória. Em muitos casos, `float64` é usado por padrão devido à sua maior precisão, a menos que você esteja trabalhando em um ambiente com restrições de memória.