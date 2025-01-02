# web-go-check
This is a small tool that will curl all the urls that you specify in [`conf.json`](https://github.com/gonzalinux/web-go-check/blob/master/credentials.json) in the interval that you prefer.

Example of configuration file conf.json
```
{
    "urls": ["https://url1", "https://url2:2999/route", "https://url3.com/"],
    "intervalMinutes": 1,
    "emailTo": "emailHere@email.email"
}
```

If an error ocurred in one of the urls, and the [`credentials.json`](https://github.com/gonzalinux/web-go-check/blob/master/credentials.json) file is correctly filled, the program will send an email to the address in `conf.json`.
![image](https://github.com/gonzalinux/web-go-check/assets/8148642/875f3c4c-c5ec-473e-9e02-5867f871ded8)


And when the page comes up again, it will send an email to alert of it.
![image](https://github.com/gonzalinux/web-go-check/assets/8148642/ab689cce-8c33-4f22-8b0c-1dc3b45ff00a)

To run this program:
```
go run mailer.go main.go
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
