# cylo-commerce-platform-assessment
This is the README file for the commerce platform (only for assessment)

**MAIN FRAMEWORK**
+ gin framwork

**Some software development principles, patterns that are used in this repository**
+ Singleton (apply for get configuration, database manager)
+ Data Mapper
+ Single responsibility principle
+ Open/closed principle
+ Dependency inversion principle


**SOURCE CODE STRUCTURE**
```
    main.go
    core 
        | config - For getting the config to run service
        | database - For managing the database connection
        | model - For defining some base models
    data  - seeding data and schema script
    middleware - For defining some middleware to use
        | authentication.middleware.go - check and get profile
        | core-response.middleware.go - transform the response data
    modules - For defining the modules of service
        | products - Product module
        | activities - Activity module
    public - serve statistic files
```

**THE REQUIREMENT BEFORE RUN ON LOCAL**
+ Mysql server
+ Config connection in config file (config.dev.json)

**SEEDING DATA**
```
go run ./data/seed-script.go
```

**HOW TO RUN**

If you want to active the hot reload for this project, so please install the gin package first and then
```
gin --appPort 5000 --excludeDir ./public
```
For case don't need, it can be run with simple command
```
go run main.go
```
And then access to the link
```
http://localhost:3000
```

**HOW TO RUN TEST**
```
go test ./...
```

**HOW TO FORMAT SOURCE CODE**

Need to setup golangci-lint before run the lint: Reference this link to install it: https://golangci-lint.run/usage/install/
Then use below command

```
golangci-lint run 
```

**SOME CURL TO TEST DATA**
1/ The curl to get the list activity
```
curl --location --request GET 'http://localhost:3000/api/activities?skip=0&take=10' \
--header 'X-Person: Jerry'
```


