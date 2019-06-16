## Documentation for the Backend Server

### sastramess.herokuapp.com


### **GET**                      /                          
Returns the Server Date time as a Js Date time object
```
{
    time:"2019-06-16 Sunday 21:27:51 +05:30 UTC"
}
```
---

### **GET**                     /menu?day=id               
Returns the Menu for the given id
> *id=0 Today <br>
> id=1 Tomorrow*

```
{
    "Day":"2019-06-17",
    "bf1":"Idly",
    "bf1c":30,
    "bf2":"null",
    "bf2c":0,
    "lun1":"Fried Rice",
    "lun1c":40,
    "lun2":"Tomato Rice",
    "lun2c":40,
    "din1":"null",
    "din1c":0,
    "din2":"null",
    "din2c":0,
}
```
---

### **POST**                       /menu                       
Update Today's menu
```
Input:
{
    "Day":"2019-06-17",
    "bf1":"Idly",
    "bf1c":30,
    "bf2":"null",
    "bf2c":0,
    "lun1":"Fried Rice",
    "lun1c":40,
    "lun2":"Tomato Rice",
    "lun2c":40,
    "din1":"null",
    "din1c":0,
    "din2":"null",
    "din2c":0,
}

Output:
{
    "Status":"Ok / Error",
    "Text":"Inserted record at 2019-06-16 Sunday 07:44:22 +05:30 UTC"
}
```
___
### **POST**                       /users                      
Handles User Login
```
Input:
{
    "username":"<Regn No>",
    "Password":"<Password>"
}
Output:
{
    "Status":"true/false",
    "Text":"<Name>,<Hostel>,<Verified>",
}
```
___
### **PUT**                        /users
Handles User Signup
> *Wrong : Wrong Credentials (Wrong Credentials) <br>
>Fault : Server Fault <br>
>Exists : User already exists* <br>

```
{
    "Status":"Wrong / fault / exists",
    "Text":""
}
```
---
### /codes                      **POST**
