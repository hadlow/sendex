# sendex

A lightweight API endpoint testing tool

## Installation

Install globally with NPM:

```bash
npm install -g sendex
```

## Usage

First you will need to initialize sendex in your project:

```bash
sendex init
```

This command will create the .sendex.yml file, which contains the global configuration sendex, and an empty set of sendex folders that are used to keep each request, response and test. The directory stucture is as follows:

- **requests**: This will store the config files for each requests. A typical config file will contain the path, method and headers for that request.
- **responses**: Responses from each request will be stored here.
- **tests**: sendex gives you the ability to write JavaScript tests for each request, those tests go here.

To create your first request, use the following command:

```bash
sendex new get posts
```

This command is made up of three parts (excluding the sendex command). The `new` command specifies that we are creating a new request. The new command takes 2 arguments: the first argument is the method that we are using for the request, for this example we are using the GET method. The method can be written uppercase or lowercase. The next command is the path, relative to the URL. If we use `posts` like in the above example, the path that will be tested is `/posts`. If we wanted to send a request to the path `/posts/all` then we would use the command `sendex new get posts/all` for example.

```bash
sendex run get posts
```

To run that command, we can then use the same stucture, but with the run command instead.

You can also run tests on a request. Create a test in the `tests` folder and run it using (replacing <TEST> with the file name of your test, excluding the file extension):

```bash
sendex test <TEST>
```

You can run all tests with:

```bash
sendex test
```

## License
[MIT](https://choosealicense.com/licenses/mit/)