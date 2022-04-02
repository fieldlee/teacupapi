# TeamCup API

## Tech stack

* API: gin
* ORM: gorm
* DB: MySQL
* Redis

## API documentation

Generate a swagger spec document from source by [go-swagger](https://github.com/go-swagger/go-swagger).

```
swagger generate spec -m -o swagger.json
// or
go generate ./...
```

View swagger document.

```
swagger serve --flavor=swagger swagger.json
```

Or open swagger.json through https://editor.swagger.io/
