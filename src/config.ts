import chalk from 'chalk';
import * as YAML from 'yaml';
import File from './file';
import parseEnv from './helpers/parseEnv';

const defaultConfig = {
	path: '_sendex',
	baseUrl: 'http://localhost/',
}

export function config(property: string): any
{
	let file: File;
	let config: string;

	try
	{
		file = new File('.sendex.yml');
	} catch(e) {
		console.log(chalk.red("Error reading config file"));
		return;
	}

	try
	{
		config = YAML.parse(parseEnv(file.read()))['config'];
	} catch(e) {
		console.log(chalk.red("Error reading config file"));
		return;
	}

	if(!config[property])
		config[property] = defaultConfig[property];

	return config[property];
}
