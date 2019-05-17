package handlers

type UserDet struct{
    Username int `json: username`
    Password int `json: password`
    Name string `json: name`
    Hostel string `json: hostel`
} //users

type Menu struct{
    day string `json:day`
    bf1 string `json:bf`
    bf1c int `json:bf1c`
    bf2 string `json:bf`
    bf2c int `json:bf1c`
    lun1 string `json:lun`
    lun1c int `json:bf1c`
    lun2 string `json:lun`
    lun2c int `json:bf1c`
    din1 string `json:din`
    din1c int `json:bf1c`
    din2 string `json:din`
    din2c int `json:bf1c`
} //postmenu

type Menus []Menu //postMenu

type Confirm struct {
    Status string `json: status`
    Text string `json: text`
} //index
