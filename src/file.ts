import * as YAML from 'yaml';
import * as fs from 'fs';
import * as path from 'path';

export default class File
{
	private path: string;

	constructor(file: string)
	{
		this.path = path.join(process.cwd(), file);
	}

	public getPath(): string
	{
		return this.path;
	}
	
	public exists(): boolean
	{
		return fs.existsSync(this.path);
	}

	public read(callback): void
	{
		return fs.readFile(this.path, {encoding: 'utf8'}, callback);
	}

	public readSync(): string
	{
		return fs.readFileSync(this.path, {encoding: 'utf8'});
	}

	public readYaml(callback: (any) => any): void
	{
		return this.read((error, data) =>
		{
			if(error)
				console.error(error);

			callback(YAML.parse(data));
		});
	}

	public readYamlSync(): object
	{
		return YAML.parse(this.readSync());
	}

	public write(data: string, callback: (any) => any): void
	{
		return fs.writeFile(this.path, data, callback);
	}

	public writeSync(data: string): any
	{
		return fs.writeFileSync(this.path, data);
	}

	public writeYaml(data: object, callback: (any) => any): void
	{
		return fs.writeFile(this.path, YAML.stringify(data), callback);
	}

	public writeYamlSync(data: object): any
	{
		return fs.writeFileSync(this.path, YAML.stringify(data));
	}

	public create(contents, callback)
	{
		fs.writeFile(this.path, contents, {}, callback);
	}
}