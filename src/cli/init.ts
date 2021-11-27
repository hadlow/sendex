import chalk from 'chalk';
import Command from './command';
import File from '../file';
import createFolderIfMissing from '../helpers/createFolderIfMissing';
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
			try
			{
				// Make config file
				configFile.writeYamlSync({
					"config": {
						"path": "_sendex",
						"baseUrl": "http://domain.com/"
					}
				});
			} catch (e: any) {
				console.log(chalk.red('Error creating config file'));
			}

			try
			{
				const folders = this.getFolderStructer(config('path'));

				this.generateFolders('', folders);
				console.log(chalk.green('Created sendex directory'));
			} catch(e: any) {
				console.log(chalk.red('Error creating _sendex folder'));
			}

			return;
		}

		console.log(chalk.blue('Sendex has already been initialized here'));
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
			createFolderIfMissing(root + folder);

			if(subfolders != {})
				this.generateFolders(folder, subfolders);
		}
	}
}
