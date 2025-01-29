<h1 align="center">
  <a href="https://sendex.dev"><img src="/logo.svg" /></a>
</h1>

<h4 align="center">Sendex is a lightweight, file-based tool to make requests to your API</h4>

<h4 align="center"><a href="https://sendex.dev">Go to full documentation</a></h4>

## Contents

- [Key Features](#key-features)
- [Installation](#installation)
- [Basic Usage](#basic-usage)
	- Create a new request
	- Editing your request
	- Running a request
- [Credits](#credits)
- [License](#license)

## Key Features

- File and command based, no GUI
- Lightweight and simple to use
- Can be used for end-to-end API testing
- Keeps your request config part of your API source-code
- Built in Go

## Installation

Install Sendex via Homebrew Tap:

```bash
brew tap hadlow/sendex
brew install sendex
```

Test installation by running:

```bash
sendex help
```

> **Note**
> Currently only tested for MacOS

## Basic Usage

### Create a new request

Requests can be created using Sendex via the `new` command. This command will create a file at the specified file path using a standard template which you can use as a starting point.

```sh
sendex new requests/get-todo.yml
```

This command will create a file at `requests/get-todo.yml` using the default GET template.

### Editing your request

Opening the file will give us this:

```yml
args:
  - id: 1 # specify 1 as default
method: GET
endpoint: http://localhost:8000/blog/{id} # we can use 'id' here
headers:
  - Content-Type: application/json
  - Accept: application/json
whitelist-headers:
  - Content-Type
  - Accept
```

> [Click here](/learn-more/request-configuration) for full explanation of all parameters that can be used.

- The args parameter allows us to pass in command line arguments. Formatted like `id=2`
- Method can be any HTTP method
- Endpoint is the API URL
- Headers must be used in list format
- Whitelist headers keeps the output clean, but only showing the headers listed

If you're just trying out Sendex, feel free to use a test API, such as [JSON Placeholder](https://jsonplaceholder.typicode.com). Replace the default endpoint with `https://jsonplaceholder.typicode.com/todos/{id}`.

### Running a request

Once your request file has been updated for your API, it can then be ran using:

```sh
sendex run request/get-todo.yml id=123
```

This should give the following response:

```sh
200 OK
Content-Type: application/json; charset=utf-8
{
  "userId": 7,
  "id": 123,
  "title": "esse et quis iste est earum aut impedit",
  "completed": false
}
```

## Credits

Sendex uses the following open source software:

- [Golang](https://go.dev)
- [Cobra](https://github.com/spf13/cobra)
- [Pretty](https://github.com/tidwall/pretty)
- [yaml](https://github.com/go-yaml/yaml/tree/v3.0.1)

## License

[MIT](https://github.com/hadlow/sendex/blob/main/LICENSE)

---

