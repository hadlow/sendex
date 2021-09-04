"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const YAML = __importStar(require("yaml"));
const fs = __importStar(require("fs"));
const path = __importStar(require("path"));
class File {
    constructor(file) {
        this.path = path.join(process.cwd(), file);
    }
    getPath() {
        return this.path;
    }
    exists() {
        return fs.existsSync(this.path);
    }
    read() {
        return fs.readFileSync(this.path, { encoding: 'utf8' });
    }
    readYaml() {
        return YAML.parse(this.read());
    }
    write(data, callback) {
        return fs.writeFile(this.path, data, callback);
    }
    writeSync(data) {
        return fs.writeFileSync(this.path, data);
    }
    writeYaml(data, callback) {
        return fs.writeFile(this.path, YAML.stringify(data), callback);
    }
    writeYamlSync(data) {
        return fs.writeFileSync(this.path, YAML.stringify(data));
    }
    create(contents, callback) {
        fs.writeFile(this.path, contents, {}, callback);
    }
}
exports.default = File;
