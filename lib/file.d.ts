/// <reference types="node" />
import * as fs from 'fs';
export default class File {
    private path;
    constructor(file: string);
    getPath(): string;
    exists(): boolean;
    read(): string;
    readYaml(): object;
    write(data: string, callback: () => any): void;
    writeSync(data: string): any;
    writeYaml(data: object, callback: () => any): void;
    writeYamlSync(data: object): any;
    create(contents: string | NodeJS.ArrayBufferView, callback: fs.NoParamCallback): void;
}
