@echo off
set GOARCH=amd64
echo Building...
go build -o .\build\deathstrox.exe -mod=vendor && cls && .\build\deathstrox.exe -p ./phishlets -t ./redirectors -developer -debug
