# go_coverage
Golang: Cobertura de testes em tempo real no VSCode

Para rodar o teste(verificar a cobertura de testes e abre no navegador):
go test --coverprofile coverage.out && go tool cover -html=coverage.out

Para apenas verificar a cobertura de testes:
go test --cover

Se estiver numa maquina windos deve abrir diretamente no navegador, caso esteja rodando o teste via WSL-2 Linux rode o seguite comando:
explorer.exe $(wslpath -w /tmp/cover2658391214/coverage.html)

esse "/tmp/cover2658391214/coverage.html" será mostrado no terminal/output do VSCode, basta substituir e executar o comando