central:
	go run Central/main.go	#Ejemplo-dist/Central/main.go

laboratorio:
	go run Ejemplo-dist/Laboratorios/pripyat.go  #Laboratorios/pripyat.go
	#go run Ejemplo-dist/Laboratorios/kampala.go	Laboratorios/kampala.go
	#go run Ejemplo-dist/Laboratorios/renca.go		Laboratorios/renca.go
	#go run Ejemplo-dist/Laboratorios/pohang.go		Laboratorios/pohang.go

#build:
	#docker build -t lab1 .

#d-central:
		#docker run -it --rm -P lab1:latest go run Ejemplo-dist/Central/main.go	#Central/main.go

#docker-central: build d-central

#d-laboratorio:
	#OJO con las direcciones OJO
	#docker run -it --rm -P -p 50051:50051 lab1:latest go run Ejemplo-dist/Laboratorios/pripyat.go		Laboratorios/pripyat.go
	#docker run -it --rm -P -p 50051:50051 lab1:latest go run Ejemplo-dist/Laboratorios/kampala.go		Laboratorios/kampala.go
	#docker run -it --rm -P -p 50051:50051 lab1:latest go run Ejemplo-dist/Laboratorios/renca.go			Laboratorios/renca.go
	#docker run -it --rm -P -p 50051:50051 lab1:latest go run Ejemplo-dist/Laboratorios/pohang.go		  Laboratorios/pohang.go

#docker-laboratorio: build d-laboratorio
