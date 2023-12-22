## API

Setting on `.env` for example 8080
`PORT=8080`

```bash
  go run main.go
```

and will program will Listening and serving HTTP on `:8080`

#### Create Invoice

- `POST`

`http://localhost:8080/api/create-invoice`

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
  `http://localhost:8080/api/update-invoice/43`

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
  `http://localhost:8080/api/invoice`

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
  `http://localhost:8080/api/detail-items/1`

#### Indexing Invoice

`GET`

`http://localhost:8080/api/invoice-indexing?issued_date=2022-10-22&due_date=2023-03-18&status=unpaid`

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
  `http://localhost:8080/api/invoice/43`

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

- Add more integrations
