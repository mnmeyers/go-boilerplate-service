# Go Boilerplate App - WIP

### Instructions on installation and running

#### 1. Install MongoDB
* Choose the proper version for your OS to [download here](https://www.mongodb.com/download-center/community). 
* Choose "current release"
* The following instructions for running Mongo assume you have a Mac. 
If not, see [here](https://docs.mongodb.com/manual/administration/install-community/) for alternative
instructions.
    * Run `brew tap mongodb/brew`
    * Run `brew install mongodb-community@4.2`
    * To start the db, run: `brew services start mongodb-community@4.2`
    * To stop the db, run: `brew services stop mongodb-community@4.2`
    * To use the shell (db must be running): run `mongo`
    * If you'd like a DB UI, you can download Robo 3T [here](https://download-test.robomongo.org/mac/robo3t-1.3.1-darwin-x86_64-7419c40.dmg)
        * Other OS options [here](https://robomongo.org/download)
        * Press command + n to create a new connection to your running db
        * Add any name. Address is `localhost`, port is `27017`
        * Click the "Test" button to verify you have valid credentials
        * If it's successful, save and double click on it to make it show up
        in side bar. 
        


#### 2. Run app (Assumes Go development environment is already set up)
* Run:
    * `go get`
    * `go build main.go`
    * `go run main.go`

**Sample data**:
* 
* 

**Sample Requests (in HTTP format)**:
```
GET http://localhost:3000/cus_H6qtZhmCZZAy5R/payments

POST /payments HTTP/1.1
 Host: localhost:3000
 Content-Type: application/json
 
 {
   "account_id": "cus_H6qtZhmCZZAy5R",
   "amount": "11.50"
 }
```

