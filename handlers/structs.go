package handlers

type Order struct{
    Username int `json:username`
    Bf string `json:bf`
    Lun string `json:lun`
    Din string `json:din`
    Snk string `jsson:snk`
}

type UserDet struct{
    Username int `json: username`
    Password string `json: password`
    Nickname string `json: nickname`
    Hostel string `json: hostel`
    Verified string `json: verified`
}

type messUser struct{
    Username string `json: username`
    Password string `json: password`
}
 //users

type CodeDet struct{
    Username int `json: username`
    Code string `json: code`
}

type Menu struct{
    Day string `json:day`
    Bf string `json:bf`
    Lun string `json:lun`
    Din string `json:din`
    Snk string `json:snk`
}

type SendMenu struct{
    TBf string `json:tbf`
    TLun string `json:lun`
    TDin string `json:tdin`
    TSnk string `json:tsnk`
}

type Code struct{
    //Username int `json: username`
    Bf string `json: bf`
    Lun string `json: lun`
    Din string `json: din`
    Snk string `json: snk`
}

type Confirm struct {
    Status string `json: status`
    Text string `json: text`
} //index
