import program from 'commander';

export default abstract class Command
{
	private program: program.Command;

	protected addCommand(command: string, description: string): program.Command
	{
		return this.program = program.command(command).description(description);
	}

	protected addAction(action: (any) => void): program.Command
	{
		return this.program.action(action);
	}
}