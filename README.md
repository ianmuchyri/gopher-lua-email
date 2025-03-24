# gopher-lua-email

## Variables to set

In main.go set the following variables:
```
from:<your google account>
password: <your google application password>
```
to get the password you can follow this [tutorial](https://dev.to/go/sending-e-mail-from-gmail-using-golang-20bi)

In temp_check.lua set the following variables:
```
recipient: <email of the recepient>
```

You can change the value of the temperature to send different emails by adjusting this variable
```
local temperature = 32 
```