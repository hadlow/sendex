import * as fs from 'fs';
import * as path from 'path';
import chalk from 'chalk';
import Command from './command';
import File from '../file';
import { config } from '../config';



export default class New extends Command
{
	constructor()
	{
		super();

		this.addCommand('new [method] [endpoint]', 'Create a new request');
		this.addAction(this.action.bind(this));
	}

	private action(method: string, endpoint: string): void
	{
		const path = `./${config('path')}/requests/${method.toUpperCase()}-${endpoint}.yml`;
		const configFile: File = new File(path);

		if(!configFile.exists())
		{
			// Make config file
			configFile.writeYamlSync({
				"method": method.toUpperCase(),
				"path": `/${endpoint}`,
				"body": "",
				"headers": "",
			});
		}

		this.displaySuccess(endpoint);
	}

	private displaySuccess(endpoint: string): void
	{
		console.log(chalk.green(`Created new request: ${endpoint}`));
	}
}