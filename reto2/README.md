# ST0263-2887 Topicos especiales en telematica
#
# Estudiante(s): Kevin Alejandro Sossa Chavarria, kasossac@eafit.edu.co
#
# Profesor: Edwin Montoya

# Reto2
#
# Descripción
Ddiseño e implementación de mínimo 2 microservicios básicos que ofrecen ambos un
servicio al API Gateway y que se deben comunicar por un middleware RPC y por un middleware
MOM. Cada uno de los microservicios debe soportar concurrencia. Para la comunicación RPC dr usa el
middleware API REST y gRPC, y para la comunicación MOM se ua RabbitMQ.

# información general de diseño de alto nivel, arquitectura, patrones, mejores prácticas utilizadas.
Todo el desarrollo esta basado en una arquitectura DDD

# Descripción del ambiente de desarrollo y técnico.
Para la ejecucion debemos garantizar que las maquinas tengan instalado:
- Docker
- Golang
- Make
  
Con esto garantizado en cada una de las maquinas correspondientes usar el comando de make dependiendo de cual instancia queremos ejecutar.
- run-gRPC-server
- run-mom-server
- run-api-gw
- run-rabbitmq

# Ips de instancias
gRPC-server: 44.217.39.70
rabbitmq: 18.207.23.165
api-gw: 3.222.98.226
rabbitmq-server: 54.197.213.223

# referencias:

[gRPC GO](https://grpc.io/docs/languages/go/basics/)

[RabbitMQ Go](https://www.rabbitmq.com/tutorials/tutorial-one-go.html)
