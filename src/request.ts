import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import chalk from 'chalk';
import * as YAML from 'yaml';
import File from './file';
import Response from './response';
import getRequestPath from './helpers/getRequestPath';
import parseEnv from './helpers/parseEnv';
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

		let file: File;
		let request: object;

		try
		{
			file = new File(getRequestPath(method, endpoint));
		} catch(e) {
			console.log(chalk.red("There was an error reading that file"));
			return;
		}

		try
		{
			request = YAML.parse(parseEnv(file.read()));
		} catch(e) {
			console.log(chalk.red("There was a YAML error in the file config. Please check."));
			return;
		}

		this.request = this.map(request);
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
			let response = new Response(error.response);

			callback(response);
		});
	}
}
