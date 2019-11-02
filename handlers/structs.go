package handlers

type Order struct{
    Username int `json:username`
    Bf int `json:bf`
    Lun int `json:lun`
    Din int `json:din`
    Snk int `jsson:snk`
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
} //postmenu

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
