import { config } from '../config';

const getRequestPath = (method: string, endpoint: string): string =>
{
	endpoint = endpoint.replace('/', '-');

    return `./${config('path')}/requests/${method.toUpperCase()}-${endpoint}.yml`;
};

export default getRequestPath;
