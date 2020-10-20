import program from 'commander';
import Command from './command';
import Run from './run';
import Init from './init';
import Test from './test';

class CLI
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

function main()
{
	let cli = new CLI();
}

main();