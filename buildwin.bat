set GOOS=windows
set GOARCH=amd64
go build -tags prod -o ../iitdapiserver.exe ./cmd/iitd_server/main.go