import * as fs from 'fs';
import chalk from 'chalk';
import Command from './command';
import createFolderIfMissing from '../helpers/createFolderIfMissing';
import File from '../file';
import { config } from '../config';

export default class Clean extends Command
{
	constructor()
	{
		super();

		this.addCommand('clean', 'Delete all saved responses');
		this.addAction(this.action.bind(this));
	}

	private action(): void
	{
        const directory = `${config('path')}/responses`;
        const gitignore = new File(`${directory}/.gitignore`);

        // @ts-ignore
        fs.rm(directory, { recursive: true });

        createFolderIfMissing(directory);
        
        gitignore.writeSync('*\n!.gitignore');

        console.log(chalk.green('Deleted all saved responses'));
	}
}