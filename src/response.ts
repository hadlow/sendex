import { config } from './config';
import File from './file';
import IResponse from './interfaces/response.interface';

export default class Response
{
	public data: object;

	public headers: object;

	public status: number;

	public statusText: string;

	public config: object;

	constructor(response: IResponse)
	{
		this.data = response.data;
		this.headers = response.headers;
		this.status = response.status;
		this.statusText = response.statusText;
		this.config = response.config;
	}

	public save()
	{
		let file: File = new File(config('path') + '/out/request.txt');

		let contents = JSON.stringify({
			"Headers": this.headers,
			"Data": this.data,
			"Status": this.status + ' ' + this.statusText,
			"Config": this.config,
		});

		file.create(contents, () => {

		});

		console.log('File saved.');
	}
}