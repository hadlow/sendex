import IResponse from './interfaces/response.interface';
export default class Response {
    data: object;
    headers: object;
    status: number;
    statusText: string;
    config: object;
    constructor(response: IResponse);
    save(): void;
}
