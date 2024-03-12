# web-go-check
This is a small tool that will curl all the urls that you specify in [`main.go`](https://github.com/gonzalinux/web-go-check/blob/ebec5ae00357c9700770cf6d69b1951bcb45d13c/main.go#L33) in the [interval](https://github.com/gonzalinux/web-go-check/blob/ebec5ae00357c9700770cf6d69b1951bcb45d13c/main.go#L10) that you prefer.

If an error ocurred in one of the urls, and the [`credentials.json`](https://github.com/gonzalinux/web-go-check/blob/master/credentials.json) file is correctly filled, the program will send an email to a [selected address](https://github.com/gonzalinux/web-go-check/blob/ebec5ae00357c9700770cf6d69b1951bcb45d13c/main.go#L9)

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
