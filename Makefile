.PHONY: run-srv

run-srv: 
	CompileDaemon -build "go build -o bin/automa8e_clone" -command "./bin/automa8e_clone"