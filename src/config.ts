import File from './file';

const defaultConfig = {
	path: '_sendex',
}

export function config(property: string): any
{
	const file: File = new File('.sendex.yml');
	const config = file.readYaml()['config'];

	if(!config[property])
		config[property] = defaultConfig[property];

	return config[property];
}

export function env(property: string): any
{
	const file: File = new File('.sendex.yml');

	return file.readYaml()['env'][property];
}