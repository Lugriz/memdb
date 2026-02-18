# memdb

memdb is a simple key-value in-memory database. Just for learning.

## Features

- in-memory key-value store
- eviction policies like LFU, LRU
- TTL (per key)
- concurrent
- HTTP/CLI access

## Running DB

Running the command

> memdb --port 9300 --eviction-policy LRU

## Data structures

This database will contain some data structures that will store data:

- **Strings**: single key-value
- **Hash**: multiple key-value data
- **Set**: a collection of not repeating values

## HTTP Specification

The path represent the data structure to use, for example: _/key, /hash, /set_ and the HTTP method the action

### Key

Creating/Replacing a key

```text
POST /key

STATUS:
201: When new Item
200: When Item already exists
400: bad input data

BODY:
{
    "key": "",
    "value": ""
}

RESPONSE:
{
    "affected_keys": 0,
}
```

Getting a key

```text
GET /key/{key}

STATUS:
200: success response
204 (no content): item does not exists

RESPONSE (when 200):
{
    "value": ""
}
```

Deleting a key

```text
DELETE /key/{key}

STATUS:
200: success response

RESPONSE:
{
    "affected_keys": 0,
}

```

### Hash

Creating/Replacing a hash

```text
POST /hash

STATUS:
201: When new Item
200: When Item already exists
400: bad input data

BODY:
{
    "key": "",
    "value": {
        "key1": "",
        "key2": ""
    }
}

RESPONSE:
{
    "affected_keys": 1,
}
```

Adding/Replacing the hash properties

```text
PUT /hash/{key}

STATUS:
200: Succes response
400: bad input data
404: key not found

BODY:
{
    "value": {
        "key1": "",
        "key2": ""
    }
}

RESPONSE:
{
    "affected_keys": 1,
}
```

Getting a hash

```text
GET /hash/{key}

STATUS:
200: success response
204 (no content): item does not exists

RESPONSE (when 200):
{
    "value": {
        "key1": "",
        "key2": ""
    }
}
```

Deleting a hash

```text
DELETE /hash/{key}

STATUS:
200: success response

RESPONSE:
{
    "affected_keys": 0,
}

```

Deleting hash properties (can delete multiple keys separated by spaces)

```text
DELETE /hash/{key}?keys={key, ...}

STATUS:
200: success response

RESPONSE:
{
    "affected_keys": 0,
}

```

## CLI Specification

Defines how to use memdb by using the CLI

### Key

Creating/Updating a key

> KEY {keyname} SET {value}

Getting a key

> KEY {keyname} GET

Deleting a key

> KEY {keyname} DEL

### Hash

Creating/Replacing the hash (can set multiple key-values separated by spaces)

> HASH {keyname} SET {key} {value} ...

Adding/Replacing the hash properties (can set multiple key-values separated by spaces)

> HASH {keyname} PUT {key} {value} ...

Getting a hash

> HASH {keyname} GET

Deleting a hash

> HASH {keyname} DEL

Deleting hash keys (can delete multiple keys separated by spaces)

> HASH {keyname} DEL {key} ...
