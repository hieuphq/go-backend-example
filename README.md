example-be
===================================
Api project use [gin](https://github.com/gin-gonic/gin) + [validator](https://github.com/go-playground/validator) + [goerr](https://github.com/dwarvesf/gerr)

<p align="center">
	<img width="600" src="img/error-l10n.jpg">
</p>

## Language
- Golang

## Structure
- Clean architecture

## How to run
- Set up tool before run
```
make setup
```

- Init Database and services by docker-compose
```
make init
```

- Seed data (if using local)
```
make seed-db-local
```

- Run project
```
make dev
```