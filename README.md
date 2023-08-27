# REST API using Go maps
This is the first REST API implemented by me. I used the following:
- **Router**. Selected for routing framework GIN `gin-gonic/gin`.
- Running HTTP server by using `net/http`.
- Temporary storage by using Go maps data structure.

---

## Run:
For the first run:

```console
go mod download
```

Then:

```console
go run ./cmd/main.go
```

---

## Structs (JSON)
```go
type Client struct {
	ID      int    `json:"id"`
	FIO     string `json:"fio"`
	Phone   string `json:"phone"`
	Adress  string `json:"adress"`
	OrderID int    `json:"order_id"`
}

type WorkingPosition struct {
	PositionID int    `json:"position_id"`
	Name       string `json:"name"`
}

type Staff struct {
	StaffID     int    `json:"staff_id"`
	FIO         string `json:"fio"`
	PositionID  int    `json:"position_id"`
	DateOfBirth string `json:"date_of_birth"`
	Salary      int    `json:"salary"`
	Phone       string `json:"phone"`
	Adress      string `json:"adress"`
}

type Order struct {
	OrderID       int     `json:"order_id"`
	ClientID      int     `json:"client_id"`
	Name          string  `json:"name"`
	Cost          float64 `json:"cost"`
	PaymentMethod string  `json:"payment_method"`
	Date          string  `json:"date"`
	FinishDate    string  `json:"finish_date"`
	Description   string  `json:"description"`
}
```


## API:
### WorkingPosition:
### POST api/working_positions
Creates new name of working position for staff

- Example Input:
```json
{
    "name": "Director"
}
```
- Example Output:
```json
{
    "position_id": 1
}
```



### GET api/working_positions
Returns all names of working positions
- Example Output:
```json
{
  "1": {
    "position_id": 1,
    "name": "Director"
  },
  "2": {
    "position_id": 2,
    "name": "Accountant"
  },
  "3": {
    "position_id": 3,
    "name": "Foreman"
  },
  "4": {
    "position_id": 4,
    "name": "Mechanic"
  }
}
```

### GET api/working_positions/:position_id
Returns the name of the working position with position_id determined in the URL

- Example Output №1 **api/working_positions/1** :
```json
{
    "position_id": 1,
    "name": "Director"
}
```
- Example Output №2:
```json
{
    "message": "Position not found."
}
```

### PUT api/working_positions/:position_id
Updates the name of the working position

- Example Input:
```json
{
    "name": "Driver"
}
```
Example Output:
```json
{
    "position_id": 1,
    "name": "Driver"
}
```
### DELETE api/working_positions/:position_id
Deletes the name of the working position
- Output:
```
Position deleted
```
---
### Staff:

### POST api/working_positions/:position_id/staff
Create employee using position_id in the URL

- Example Input for **api/working_positions/1/staff**:
```json
{
    "fio": "Ivanov Ivan Ivanovich",
    "date_of_birth": "01.05.2000",
    "salary": 100000,
    "phone": "+79998887766",
    "adress": "Novosibirsk, st. Titova 1"
}
```
- Example Output:
```json
{
    "position_id": 1,
    "staff_id": 1
}
```

### GET api/staff
Returns information about all employees

- Example Output:
```json
{
    "1": {
        "staff_id": 1,
        "fio": "Ivanov Ivan Ivanovich",
        "position_id": 1,
        "date_of_birth": "01.05.2000",
        "salary": 100000,
        "phone": "+79998887766",
        "adress": "Novosibirsk, st. Titova 1"
    },
    "2": {
        "staff_id": 2,
        "fio": "Semenov Semen Semenovich",
        "position_id": 2,
        "date_of_birth": "01.05.2003",
        "salary": 50000,
        "phone": "",
        "adress": "Novosibirsk, st. Karla Marksa 10"
    }
}
```


### GET api/staff/:staff_id
Returns information about employee with staff_id

- Example Output for **api/staff/1**:
```json
{
    "staff_id": 1,
    "fio": "Ivanov Ivan Ivanovich",
    "position_id": 1,
    "date_of_birth": "01.05.2000",
    "salary": 100000,
    "phone": "+79998887766",
    "adress": "Novosibirsk, st. Titova 1"
}
```

