# Users endpoint

## This handles the users of the site

### POST /api/users

creates a new user

##### Request Body

```json 

{
    "email": "exemple@gmail.com",
    "password": "10284907",
}

```

##### Response Body

```json 

{
    "email": "exemple@gmail.com",
    "id": 10,
    "is_chirpy_red": false,
}

```

### PUT /api/users

Update user email and password

##### Request Body

```json 

{ 
    "email": "exemple@gmail.com",
    "password": "08124082390",
}

```

It also needs a token to authenticate the user in the authentication header

##### Response Body

```json 

{   
    "id": "1",
    "email": "exemple@gmail.com",
    "is_chirpy_red": false,
}

```