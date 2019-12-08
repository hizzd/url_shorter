@echo off
SET CGO_ENABLED=0

SET GOARCH=amd64

echo start build url_short on linux
SET GOOS=linux
go build -o ./url_short internal/main.go

echo start build url_short on windows
SET GOOS=windows
go build -o ./url_short.exe internal/main.go


echo build url_short complete.