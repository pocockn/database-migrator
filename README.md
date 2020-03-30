# database-migrator

Migrates and seeds a local database

## Getting Started

```
make run
```

I run it within a docker-compose stack to migrate and seed the local database to help with local development. 

Example docker-compose file

```
  database:
    image: mysql:5.6.34
    environment:
      - MYSQL_ROOT_PASSWORD=password
      - MYSQL_DATABASE=database
      - MYSQL_USER=username
      - MYSQL_PASSWORD=password
    expose:
      - 3306
    ports:
      - 3306:3306
  data-migrator:
    image: pococknick91/database-migrator:0.16.0
    environment:
      - ENV=docker
```

### Installing

```
make install
```

## Running the tests

```
make test
```

## Built With

* [Go](https://golang.org/)
* [GORM](https://gorm.io/) - ORM Library

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/pocockn/recs-api/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
