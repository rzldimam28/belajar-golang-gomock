curl localhost:8899/ping
curl localhost:8899/cats
curl -X POST localhost:8899/cats -H 'Content-Type:application/json' -d '{"name":"Puss"}'