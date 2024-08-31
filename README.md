# Sendex

The simple API tool

## Run a request

`sendex requests/post-user.yml`
`sendex requests/get-user.yml`

## Create a new request

`sendex requests/get-blog.yml --create`

## Pass in args

`sendex requests/get-blog.yml -- id:4371`

## Save response

`sendex requests/get-blog.yml --save`

## Example config file

```yaml
args:
  - id: 1 # specify 1 as default
method: GET
endpoint: http://localhost:8000/blog/{id} # we can use 'id' here
headers:
  - Content-Type: application/json
  - Accept: application/json
```
