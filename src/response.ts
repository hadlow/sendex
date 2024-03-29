import chalk from 'chalk';
import File from './file';
import IResponse from './interfaces/response.interface';
import { config } from './config';
import getRequestPath from './helpers/getRequestPath';

export default class Response
{
	public data: object;

	public headers: object;

	public status: number;

	public statusText: string;

	public config: object;

	constructor(response: IResponse)
	{
		if(!response)
		{
			this.status = 999;
			return;
		}

		this.data = response.data;
		this.headers = response.headers;
		this.status = response.status;
		this.statusText = response.statusText;
		this.config = response.config;
	}

	public save(method: string, endpoint: string)
	{
		const d = new Date();
		const date = `${d.getFullYear()}${("0" + (d.getMonth() + 1)).slice(-2)}${("0" + d.getDate()).slice(-2)}`;
		const time = `${("0" + d.getHours()).slice(-2)}${("0" + d.getMinutes()).slice(-2)}${("0" + d.getSeconds()).slice(-2)}`;
		const datetime = `${date}_${time}`;
		endpoint = endpoint.replace('/', '-');
		const file: File = new File(`${config('path')}/responses/${datetime}_${method.toUpperCase()}-${endpoint}.txt`);

		const contents = `${method.toUpperCase()} /${endpoint}
${this.status + ' ' + this.statusText}
---
Headers
${JSON.stringify(this.headers, null, 4)}
---
Body
${JSON.stringify(this.data, null, 4)}
`

		file.writeSync(contents);

		console.log(chalk.cyan('Status') + chalk.red(': ') + chalk.green(`${this.status} ${this.statusText}`));
		console.log(`Response saved at ${chalk.underline(file.getPath())}`);
	}

	public print()
	{
		if(this.status == 999)
		{
			console.log(chalk.red("Connection closed"));
			return;
		}

		console.log(`${JSON.stringify(this.data, null, 4)}`);

		let color = chalk.red;

		switch(Math.floor(this.status / 100))
		{
			case 2:
				color = chalk.green;
				break;

			case 3:
				color = chalk.cyan;
				break;

			case 4:
				color = chalk.yellow;
				break;

			case 5:
				color = chalk.red;
				break;
		}

		console.log(color(`${this.status} ${this.statusText}`));
	}
}
