# Options Pattern em Go

## O que é

O **Options Pattern** (padrão de opções) é uma técnica idiomática do Go para passar configurações opcionais a construtores ou funções sem comprometer a clareza da API nem exigir alterações na assinatura da função quando novos parâmetros são adicionados.

## Problema que resolve

Imagine um construtor com muitos parâmetros, onde parte deles é obrigatória e parte é opcional com valores padrão:

```go
// Problemático: difícil de ler, frágil a mudanças
func NewServer(maxConn int, id string, tls bool, optionA bool, optionB bool) *Server
```

Adicionar um novo parâmetro opcional quebra todos os lugares que chamam essa função. O Options Pattern resolve isso.

## Como funciona

### 1. Defina o tipo funcional

```go
type Options func(*Server)
```

Cada "opção" é uma função que recebe um ponteiro para a struct e modifica seus campos.

### 2. Crie funções `With*` para cada opção

```go
func WithOptionA(b bool) Options {
    return func(s *Server) {
        s.optionA = b
    }
}
```

Por convenção, essas funções seguem o prefixo `With`.

### 3. Use variadic no construtor

```go
func NewServer(maxConn int, id string, tls bool, opts ...Options) *Server {
    s := &Server{maxConn: maxConn, id: id, tls: tls}
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

Os parâmetros obrigatórios são posicionais; as opções são aplicadas em loop.

## Uso

```go
// Apenas com valores obrigatórios — opções ficam com o padrão
s1 := NewServer(100, "server1", true)

// Com opções explícitas
s2 := NewServer(100, "server1", true,
    WithOptionA(true),
    WithOptionB(true),
)
```

## Vantagens

| Aspecto | Benefício |
|---|---|
| **Extensibilidade** | Novas opções não quebram chamadas existentes |
| **Legibilidade** | Cada opção é nomeada e auto-explicativa |
| **Valores padrão** | Definidos no construtor, sem necessidade de structs de config separadas |
| **Composabilidade** | Opções podem ser combinadas livremente na chamada |

## Variação: Options sobre uma struct de config

Em projetos maiores, é comum usar uma struct intermediária:

```go
type options struct {
    optionA bool
    optionB bool
}

type Option func(*options)

func WithOptionA(b bool) Option {
    return func(o *options) { o.optionA = b }
}

func NewServer(maxConn int, id string, tls bool, opts ...Option) *Server {
    cfg := &options{} // defaults
    for _, opt := range opts {
        opt(cfg)
    }
    return &Server{maxConn: maxConn, id: id, tls: tls, optionA: cfg.optionA}
}
```

Essa variação separa as preocupações da struct principal das opções de construção.

## Referências

- [Functional Options Pattern — Rob Pike](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)
- [Dave Cheney — Functional options for friendly APIs](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
