export default class File {
    private path;
    constructor(file: string);
    getPath(): string;
    exists(): boolean;
    read(callback: any): void;
    readSync(): string;
    readYaml(callback: (any: any) => any): void;
    readYamlSync(): object;
    write(data: string, callback: (any: any) => any): void;
    writeSync(data: string): any;
    writeYaml(data: object, callback: (any: any) => any): void;
    writeYamlSync(data: object): any;
    create(contents: any, callback: any): void;
}
