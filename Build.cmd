@echo off

set app=TorMon

set GOARCH=amd64

call :Build

set GOARCH=386

call :Build

pause

exit /b

:Build

set GOOS=windows
set file=Release/%app%_%GOOS%_%GOARCH%.exe
echo Building %file%...
go build -o %file% %app%.go

set GOOS=linux
set file=Release/%app%_%GOOS%_%GOARCH%
echo Building %file%...
go build -o %file% %app%.go

set GOOS=darwin
set file=Release/%app%_%GOOS%_%GOARCH%.app
echo Building %file%...
go build -o %file% %app%.go

exit /b