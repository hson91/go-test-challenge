
- [Requirements](#requirements)
- [Run](#run)
- [Build](#build)
- [Usage](#usage)
- [Find User](#find-user)
- [Find Organization](#find-organization)
- [Find Ticket](#find-ticket)

## Requirements

- Go 1.10 or higher. We aim to support the latest supported versions of go.

## Run
```go 
    git clone https://github.com/hson91/go-test-challenge.git
    cd go-test-challenge
    go run app.go
```

## Build
```go 
    go build -o app
```

## Test
```go 
    go test -v
```

## Usage
- Help: ```$ help``` 

- Reload:```$ reload```

- Exit:  ```$ quit``` or ```$ exit```

- Show struct user:  ```$ user struct```

- Show struct organization:  ```$ organization struct```

- Show struct ticket: ```$ tiket struct```

## Find User
Syntax: ```$ user find attribute=value ```\
Attributes of User find
``` 
    id int
    organization_id int
    alias <string>
    email <string>
    name <string>
    phone <string>
    role `agent` or `admin` or `end-user`
    tag <string>
    active true or false
    shared true or false
    suspended true or false
    verified true or false
```
**Example with find user:**\
Find user by id = 1\
```$ user find id=1```

Find User with name = "Francisca Rasmussen"\
```$ user find name="Francisca Rasmussen"```

Find User with active = true\
```$ user find active=true```

## Find Organization
Syntax: ```$ organization find attribute=value```\
Attributes of Organization find
``` 
    id int
    name <string>
    tag <string>
    domain <string>
    shared_tickets true or false
```
**Example with find organization:**\
Find Organization with id = 124\
```$ organization find id=124```

Find Organization with name="Qualitern"\
```$ organization find name='Qualitern'```

Find Organization with shared_tickets=true\
```$ organization find shared_tickets=true```

## Find Ticket
Syntax: ```$ ticket find attribute=value ```\
Attributes of Organization find
``` 
    id : int
    status :`open`, `pending`, `hold`, `solved`, `closed`
    subject : string
    type : `incident`, `problem`, `question`,`task`
    tag :string
    via : `web`, `chat`, `voice`
    submitter_id : int
    assignee_id : int
    organization_id : int
```
**Example with find ticket:** \
Find ticket with id = "436bf9b0-1147-4c0a-8439-6f79833bff5b"\
```$ ticket find id="436bf9b0-1147-4c0a-8439-6f79833bff5b" ```

Find ticket with status = pending\
```$ ticket find status=pending ```
## Thanks and Acknowledgement 
