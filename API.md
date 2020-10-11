# API
## Resources
- [GET /vendors](#get-vendors)
- [GET /vendors/{id}](#get-vendorsid)
- [POST /vendors](#post-vendors)
- [PUT /vendors/{id}](#put-vendorsid)
- [DELETE /vendors/{id}](#delete-vendorsid)
- [GET /vendors/{vendor_id}/dishes](#get-vendorsvendor_iddisdhes)
- [GET /vendors/{vendor_id}/dishes/{id}](#get-vendorsvendor_iddisdhesid)
- [POST /vendors/{vendor_id}/dishes](#post-vendorsvendor_iddisdhes)
- [PUT /vendors/{vendor_id}/dishes/{id}](#put-vendorsvendor_iddisdhesid)
- [DELETE /vendors/{vendor_id}/dishes/{id}](#delete-vendorsvendor_iddisdhesid)
- [GET /orders](#get-orders)
- [POST /orders](#post-orders)

### GET /vendors
* **URL Query:**
  ```
  tags[]=[string]
  ```
* **Response:**
  ```
  [
    {
      "id":"jsadfb12",
      "name":"vendor1",
      "description":"foo"
    },
    {
      "id":"joaqweb12",
      "name":"vendor2",
      "description":"foo"
    }
  ]
  ```
### GET /vendors/{id}
* **Response:**
  ```
  {
    "id":"jsadfb12",
    "name":"vendor1",
    "description":"foo"
  }
  ```
### POST /vendors
* **Request:**
  ```
  {
    "name":"new vendor",
    "description":"foo"
  }
  ```
### PUT /vendors/{id}
* **Request:**
  ```
  {
    "name":"updated vendor",
    "description":"foo"
  }
  ```
### DELETE /vendors/{id}
### GET /vendors/{vendor_id}/dishes
* **Response:**
  ```
  [
    {
      "id":"jsadfb12",
      "vendor_id":"asqwfasfv",
      "name":"vendor1",
      "price":1000
    },
    {
      "id":"jsadasdc2",
      "vendor_id":"asqwfasfv",
      "name":"vendor2",
      "price":1000
    }
  ]
  ```
### GET /vendors/{vendor_id}/dishes/{id}
* **Response:**
  ```
  {
    "id":"jsadasdc2",
    "vendor_id":"asqwfasfv",
    "name":"dish1",
    "price":1000
  }
  ```
### POST /vendors/{vendor_id}/dishes
* **Request:**
  ```
  {
    "name":"new dish",
    "price":5000
  }
  ```
### PUT /vendors/{vendor_id}/dishes/{id}
* **Request:**
  ```
  {
    "name":"updated dish",
    "price":5000
  }
  ```
### DELETE /vendors/{vendor_id}/dishes/{id}
### GET /orders
* **Response:**
  ```
  [
    {
      "id":"ohdtvhgb",
      "items":[
        {
          "dish_id":"jsadasdc2",
          "vendor_id":"asqwfasfv",
          "amount":1,
          "request":"request for this dish"
        },
        {
          "dish_id":"asdasdlm2",
          "vendor_id":"asqwfasfv",
          "amount":3,
          "request":"request for this dish"
        },
      ]
    },
    {
      "id":"oijsdagb",
      "items":[
        {
          "dish_id":"ijasddc2",
          "vendor_id":"asqwfasfv",
          "amount":3,
          "request":"request for this dish"
        },
      ]
    }
  ]
  ```
### POST /orders
* **Request:**
  ```
  [
    {
      "dish_id":"jsadasdc2",
      "vendor_id":"asqwfasfv",
      "amount":1,
      "request":"request for this dish"
    },
    {
      "dish_id":"asdasdlm2",
      "vendor_id":"asqwfasfv",
      "amount":3,
      "request":"request for this dish"
    },
  ]
  ```
