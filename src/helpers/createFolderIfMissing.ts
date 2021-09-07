import * as fs from 'fs';
import * as path from 'path';

const createFolderIfMissing = (folder: string): void =>
{
    const absPath = path.join(process.cwd(), folder);

    if(!fs.existsSync(absPath))
        fs.mkdirSync(absPath);
}

export default createFolderIfMissing;