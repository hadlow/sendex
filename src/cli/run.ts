const clear = require('clear');
import Command from './command';
import { config } from '../config';
import { Request } from '../request';

export default class Run extends Command
{
    constructor()
	{
		super();

		this.addCommand('run <endpoint>', 'Make a request to an API');
		this.addAction(this.action);
    }
    
    private action(endpoint: string): void
    {
        let path = '../' + config('path') + '/requests/' + endpoint + '.yml';
        let request = new Request(path);

        request.execute();
    }
}