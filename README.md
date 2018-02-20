Go GraphQL Example
=====

A simple blog app using [go-graphql](https://github.com/graphql-go/graphql) and [Vue.js](https://vuejs.org/).

Get started
--------

Just clone the repository, and get the dependencies

```shell
git clone https://github.com/felipecaputo/go-graphql-example
cd go-graphql-example
go get
go build -o bin/go-graphql-example
./bin/go-graphql-example
```

If you access `http://localhost:8000/` you will see the blog app.

You can also access `http://localhost:8000/graphql`. If accessing from a browser, it will show a `Graphiql` interface that you can query the blog Graql API.