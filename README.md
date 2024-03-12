# web-go-check
This is a small tool that will curl all the urls that you specify in `main.go` in the interval that you prefer.

If an error ocurred in one of the urls, and the credentials.json file is correctly filled, the program will send an email to a selected address

To run this program:
```
go run .\mailer.go .\main.go
```

Or build it and run it

```
go build
```

Then in windows:
```
start web-go-check.exe
```
(or double click)

And in linux
```
./web-go-check
```