# Godel

WIP

## Local setup

1. Build the application
```
go build
```

2. Set the expected environment variables in `.env.local` file

3. Start the docker-compose for the database
```
docker-compose -f docker-compose.db.yml up
```

4. Start the server
```
./go-del
```