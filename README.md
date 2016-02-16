# Counter

Counter is a web application that increases the value of a counter every time
it is clicked. The application uses PostgreSQL as its store.

## Connecting to PostgreSQL

This is configured by setting the environment variable `POSTGRES_URL` to a postgres endpoint such
as `postgres://localhost`.

```
# With binary and local Postgres
$ POSTGRES_URL=postgres://localhost counter-linux

# With docker and local Postgres
docker run -d -e POSTGRES_URL=postgres://localhost stigkj/counter

# Or with link, assuming a Postgres container named postgres is running
docker run -d --link postgres -e POSTGRES_URL=postgres://postgres@postgres stigkj/counter
```

