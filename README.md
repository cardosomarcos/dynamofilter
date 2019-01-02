# Dynamo Filter

A ideia do dynamo filter é padronizar os filtros do dynamodb usando 
[guregu/dynamo](https://github.com/guregu/dynamo) para se conectar ao dynamo.

## Install

```bash
go get github.com/cardosomarcos/dynamofilter
```

## Usage and Examples
Definimos um novo filtro
```
filter := dynamofilter.NewFilter()
filter.Equals("name", "Stor Gendibal")
```
Agora, utilizando a [guregu/dynamo](https://github.com/guregu/dynamo) como conector passamos nosso filtro
```
// Variavel que recebe o resultado da query
var data []map[string]interface{}

s, _ := session.NewSession()
// Abre a conexão com o dynamoDB
dynamoConnection := dynamo.New(s, &aws.Config{
    // sa-east-1 => região de são paulo
    Region:     aws.String("sa-east-1"),
	MaxRetries: aws.Int(3),
})

// Cria um novo filtro
filter := dynamofilter.NewFilter()

// Cria um filtro do tipo no nome
filter.Equals("name", "Stor Gendibal")

// Converte para poder usar no guru/dynamo
f, args := filter.Builder()

// Executa a query no dynamo passando nosso filtro
err := dynamoConnection.Table("user").Scan().Filter(f, args...).All(&data)

log.Println("result: ", data)
log.Println("err: ", err)

```

## License

The MIT License (MIT) - see [LICENSE.md](https://github.com/cardosomarcos/dynamofilter/blob/master/LICENSE) for more details