import program from 'commander';
import Command from './cli/command';
import Run from './cli/run';
import Init from './cli/init';
import Test from './cli/test';

class Main
{
	private commands: Command[] = [];
	
	constructor()
	{
		this.addCommand(new Run());
		this.addCommand(new Init());
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