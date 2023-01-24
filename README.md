# EDA - Event Driven Architecture

### Iniciando com eventos

**Eventos** são efeitos colaterais de um sistema, que aconteceram no passado, e apartir deles outros processos são acionados e decisões são tomadas.

- Situação que ocorreram no passado
- Normalmente deixa efeitos colaterais. Ex.: Porta do carro abriu... Ligar a luz interna.
- Pode atuar de forma internalizada (mesmo sistema) no software ou externalizada (sistemas externos)
- É comum ter a necessidade de externalizar esses eventos.
- Domain Events: Eventos de dominio: Algo aconteceu na camada de dominio, camada de regras de negócio. Uma mudança no estado interno aconteceu --> ex: Agregados


### 3 Tipos de eventos

1. **Event Notification:** Quando um sistema deseja notificar uma mudança. Forma curta de comunicação.
```json 
{"pedido":1, "status":"aprovado"}
```
2. **Event Carried State Transfer:** Formato completo para trafegar as informações. Modalidade para trafegar dados maiores que precisam ser gravadas. "Stream de dados". É enviado o dados e não a notificação de uma mudança.
```json
{"pedido":1, "status":"aprovado", "produtos":[{}], "tax":"1%", "comprador":"Robson"}
```
3. **Event Sourcing:** É a forma de conseguir capturar eventos e armazena-los em um banco de dados, e o mesmo ser usado como **time series database.** Ex.: Saldo de conta bancaria, é a soma/subtração das transações. É possivel fazer um replay de eventos passados, fazer auditoria.

### Event Collaboration
https://martinfowler.com/eaaDev/EventCollaboration.html

**Método tradicional**
Etapas distintas mas que acontecem sequencialmente.
Comprou um produto **->** Estoque do produto comprado **->** Muda o catálogo **->** emite nota **->** Separação

**Com Event Collaboration** Tudo gera eventos.
Parto do principio que já tenho todas as informações para uma determinada ação.
Não precisa pedir, parte do principio que tudo já esta sendo atualizado.

- Fulano Comprou
- Estoque mudou
- Cor mudou
- Não foi emitida
- Erro aconteceu
- Produto mudou a descrição

### Entendendo CQRS (Command Query Responsability Segregation)

https://github.com/keyvanakbary/cqrs-documents
https://eximia.co/command-query-responsibility-segregation-cqrs/

**CQS vs CQRS** / Maior diferença é o nível de granularidade

Separar Comando de Consulta

**Comando: (Escrita/Alteração)** intenção de MUDANÇA do usuário. "Criar produto"
Comando apenas muda/cria algo, não tem retorno porém, esse processo envolve muitas regras que passam pela camada de dominio.

**Consultas: (Leitura)** intensão de obter informação, para por exemplo, hidratar um objeto de dominio, se necessário. Mas, quando separamos o sistema em uma parte que apenas realizam consultas, obviamente que não precisamos do modelo de dominio(DDD). **Foco apenas no dado/retorno.**

### Separação fisica de dados

É possivel guardar os dados fisicamente diferente: Banco de escrita e de leitura.

Com o banco de leitura (NOSQL) podemos ter uma view materializada, evitando JOINS, facilitando a busca.

### Event sourcing vs Command sourcing

**Event sourcing:** Possibilita capturar events e materializar as views.

**Command sourcing:** Estratégia de armazenar os comandos e realizar playback.

### Como implementar CQRS

```bash
/Domain
  /Commands
    /UseCases
      /DAOs vs Repositories
/Queries
  /UseCases #Pode ser meu controller? É possivel evitar complexidade nesse caso?
    /DAOs vs Repositories 
```

### Eventos
- Algo que aconteceu no passado
- Inserir o registro -> Registro Insderido

O que pode ocorrer ao inserir um novo cliente
- Disparar Email
- Publicar uma mensagem na fila
- Notificar um usuário no slack
- Inserir esse usuário no Salesdorce

### Elementos táticos de um contexto de eventos 
- Evento (Carregar dados) -> Elemento principal
- Operações que serão executadas quando um evento é chamado
- Gerenciador dos nossos eventos/operações
  - Registrar os eventos e suas operações
  - Despachar/Fire no evento para que suas operações sejam executadas