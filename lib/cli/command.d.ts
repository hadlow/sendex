import program from 'commander';
export default abstract class Command {
    private program;
    protected addCommand(command: string, description: string): program.Command;
    protected addAction(action: (any: any) => void): program.Command;
}
