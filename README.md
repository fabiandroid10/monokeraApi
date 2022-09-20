# MonokeraApi

#### Frontend Modules:

* Realizar ordenes para comprar y vender
* Realizar debitos al momento de realizar una compra
* Realizar creditos al momento de realizar una venta


#### Back-End Modules: 

* Escalar la API hasta poder soportar mas de 5000 request por segundo 
* Trazabilidad e historial de transacciones
* Performance: Se desea implementar mecanismo de cache


##### Asumptions:

* Se asume que los usuarios ya existen
* Se desprecian los saldos de cuenta de los usuarios
* Se asume que el transaction id no van a repetirse (se considera over engineering implementar el id con base de datos)

## Desarrollo del servicio

* GO version 1.18.1.
* Docker
* Framework gin-gonic
* Despliegue en AWS Elastic Beanstalk (EC2 Instances, Elastic load Balancer, AutoScaling)
* Monitoreo - Amazon CloudWatch

### Diagrama Opcion 1
![architecture](https://github.com/fabiandroid10/monokeraApi/blob/master/Op1.jpeg)

### Diagrama Arquitectura Elegida
![architecture](https://github.com/fabiandroid10/monokeraApi/blob/master/DiagramArch.jpeg)

### Diagrama AWS
![architecture](https://github.com/fabiandroid10/monokeraApi/blob/master/DiagramAWS.jpeg)


## Consumo del servicio

### Consumo local:

Correr la imagen local: 

1. Descargar docker (según el sistema operativo)
2. Compilar la imagen: 
  - docker build --tag order-book .
3. Correr la imagen:
  - docker run -d -p 80:80 order-book 
4. Realizar la petición usando como body el request encontrado en el archivo ExampleRequest.json al localhost
  - POST http://localhost:80/purchase

Request Example
```bash
{
    "transaction_id": "312312312312",
    "customer_info":{
        "customer_name": "User 1",
        "customer_account": "23131342342342321",
        "ammount": 20000 
    },
    "seller_info":{
        "seller_name": "User 2",
        "seller_account": "23132131"
    }
    }
```


### Consumo hacia AWS:
POST
```
http://goapp-env-2.eba-i3zddipn.us-east-1.elasticbeanstalk.com/purchase
```

#### Logs generados a partir del consumo:

```
2022/09/19 15:22:55.591607 transactionService.go:34: -------Inicio transaccion-------
2022/09/19 15:22:55.591607 transactionService.go:35: Request transaccion: 
2022/09/19 15:22:55.591607 transactionService.go:37: transaction id: 312312312312 Customer name: User 1 Customer account: 23131342342342321 Seller name: User 2 Seller account: 23132131 Transaction ammount 20000
2022/09/19 15:22:55.591607 transactionService.go:54: Response transaccion: 
2022/09/19 15:22:55.591607 transactionService.go:55: Transaction id 312312312312 Customer name: User 1 Customer account: 23131342342342321 Debit Ammount: 20000 Seller Name: User 2 Seller Account: 23132131 Credit Ammount: 20000
2022/09/19 15:22:55.591607 transactionService.go:63: -------Fin Transaccion-------
```

#### Aplicar CI/CD con CodePipeline o con Jenkins:

* Diagrama - aplicacion de CI/CD (Pipeline), tambien se define servicios para escalar horizontalmente los servidores de aplicaciones, aumentar el rendimiento, disminuir la latencia, conseguir tolerancia a fallos y aumentar la disponibilidad.
----
![architecture](https://github.com/fabiandroid10/monokeraApi/blob/master/Flujoci-cd.png)

#### Monitoreo desde CloudWatch:

![cloudwatch](https://github.com/fabiandroid10/monokeraApi/blob/master/CloudWatch.png)
----
![cloudwatch](https://github.com/fabiandroid10/monokeraApi/blob/master/CloudWatch2.png)


