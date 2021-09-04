"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const config_1 = require("./config");
const file_1 = __importDefault(require("./file"));
class Response {
    constructor(response) {
        this.data = response.data;
        this.headers = response.headers;
        this.status = response.status;
        this.statusText = response.statusText;
        this.config = response.config;
    }
    save() {
        let file = new file_1.default(config_1.config('path') + '/responses/request.txt');
        let contents = JSON.stringify({
            "Headers": this.headers,
            "Data": this.data,
            "Status": this.status + ' ' + this.statusText,
            "Config": this.config,
        });
        file.create(contents, () => {
        });
        console.log('File saved.');
    }
}
exports.default = Response;
