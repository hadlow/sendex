import * as fs from 'fs';
import Command from './command';
import File from '../file';
import { config } from '../config';

export default class Test extends Command
{
	private tests: string;

	constructor()
	{
		super();

		this.addCommand('test [endpoint]', 'Test an API enpoint');
	}

	private action(path: string)
	{
		fs.readdirSync(config('path') + '/tests/').forEach(file => {
			console.log(file)
		});

		let file = new File(path);

		this.tests = file.read();
	}
}