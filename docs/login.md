# Login endpoint

## This handles the login of the site

### POST /api/login

##### Request body

```json

{ 
    "email": "email@example.com",
    "password": "password2080",
}

```

##### Response body

```json 

{ 
    "id": "1",
    "email": "email@example.com",
    "is_chirpy_red": false,
    "token": "INAS9FH89WUE890-98HAIUDFHNIUWAEH",
    "refresh_token": "IUBFS89EB8RFBE"
}

```