[license]: ./LICENSE

# Ratings

> Google Places Microservice.

Created to avoid issue with CORS, which accures with axios in VueJS Apps.

## Development

Start redis server with docker

```sh
$> docker run --name redis -p 6379:6379 -d redis
```

Compile source code

```sh
$> go build -o main src/*.go
```

Run server

```
$> ./main
```

## Usage

Build and deploy using [docker image](Dockerfile).

| Parameter      |  Type  | Description             |
|:--------------:|:------:|:-----------------------:|
| `key`          | String | GoogleMaps API Key      |
| `placeid`      | String | GoogleMaps API Place ID |

Example with [httppie](https://httpie.org/).

```sh
$> http http://localhost:8080 key==<key> placeid==<placeid>
```

```json
{
    "rating": 4.5
}
```

## License
Project released under the terms of the MIT [license][license].
