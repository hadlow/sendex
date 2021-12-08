import program from 'commander';
import Command from './cli/command';
import Init from './cli/init';
import New from './cli/new';
import Run from './cli/run';
import Peek from './cli/peek';
import Clean from './cli/clean';
import Test from './cli/test';

class Main
{
	private commands: Command[] = [];
	
	constructor()
	{
		require('dotenv').config();

		this.addCommand(new Init());
		this.addCommand(new New());
		this.addCommand(new Run());
		this.addCommand(new Peek());
		this.addCommand(new Clean());
		this.addCommand(new Test());

		program.parse(process.argv);
	}

	private addCommand(command: Command): void
	{
		this.commands.push(command);
	}
}

;(async () =>
{
    new Main();
})();
