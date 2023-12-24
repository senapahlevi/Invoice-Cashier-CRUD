
# Invoice Cashier

Backend used for to create cashier invoices quickly and easily. This application is designed to help small and medium businesses manage their sales transactions.



## Tech Stack
You can try using `postman` and use demo 
using these URL `invoice-cashier-crud-production.up.railway.app` 

**Database:** MySQL

**Server:** Golang, Gin, Gorm


## Design

[Figma](https://www.figma.com/file/K27Xkg85MNvWKrJzXudpxI/Invoice-System-Cashier?type=design&node-id=0%3A1&mode=design&t=HA7D6oRAAQtkTd2b-1)


## Run Locally

Clone the project

```bash
  git clone https://github.com/senapahlevi/indonesia-University-API
```

Go to the project directory

```bash
  cd my-project
```

Start the server

```bash
go run main.go
```


## API

Setting on `.env` for example 8080
`PORT=8080`

`URL = http://localhost:8080` OR `invoice-cashier-crud-production.up.railway.app`

```bash
  go run main.go
```
and will program will Listening and serving HTTP on `:8080`
#### Create Invoice 
- `POST`

`URL/api/create-invoice`

body json 
```bash
{
    "customer": "budi",
    "status": "paid",
    "address": "Test address",
    "issued_date": "2023-02-03",
    "due_date": "2020-02-03",
    "subject": "test subject",
    "detail_item_json": "[{\"invoice_id\":1, \"item_name\":\"toys\",\"item_type\":\"service\",\"quantity\":10,\"unit_price\":20 },{\"invoice_id\":1, \"item_name\":\"ice cream\",\"item_type\":\"service\",\"quantity\":1,\"unit_price\":690 }]"
}
```
response json:
```bash
{
    "data": {
        "invoice_id": 44,
        "subject": "test subject",
        "status": "paid",
        "issued_date": "2023-02-03T00:00:00Z",
        "due_date": "2020-02-03T00:00:00Z",
        "sub_total": 7590,
        "tax": 759,
        "grand_total": 6831,
        "detail_item_json": "[{\"invoice_id\":1, \"item_name\":\"toys\",\"item_type\":\"service\",\"quantity\":10,\"unit_price\":20 },{\"invoice_id\":1, \"item_name\":\"ice cream\",\"item_type\":\"service\",\"quantity\":1,\"unit_price\":690 }]",
        "customer": "boedi",
        "address": "Test address",
        "detail_items": null
    }
}
```
#### Update Invoice  
- `PUT`
`URL/api/update-invoice/43`
```bash
{
    "data": {
        "invoice_id": 43,
        "subject": "test subject",
        "status": "paid",
        "issued_date": "2020-01-02T00:00:00Z",
        "due_date": "2020-02-03T00:00:00Z",
        "sub_total": 6990,
        "tax": 699,
        "grand_total": 6291,
        "detail_item_json": "[{ \"item_name\":\"Toyowheels\",\"item_type\":\"tools\",\"quantity\":1,\"unit_price\":6990}]",
        "customer": "rudi tabuti",
        "address": "newcastle",
        "detail_items": null
    }
}
```

#### Detail All Invoice /GET All invoice
- `GET`
`URL/api/invoice`

```bash
{
    "data": [
        {
            "invoice_id": 43,
            "subject": "test subject",
            "status": "paid",
            "issued_date": "2020-01-02T07:00:00+07:00",
            "due_date": "2020-02-03T07:00:00+07:00",
            "sub_total": 6990,
            "tax": 699,
            "grand_total": 6291,
            "detail_item_json": "[{ \"item_name\":\"Toyowheels\",\"item_type\":\"tools\",\"quantity\":1,\"unit_price\":6990}]",
            "customer": "rudi tabuti",
            "address": "newcastle",
            "detail_items": null
        },
        {
            "invoice_id": 44,
            "subject": "test subject",
            "status": "paid",
            "issued_date": "2023-02-03T07:00:00+07:00",
            "due_date": "2020-02-03T07:00:00+07:00",
            "sub_total": 7590,
            "tax": 759,
            "grand_total": 6831,
            "detail_item_json": "[{\"invoice_id\":1, \"item_name\":\"toys\",\"item_type\":\"service\",\"quantity\":10,\"unit_price\":20 },{\"invoice_id\":1, \"item_name\":\"ice cream\",\"item_type\":\"service\",\"quantity\":1,\"unit_price\":690 }]",
            "customer": "boedi",
            "address": "Test address",
            "detail_items": null
        }
    ]
}
```
#### Delete Invoice button  
- `DELETE`
`URL/api/detail-items/1`
 
 
#### Indexing Invoice 
`GET`

`URL/api/invoice-indexing?issued_date=2022-10-22&due_date=2023-03-18&status=unpaid` 



response json:
```bash
{
    "data": [
        {
            "invoice_id": 43,
            "subject": "test subject",
            "status": "paid",
            "issued_date": "2023-02-03T07:00:00+07:00",
            "due_date": "2020-02-03T07:00:00+07:00",
            "sub_total": 7590,
            "tax": 759,
            "grand_total": 6831,
            "detail_item_json": "[{\"invoice_id\":1, \"item_name\":\"toys\",\"item_type\":\"service\",\"quantity\":10,\"unit_price\":20 },{\"invoice_id\":1, \"item_name\":\"ice cream\",\"item_type\":\"service\",\"quantity\":1,\"unit_price\":690 }]",
            "customer": "boedi",
            "address": "Test address",
            "detail_items": [
                {
                    "item_id": 35,
                    "InvoiceID": 43,
                    "item_name": "toys",
                    "item_type": "service",
                    "quantity": 10,
                    "unit_price": 20,
                    "amount": 200
                },
                {
                    "item_id": 36,
                    "InvoiceID": 43,
                    "item_name": "ice cream",
                    "item_type": "service",
                    "quantity": 1,
                    "unit_price": 690,
                    "amount": 690
                }
            ]
        }
    ]
}
```
#### Detail ID Invoice  

- `GET`
`URL/api/invoice/43`
```bash
{
    "data": {
        "invoice_id": 43,
        "subject": "test subject",
        "status": "paid",
        "issued_date": "2020-01-02T07:00:00+07:00",
        "due_date": "2020-02-03T07:00:00+07:00",
        "sub_total": 6990,
        "tax": 699,
        "grand_total": 6291,
        "detail_item_json": "[{ \"item_name\":\"Toyowheels\",\"item_type\":\"tools\",\"quantity\":1,\"unit_price\":6990}]",
        "customer": "rudi tabuti",
        "address": "newcastle",
        "detail_items": [
            {
                "item_id": 39,
                "InvoiceID": 43,
                "item_name": "Toyowheels",
                "item_type": "tools",
                "quantity": 1,
                "unit_price": 6990,
                "amount": 6990
            }
        ]
    }
}
```



