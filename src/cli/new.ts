import * as fs from 'fs';
import * as path from 'path';
import chalk from 'chalk';
import Command from './command';
import File from '../file';
import getRequestPath from '../helpers/getRequestPath';
import Methods from '../enums/methods.enum';
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
		const errors = this.validate(method, endpoint);

		if(errors.length)
		{
			console.log(chalk.red(`There was ${errors.length} problem${errors.length == 1 ? '' : 's'} creating that request:`));
			
			for(const error of errors)
				console.log(`  * ${error}`);

			return;
		}

		// Remvoe slash at start of endpoint
		while(endpoint.charAt(0) === '/')
			endpoint = endpoint.substring(1);

		const configFile: File = new File(getRequestPath(method, endpoint));

		if(!configFile.exists())
		{
			// Make command configuration file
			configFile.writeYamlSync({
				"method": method.toUpperCase(),
				"endpoint": `/${endpoint}`,
				"params": {},
				"body": "",
				"headers": {
					"Content-Type": "application/json",
					"Accept": "application/json",
				},
			});

			console.log(chalk.green(`Created new request: ${method.toUpperCase()}-${endpoint}.yml`));
		} else {
			console.log(chalk.red(`The command "${method.toUpperCase()} ${endpoint}" already exists.`));
		}
	}

	private validate(method: string, endpoint: string): string[]
	{
		let errors = [];

		// Validate method
		if(method === undefined)
		{
			errors.push(`Missing argument [method].`);
		} else {
			if(!Object.values<string>(Methods).includes(method.toUpperCase()))
				errors.push(`The HTTP method "${method}" is not valid. See https://sendexapi.com/docs/methods for a list of available HTTP methods.`);
		}

		// Validate endpoint
		if(endpoint === undefined)
			errors.push(`Missing argument [endpoint].`);

		return errors;
	}
}
