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

Once setup, you can start adding requests to the `_sendex` folder. These requests live in the `requests` folder. You can then make that request by running (replacing <REQUEST> with your request file name, excluding the file extension):

```bash
sendex run <REQUEST>
```

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