// POST /v1/user/register
const body = {   
    "username": "",
    "name" : "",
    "password": "",
}

//response mengandung token
//response
//200
const response = {
    "message" : "User Registered Succesfully",
    "data" : {
        "username" : "",
        "name" : "",
        "access_token" : "", //2 menit login
    }
}

// POST /v1/user/login
const body_1 = {
    "username": "",
    "password": "",
}