### PUT api/staff/:staff_id
Updates information about employee

- Example Input for **api/staff/1** :
```json
{
    "fio": "Ivanov Maksim Semenovich",
    "date_of_birth": "02.03.2001",
    "position_id" : 1,
    "salary": 100000,
    "phone": "+79233657622",
    "adress": "Novosibirsk, st. Karla Marksa 10"
}
```

- Example Output:
```json
{
  "staff_id" : 1
}
```

### DELETE api/staff/:staff_id
Deletes information about employee

- Example output for **api/staff/1** :
```
Employee deleted
```
---
### Client

### POST api/clients
Creates new client
- Example Input:
```json
{
    "fio" : "Krasnov Viktor Mikhailovich",
    "phone" : "+79997556543",
    "adress" : "Moscow, st. Lenina 5"
}
```
- Example Output:
```json
{
    "id": 1
}
```
### GET api/clients
Returns information about all clients

- Example Output:
```json
{
    "1": {
        "id": 1,
        "fio": "Krasnov Viktor Mikhailovich",
        "phone": "+79997556543",
        "adress": "Moscow, st. Lenina 5",
        "order_id": 0
    },
    "2": {
        "id": 2,
        "fio": "Maksimov Nikolai Fedorovich",
        "phone": "+79993332232",
        "adress": "Kazan, st. Pushkina 10",
        "order_id": 0
    }
}
```

### GET api/clients/:id
Returns information about client

- Example Output for **api/clients/1**:
```json
{
    "id": 1,
    "fio": "Krasnov Viktor Mikhailovich",
    "phone": "+79997556543",
    "adress": "Moscow, st. Lenina 5",
    "order_id": 0
}
```

### PUT api/clients/:id
Updates information about client

- Example Input for **api/clients/2** (update only phone number):
```json
{
    "fio" : "Maksimov Nikolai Fedorovich",
    "phone" : "+79993325533",
    "adress" : "Kazan, st. Pushkina 10"
}
```

- Example Output:
```json
{
    "id": 2
}
```
### DELETE api/clients/:id

- Example Output:
```
Client deleted
```

---

### Order

###  POST api/orders
Creates new order

- Example Input (with client_id):
```json
{
    "client_id": 1,
	"name": "Build the roof of the house",
	"cost": 10000,
	"payment_method": "Internet bank",
	"date": "05.12.2022",
	"finish_date": "12.01.2023",
	"description": "Construction of the roof of a private house, 15x20 meters"
}
```
- Example Output:
```json
{
    "order_id": 1
}
```
Now we can check order from client with client_id = 1:  
**GET api/clients/1/order**
- Example Output:
```json
{
    "order_id": 1
}
```  
**GET api/orders/1/client**
- Example Output:
```json
{
    "client_id": 1
}
```

### GET api/orders
Returns information about all orders
- Example Output:
```json
{
    "1": {
        "order_id": 1,
        "client_id": 1,
        "name": "Build the roof of the house",
        "cost": 10000,
        "payment_method": "Internet bank",
        "date": "05.12.2022",
        "finish_date": "12.01.2023",
        "description": "Construction of the roof of a private house, 15x20 meters"
    }
}
```

### GET api/orders/:order_id
Returns information about order
- Example Output for **api/orders/1**:
```json
{
    "order_id": 1,
    "client_id": 1,
    "name": "Build the roof of the house",
    "cost": 10000,
    "payment_method": "Internet bank",
    "date": "05.12.2022",
    "finish_date": "12.01.2023",
    "description": "Construction of the roof of a private house, 15x20 meters"
}
```

### PUT api/orders/:order_id
Updates information about order
- Example Input for **api/orders/1**:
```json
{
    "client_id": 1,
	"name": "Build house",
	"cost": 5000000,
	"payment_method": "Internet bank",
	"date": "10.05.2023",
	"finish_date": "12.05.2024",
	"description": "Build a house outside the city, budget 5000000"
}
```
- Example Output:
```json
{
    "order_id": 1
}
```

### DELETE api/orders/:order_id
- Example Output for **api/orders/1**:
```
Order deleted
```

---
## TODO:
- Do REST API using database.

---
Mikhail Rogalsky