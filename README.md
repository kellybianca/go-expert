# go-expert

Runtime Go
go build main.go

GOOS=windows go build main.go
ou
GOOS=linux go build main.go

go tool dist listc // lista todas as arquiteturas

go env GOOS GOARCH // mostra a versao do so e a arquitetura
