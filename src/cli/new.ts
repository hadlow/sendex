import * as fs from 'fs';
import * as path from 'path';
import chalk from 'chalk';
import Command from './command';
import File from '../file';
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
				console.log(` * ${error}`);

			return;
		}

		const path = `./${config('path')}/requests/${method.toUpperCase()}-${endpoint}.yml`;
		const configFile: File = new File(path);

		if(!configFile.exists())
		{
			// Make config file
			configFile.writeYamlSync({
				"method": method.toUpperCase(),
				"path": `/${endpoint}`,
				"params": [],
				"body": "",
				"headers": [],
			});
		}

		console.log(chalk.green(`Created new request: ${method.toUpperCase()}-${endpoint}.yml`));
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

		// Validate method
		if(endpoint === undefined)
			errors.push(`Missing argument [endpoint].`);

		return errors;
	}
}