# Welcome to the order-manager #
## Workstation Setup ##
* Download & install the Go tools and ensure you add the `bin` directory to your `PATH` as well.<br>

* Choose the installation steps for your operating system at golang.org
A Go version of 1.12 or later is needed to run this code.<br>

* Once go is installed, you can run the following command in your terminal to verify the version installed.
```
❯ go version
go version go1.13 darwin/amd64
```
* The default workspace directory is `$HOME/go`. 
If you'd like to use a different directory, you will need to setup your [`GOPATH`](https://github.com/golang/go/wiki/SettingGOPATH) environment variable.

* Test your installation.<br>
    - Clone this repository in your go working directory. <br>
    `example working directoy: $GOPATH/src/github.com`
    - Using your terminal navigate to the `greeter` folder of this project 
    - run the following command where name is your name:
```
❯ go run hello.go Name
Hey {Name}, welcome to the project!
```   
## The project ##
This is a simple project that consists of an in-memory database which can perform
the following operations based on the HTTP request:
* Create a new order: `POST localhost:8080/orders`
* Find all orders: `GET localhost:8080/orders`
* Find a particular order: `GET localhost:8080/orders/{order_id}`
* Delete an order: `DELETE localhost:8080/orders/{order_id}`

To run the project, navigate on the cmd folder:
```
❯ go run server.go
```   

On your browser go to `localhost:8080`
* The browser will display the JSON response.

Create an order:
* The file `postBody.json` contains a mock data to create an order, you can create an order 
by using the following curl command: 
```
curl -v -H "Content-Type: application/json"  --data @postBody.json http://localhost:8080/orders
```   
* You can run the above command multiple times so that more orders are generated.

Find an order:
* On your browser go to `localhost:8080/orders/{order_id}`
* An example order id: `order-0`
* The server serializes the response, but the item is visible in your terminal.

Find all orders:
* On your browser go to `localhost:8080/orders`
* The server serializes the response, but the items are visible in your terminal.

Delete an order:
* On your terminal run the following command:
```
curl -X "DELETE" http://localhost:8080/orders/{order_id}
```

## Technologies Used ##
* Go 1.13
* [`Gin web framework`](https://godoc.org/github.com/gin-gonic/gin)
