const clear = require('clear');
import Command from './command';
import { config } from '../config';
import { Request } from '../request';

export default class Run extends Command
{
    constructor()
	{
		super();

		this.addCommand('run [method] [endpoint]', 'Make a request to an API');
		this.addAction(this.action.bind(this));
    }
    
    private action(method: string, endpoint: string): void
    {
        const path = `./${config('path')}/requests/${method.toUpperCase()}-${endpoint}.yml`;
        const request = new Request(path);

        request.execute();
    }
}