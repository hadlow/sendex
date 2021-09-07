import { config } from '../config';

const getRequestPath = (method: string, endpoint: string): string =>
{
    return `./${config('path')}/requests/${method.toUpperCase()}-${endpoint}.yml`;
};

export default getRequestPath;