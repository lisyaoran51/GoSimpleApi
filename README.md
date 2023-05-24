# GoSimpleApi

### start

```bash
make build
docker run -p 5000:5000 lisyaoran51/go-simple-api
```


### postman

```json
{
	"info": {
		"_postman_id": "f6751d1e-04cb-40db-9ec0-3990a0b2b9f3",
		"name": "GoSimpleApi",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "23492686"
	},
	"item": [
		{
			"name": "create",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"account\":\"abcabc1\",\n    \"password\":\"ppppppppP1\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:5000/v1/users/create",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"v1",
						"users",
						"create"
					]
				}
			},
			"response": []
		},
		{
			"name": "validate",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"account\":\"abcabc12\",\n    \"password\":\"ppppppwppP1\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:5000/v1/users/validate",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"v1",
						"users",
						"validate"
					]
				}
			},
			"response": []
		}
	]
}
```