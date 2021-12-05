import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import chalk from 'chalk';
import * as YAML from 'yaml';
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
		const withEnv = this.parseEnv(file.read());

		try
		{
			const request = YAML.parse(withEnv);

			this.request = this.map(request);
		} catch (e) {
			console.log(chalk.red("There was a YAML error in the file config. Please check."));
		}
	}
	
	private parseEnv(contents: string): string
	{
		const regex = /\${(.*?)\}/gmi;
		const matches = contents.match(regex);

		if(!process?.env) return contents;

		for(const match of matches)
		{
			const envVar = match.slice(2, -1);

			if(envVar in process.env)
				contents = contents.replace(match, process.env[envVar]);
		}

		return contents;
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
