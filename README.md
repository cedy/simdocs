# Simdocs
## Overview
Simdocs is simple web application to keep track of orders/receipts/etc. This is build for people who still use just a paper copy of their documents.

## Setup & Run
To change a categories please edit list here [models/record.go](https://github.com/cedy/simdocs/blob/master/models/record.go#L22)

Executable takes 3 possitional arguments: username, password, port.
For example, to have it password protected and run it on port 8888 start the application this way:
```
./main my_username my_secret_password 8888
```
You will be able to access the page at http://host_ip/records/

If executed without args, application will run on port 8080 without basic authorization. 
