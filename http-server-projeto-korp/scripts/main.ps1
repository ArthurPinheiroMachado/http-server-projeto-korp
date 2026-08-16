$env:HTTP_PORT= "8080"
$env:TIMEOUT_TIME= "3"
$env:PROJECT_NAME= "Projeto Korp"

go run $PSScriptRoot\..\cmd\server\main.go