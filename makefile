central:
	go run Central/main.go	#Ejemplo-dist/Central/main.go

laboratorio:
	go run Laboratorios/pripyat.go  #Laboratorios/pripyat.go
	#go run Laboratorios/kampala.go	Laboratorios/kampala.go
	#go run Laboratorios/renca.go		Laboratorios/renca.go
	#go run Laboratorios/pohang.go		Laboratorios/pohang.go

docker-central:
	docker build -t lab1 .
	docker run -it --rm -P lab1:latest go run Central/main.go	#Central/main.go

docker-laboratorio:
	docker build -t lab1 .
	docker run -it --rm -P -p 50051:50051 lab1:latest go run Laboratorios/pripyat.go
	#docker run -it --rm -P -p 50051:50051 lab1:latest go run Laboratorios/kampala.go
	#docker run -it --rm -P -p 50051:50051 lab1:latest go run Laboratorios/renca.go
	#docker run -it --rm -P -p 50051:50051 lab1:latest go run Laboratorios/pohang.go
