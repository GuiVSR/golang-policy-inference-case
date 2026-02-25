# Documentação do Projeto

## Como Buildar e Fazer Deploy

### 1. Build da imagem para AWS Lambda
Execute este comando para compilar o binário compatível com o ambiente Lambda da AWS:
```bash
GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go build -tags lambda.norpc -o bootstrap cmd/lambda/main.go
```

2. Compactar o arquivo

Após o build, compacte o binário no formato exigido pela AWS:
```bash
zip go_lambda.zip bootstrap
```

3. Faça o upload do arquivo go_lambda.zip para o AWS Lambda através do console AWS ou AWS CLI.


Como Testar

Em desenvolvimento - área reservada para instruções de teste.


Decisões Técnicas
Arquitetura

Escolhi o **AWS Lambda** por oferecer:

    Custo extremamente baixo para aplicações stateless

    Ausência de integrações com bancos de dados

    Escala automática e gerenciamento de infraestrutura reduzido

Estrutura de Código

    Optei por uma estrutura de código simples, com o mínimo de abstrações, visando maximizar a performance

Algoritmo e Estruturas de Dados

    Árvore Customizada: Implementei uma estrutura de dados em árvore própria, pois a gographviz não atendia à necessidade de condições customizáveis

    DFS (Depth-First Search): Escolhi este algoritmo por ser o mais adequado para navegar até as folhas da árvore de forma eficiente

Bibliotecas Utilizadas

    aws-lambda-go/events: Biblioteca oficial para integração com AWS Lambda

    govaluate: Utilizada para avaliação das expressões recebidas através do parâmetro "cond=" no modelo de digraph

    gographviz: Utilizada apenas para o parsing inicial do digraph em uma árvore abstrata

Configuração e Performance

    Memória: Configuração mínima de 128MB

    Performance: Testes realizados com 1000 requisições/segundo apresentaram p95 de 15ms de duração na AWS

    Custos: Baseado na tabela de preços da AWS, com estas configurações, o custo estimado seria de USD$0.002 para cada 1 milhão de requests.

Endpoints

`GET /healthcheck`
```bash
curl --location 'https://noaca6ofui2g2wcie7ipnfp3i40jzipi.lambda-url.sa-east-1.on.aws/healthcheck'
```

`POST - /infer`
```bash
curl --location 'https://noaca6ofui2g2wcie7ipnfp3i40jzipi.lambda-url.sa-east-1.on.aws/infer' \
--header 'Content-Type: application/json' \
--data '{
  "policy_dot": "digraph Policy { start [result=\"\"]; approved [ result=\"approved=true,segment=prime\" ]; rejected [ result=\"approved=false\" ]; review [ result=\"approved=false,segment=manual\" ]; start -> approved [cond=\"age>=18 && score>700\"]; start -> review [cond=\"age>=18 && score<=700\"]; start -> rejected [cond=\"age<18\"]; }",
  "input": {
    "age": 25,
    "score": 720
  }
}'
```

`POST - /visualize`
```bash
curl --location 'https://noaca6ofui2g2wcie7ipnfp3i40jzipi.lambda-url.sa-east-1.on.aws/visualize' \
--header 'Content-Type: application/json' \
--data '{
  "policy_dot": "digraph Policy { start [result=\"\"]; approved [ result=\"approved=true,segment=prime\" ]; rejected [ result=\"approved=false\" ]; review [ result=\"approved=false,segment=manual\" ]; start -> approved [cond=\"age>=18 && score>700\"]; start -> review [cond=\"age>=18 && score<=700\"]; start -> rejected [cond=\"age<18\"]; }"
}'
```