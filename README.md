<p align="center">
  <a href="https://github.com/saguywalker/sitcompetence/actions"><img alt="GitHub Actions status" src="https://https://github.com/saguywalker/sitcompetence/workflows/Go/badge.svg"></a>
</p>

## Production
To start the web application run
```
docker compose up -d
```

To stop
```
docker compose down
```

## Development
Open 2 terminal tab
1. Go server
```
go run main.go dev
```

2. Web UI
```
cd frontend
yarn serve
```
