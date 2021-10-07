import chalk from 'chalk';
import File from './file';
import IResponse from './interfaces/response.interface';
import { config } from './config';

export default class Response
{
	public data: object;

	public headers: object;

	public status: number;

	public statusText: string;

	public config: object;

	constructor(response: IResponse)
	{
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
		const file: File = new File(`${config('path')}/responses/${datetime}_${method.toUpperCase()}-${endpoint}.json`);

		const contents = JSON.stringify({
			"Headers": this.headers,
			"Data": this.data,
			"Status": this.status + ' ' + this.statusText,
			"Config": this.config,
		}, null, 4);

		file.writeSync(contents);

		console.log(chalk.cyan('Status') + chalk.red(': ') + chalk.green(`${this.status} ${this.statusText}`));
		console.log(`Response saved at ${chalk.underline(file.getPath())}`);
	}

	public print()
	{
		const message = {
			'Status': `${this.status} ${this.statusText}`,
			'Headers': `${JSON.stringify(this.headers, null, 4)}`,
			'Data': `${JSON.stringify(this.data, null, 4)}`,
		}

		for(const label in message)
		{
			console.log(chalk.cyan(label) + chalk.red(': ') + chalk.green(`${message[label]}`));
		}
	}
}
