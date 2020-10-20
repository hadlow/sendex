import File from './file';

const defaultConfig = {
	path: '_sendex',
}

export function config(property: string): any
{
	let file: File = new File('.sendex.yml');
	let config = file.readYamlSync()['config'];

	if(!config[property])
		config[property] = defaultConfig[property];

	return config[property];
}

export function env(property: string): any
{
	let file: File = new File('.sendex.yml');

	return file.readYamlSync()['env'][property];
}