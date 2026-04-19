# memdb

memdb is a simple key-value in-memory database. Just for learning. This database will contain different data types to store data.

## Features

- Multiple data types support
  - [ ] String
  - [ ] Hash
  - [ ] Set
- [ ] HTTP/CLI access
- [ ] Concurrent
- [ ] Eviction policies like LFU, LRU
- [ ] TTL (per key)

## Running DB

Running the command

> memdb --port 0000 --eviction-policy LRU

## HTTP API Specification

All request operations will be a *POST* request, the path defines the data type to use and the operation. This is the path structure:

> POST /{data_type}/{operation}

And the body will be json format that contains a *key* and *value* properties for requests. This is the body structure:

```text
{
    key: string,
    value: any?
}
```

**NOTE**: *value* is omited depending the operation.

### Error responses

Error status codes

- 400 (bad request): invalid input, either body or path
- 404 (not found): key not found
- 500 (internal error): internal server error

Error responses structure:

```text
{
    "type": ""
    "message": "",
}
```

Where _**type**_ is the error type and _**message**_ the error message.

### Examples

#### Key

Creating/Replacing a key

```json
POST /key/set

BODY:
{
    "key": "",
    "value": ""
}

STATUS:
200: success response

RESPONSE:
{
    "affected_key": true, // whether the key was affected
}
```

Getting a key

```json
POST /key/get

BODY:
{
    "key": "",
}

STATUS:
200: success response
204 (no content): item does not exists

RESPONSE:
{
    "value": ""
}
```

Deleting a key

```json
POST /key/del

BODY:
{
    "key": "",
}

STATUS:
200: success response

RESPONSE:
{
    "affected_key": true, // whether the key was affected
}
```

### Hash

Creating/Replacing a hash

```json
POST /hash/set

BODY:
{
    "key": "",
    "value": {
        "key1": "",
        "key2": ""
    }
}

STATUS:
200: When Item already exists

RESPONSE:
{
    "affected_key": true, // whether the key was affected
}
```

Adding/Replacing the hash properties

```json
POST /hash/add

BODY:
{
    "key": "",
    "value": {
        "key1": "",
        "key2": ""
    }
}

STATUS:
200: Success response
404: key not found

RESPONSE:
{
    "affected_key": true,
}
```

**NOTE:** If the key does not exists, it returns *404* status code

Getting a hash

```json
POST /hash/get

BODY:
{
    "key": "",
}

STATUS:
200: success response
204 (no content): item does not exists

RESPONSE:
{
    "value": {
        "key1": "",
        "key2": ""
    }
}
```

Deleting a hash

```json
POST /hash/del

BODY:
{
    "key": "",
}

STATUS:
200: success response

RESPONSE:
{
    "affected_key": true, // whether the key was affected
}
```
