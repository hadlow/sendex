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

	public read(): string
	{
		return fs.readFileSync(this.path, {encoding: 'utf8'});
	}

	public readYaml(): object
	{
		return YAML.parse(this.read());
	}

	public write(data: string, callback: () => any): void
	{
		return fs.writeFile(this.path, data, callback);
	}

	public writeSync(data: string): any
	{
		return fs.writeFileSync(this.path, data);
	}

	public writeYaml(data: object, callback: () => any): void
	{
		return fs.writeFile(this.path, YAML.stringify(data), callback);
	}

	public writeYamlSync(data: object): any
	{
		return fs.writeFileSync(this.path, YAML.stringify(data));
	}

	public create(contents: string | NodeJS.ArrayBufferView, callback: fs.NoParamCallback)
	{
		fs.writeFile(this.path, contents, {}, callback);
	}
}
