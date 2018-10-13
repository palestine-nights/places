[license]: ./LICENSE

# Ratings

> Google Places Microservice.

Created to avoid issue with CORS, which accures with axios in VueJS Apps.

## Usage

Build and deploy using [docker image](Dockerfile).

| Parameter      |  Type  | Description             |
|:--------------:|:------:|:-----------------------:|
| `key`          | String | GoogleMaps API Key      |
| `placeid`      | String | GoogleMaps API Place ID |

Example with [httppie](https://httpie.org/).

```sh
http http://localhost:8000 key==<key> placeid==<placeid>
```

```json
{
    "rating": 4.5
}
```

## License
Project released under the terms of the MIT [license][license].
