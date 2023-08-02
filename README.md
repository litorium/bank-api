# Bank-Merchant API

## Steps to use

### I use 8080 port (localhost:8080)

#### First you need to create a account on

```golang
/user //POST METHOD

{
    "username" : "username",
    "password" : "password"
}
```

#### After that you need to login to get the token

```golang
/login //POST METHOD

{
    "username" : "username",
    "password" : "password"
}
```

#### In user you can CRUD (must login except post method)

```golang
/user //POST METHOD

/user //GET METHOD (No need id because already use sessions)

/user //DELETE METHOD (No need id because already use sessions)

/user //PUT METHOD (No need id because already use sessions)
{
    "username" : "username",
    "password" : "password"
}
```

#### In merchant u can CRUD no need to login (Dummy Merchant)

```golang
/merchant //POST METHOD
{
    "name" : "nama_merchant",
    "no_rek": "no_rek"
}

/merchant/nama_merchant //GET METHOD

/merchant //PUT METHOD
{
    "id" : "id",
    "name" : "nama_merchant",
    "no_rek": "no_rek"
}

/merchant/nama_merchant //DELETE METHOD
```

#### In payment u can read and create (must login)

```golang
/payment //POST METHOD (No need id because already use sessions)
{
    "merchant_no_rek" : "no_rek",
    "amount" : "amount"
}

/payment //GET METHOD (No need id because already use sessions)
it will show all ur userid transaction
```

#### You can logout

```golang
/logout //POST METHOD
```
