import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import File from './file';
import Response from './response';
import { config } from './config';

export class Request
{
	private request: object;

	constructor(path: string)
	{
		let file: File = new File(path);
		
		// Map the YAML config to an object for axios
		this.request = this.map(file.readYaml());
	}

	private map(request: object): object
	{
		return {
			method: request['method'],
			baseURL: config('baseUrl'),
			url: request['url'],
			data: request['body'],
			headers: request['headers'],
		};
	}

	public execute()
	{
		console.log(this.request);
		
		axios(this.request).then((resp) =>
		{
			let response = new Response(resp);

			response.save();
		}).catch((error) =>
		{
			console.log(error)
		});
	}
}