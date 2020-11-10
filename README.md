# Simdocs

## Disclaimer
This is a hobby project to better learn Go. It breaks bunch of design patterns and principles of software engineering, doesn't have tests and may not run on your machine. I'm ashame of this code, and please look at this code as an example of how not to write code.

## Overview
Simdocs is simple web application to keep track of orders/receipts/etc. This is build for people who still use just a paper copy of their documents.

## Setup & Run
To change categories please edit list here [models/record.go](https://github.com/cedy/simdocs/blob/master/models/record.go#L22)

Executable takes 3 possitional arguments: username, password, port.
For example, to have it password protected and run it on port 8888 start the application this way:
```
./main my_username my_secret_password 8888
```
You will be able to access the page at http://host_ip:8888/records/

If executed without args, application will run on port 8080 without basic authorization. 
