import * as fs from 'fs';
import * as path from 'path';
import chalk from 'chalk';
import Command from './command';
import File from '../file';
import { config } from '../config';

export default class Init extends Command
{
	constructor()
	{
		super();

		this.addCommand('init', 'Setup an sendex project');
		this.addAction(this.action.bind(this));
	}

	private action(): void
	{
		const configFile: File = new File('.sendex.yml');

		if(!configFile.exists())
		{
			// Make config file
			configFile.writeYamlSync({
				"config": {
					"path": "_sendex",
					"baseUrl": "http://domain.com/"
				}
			});
		}

		const folders = this.getFolderStructer(config('path'));

		this.generateFolders('', folders);
		this.displaySuccess();
	}

	private displaySuccess(): void
	{
		console.log(chalk.green('Created sendex directory'));
	}

	private getFolderStructer(root: string): object
	{
		return {
			[root]: {
				"responses": {},
				"requests": {},
				"tests": {}
			}
		};
	}

	private generateFolders(root: string, folders: object): void
	{
		root = (root != '' ? root + '/' : root);

		for(let [folder, subfolders] of Object.entries(folders))
		{
			this.createFolderIfMissing(root + folder);

			if(subfolders != {})
				this.generateFolders(folder, subfolders);
		}
	}

	private createFolderIfMissing(folder: string): void
	{
		const absPath = path.join(process.cwd(), folder);

		if(!fs.existsSync(absPath))
			fs.mkdirSync(absPath);
	}
}