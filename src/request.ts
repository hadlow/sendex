import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import File from './file';
import Response from './response';
import getRequestPath from './helpers/getRequestPath';
import { config } from './config';

export default class Request
{
	private method: string;

	private endpoint: string;

	private request: object;

	constructor(method: string, endpoint: string)
	{
		this.method = method;
		this.endpoint = endpoint;

		const file: File = new File(getRequestPath(method, endpoint));
		
		// Map the YAML config to an object for axios
		this.request = this.map(file.readYaml());
	}

	private map(request: object): object
	{
		return {
			method: request['method'],
			baseURL: config('baseUrl'),
			url: request['endpoint'],
			data: request['body'],
			headers: request['headers'],
			params: request['params'],
		};
	}

	public execute(callback: any)
	{
		axios(this.request).then((resp) =>
		{
			let response = new Response(resp);

			callback(response);
		}).catch((error) =>
		{
			console.log(error)
		});
	}
}
