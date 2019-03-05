```console
# run http server
$ RPC_ENDPOINT=http://localhost:10030 go run .
```

```console
# run test xmlrcp server
$ ruby test_server.rb
```

```console
# make a request (httpie)
$ http -v --form POST localhost:30303/login user=foo pass=bar
```
