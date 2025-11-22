# RC Pairing Interview

https://www.recurse.com/pairing-tasks

## Database server

> Before your interview, write a program that runs a server that is accessible on
> http://localhost:4000/. When your server receives a request on
> http://localhost:4000/set?somekey=somevalue it should store the passed key and value in memory.
> When it receives a request on http://localhost:4000/get?key=somekey it should return the value
> stored at somekey.
During your interview, you will pair on saving the data to a file. You can start with simply
appending each write to the file, and work on making it more efficient if you have time.

## Usage

Start DB

```sh
go run cmd/db/main.go
```

Make requests

```sh
curl -i 'http://localhost:4000/get?a'
curl -i 'http://localhost:4000/set?a=1'
curl -i 'http://localhost:4000/get?a'
curl -i 'http://localhost:4000/set?a=2'
curl -i 'http://localhost:4000/get?a'
```

## Tests

Run tests

```sh
go test ./...
```

## Thoughts

I would probably

* turn `http://localhost:4000/set?somekey=somevalue` into `POST http://localhost:4000/somekey` with
the value in the body
* add `DELETE http://localhost:4000/somekey` as I cannot differentiate if a user wants to set the
key to an empty string or delete it as these are the same to me `http://localhost:4000/set?somekey`
and `http://localhost:4000/set?somekey=`

