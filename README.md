# Sendex

Sendex is a lightweight, file-based tool to make requests to your API

## Installation

Coming soon

## Usage

### Create a new request with default template

`sendex new requests/get-blog.yml`

### Create a new request with the post template

`sendex new requests/post-user.yml --post`

### Run a request

`sendex run requests/post-user.yml`
`sendex run requests/get-user.yml`

### Pass in args

`sendex run requests/get-blog.yml id=4371`

### Show specific fields

`sendex run requests/get-blog.yml --status`
`sendex run requests/get-blog.yml --body`
`sendex run requests/get-blog.yml --headers`

### Save response

`sendex save requests/get-blog.yml`

### Save response and specify output file path

`sendex save requests/get-blog.yml --output=responses/get-blog.txt`

### Args also work on save

`sendex save requests/get-blog.yml id=4371 --output=responses/get-blog.txt`

### Example config file

```yaml
args:
  - id: 1 # specify 1 as default
method: GET
endpoint: http://localhost:8000/blog/{id} # we can use 'id' here
headers:
  - Content-Type: application/json
  - Accept: application/json
whitelist-headers: # only show these headers in the output
  - Content-Type
  - Etag
```
