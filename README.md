restdemo
========

A demo showing how to use Go to serve an Oauth2 protected REST API.

To get the binary, run:

    go install github.com/shutej/restdemo/basicrestdemo

You will **still** need to clone the project so you can get at the SSL
certificates and the static files.  Run the "basicrestdemo" binary in the root
of the repository, then
[https://localhost:8000/authorize?response_type=token&client_id=client1&redirect_uri=https:%2F%2Flocalhost%3A8000%2Fstatic%2Findex.html&scope=](click
here).