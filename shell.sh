# curl endpoint
curl localhost:8899/ping
curl localhost:8899/cats
curl -X POST localhost:8899/cats -H 'Content-Type:application/json' -d '{"name":"Puss"}'

# generate mock file with mockery
mockery --dir=db --name=DB --filename=db.go --output=db/mocks --outpkg=db_mocks # tidak dipake
mockery --dir=repository --name=Repository --filename=repository.go --output=repository/mocks --outpkg=repository_mocks
mockery --dir=service --name=Service --filename=service.go --output=service/mocks --outpkg=service_mocks