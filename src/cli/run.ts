const clear = require('clear');
import chalk from 'chalk';
import Command from './command';
import File from '../file';
import Request from '../request';
import getRequestPath from '../helpers/getRequestPath';

export default class Run extends Command
{
    constructor()
	{
		super();

		this.addCommand('run [method] [endpoint]', 'Execute a request');
		this.addAction(this.action.bind(this));
    }
    
    private action(method: string, endpoint: string): void
    {
        const errors = this.validate(method, endpoint);

        if(errors.length)
		{
			console.log(chalk.red(`There was ${errors.length} problem${errors.length == 1 ? '' : 's'} executing that request:`));
			
			for(const error of errors)
				console.log(`  * ${error}`);

			return;
		}

		const configFile: File = new File(getRequestPath(method, endpoint));

		if(!configFile.exists())
		{
			console.log(chalk.red(`The command "${method.toUpperCase()} ${endpoint}" does not exist.`));
		} else {
			const request = new Request(method, endpoint);

            request.execute();
		}
    }

    private validate(method: string, endpoint: string): string[]
	{
		let errors = [];

		// Validate method
		if(method === undefined)
			errors.push(`Missing argument [method].`);

		// Validate endpoint
		if(endpoint === undefined)
			errors.push(`Missing argument [endpoint].`);

		return errors;
	}
}