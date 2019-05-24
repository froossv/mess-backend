package handlers

type Order struct{
    Username int `json:username`
    Bf1 int `json:bf1`
    Bf2 int `json:bf2`
    Lun1 int `json:lun1`
    Lun2 int `json:lun2`
    Din1 int `json:din1`
    Din2 int `json:din2`
}

type UserDet struct{
    Username int `json: username`
    Password string `json: password`
    Name string `json: name`
    Hostel string `json: hostel`
    Verified string `json: verified`
} //users

type CodeDet struct{
    Username int `json: username`
    Code string `json: code`
}

type Menu struct{
    Day string `json:day`
    Bf1 string `json:bf1`
    Bf1c int `json:bf1c`
    Bf2 string `json:bf2`
    Bf2c int `json:bf2c`
    Lun1 string `json:lun1`
    Lun1c int `json:lun1c`
    Lun2 string `json:lun2`
    Lun2c int `json:lun2c`
    Din1 string `json:din1`
    Din1c int `json:din1c`
    Din2 string `json:din2`
    Din2c int `json:din2c`
} //postmenu

type Confirm struct {
    Status string `json: status`
    Text string `json: text`
} //index
