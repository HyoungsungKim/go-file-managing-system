# Golang file server
Environment of docker:
```
SERVER_ADDRESS="172.32.0.1"
SERVER_PORT="9010"
```

If you run `docker-compose up`, server will listen on `172.32.0.1:9010`.
  
Thus, if you want to test below commends by attaching container, you have to replace `localhost` to `172.32.0.1`. However, if you want to test below commends on host, you don't need to replace `localhost`.

---

## Example commands for testing http methods
Example commends for testing http methods.

1. Check server is listening on defined IP address
2. Test uploading file (single file)

### Check server is listening on defined IP address
> Uncomment `router.GET("/ping", .. )` of `src/internal/controller/httpMethodHandler.go` to test server is listening on defined IP address

Check server is listening on defined IP address using GET method.

Example
```
curl -X GET http://localhost:9010/ping
```


### Test uploading file (single file)
Upload `helloworld.txt` to `storage/SOME-UUID` using POST method.

Example
```
curl -X POST http://localhost/upload/SOME-UUID -F "file=@./testResource/helloworld.txt" -H "Content-Type: multipart/form-data"
```
