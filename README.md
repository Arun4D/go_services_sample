# go_services_sample

## service 
Sample microservices 

### Customer

customer microservice to store customer informations


## Build

````
cd service/customer_api

go get .

export DBUSER=root
export DBPASS=mysqlrootpassword

go run .
````
## Setup mysql docker

````
docker pull mysql
docker run --name ad-mysql -e MYSQL_ROOT_PASSWORD=mysqlrootpassword -p 3306:3306 -d mysql

mysql -u root -p
create database customer;
use customer
````

## Run 

````
curl http://localhost:8080/customers

curl http://localhost:8080/customers/

curl http://localhost:8080/customers \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": 0, "first_name": "test","last_name": "test","date_of_birth": "01-01-1984"}'

1

````