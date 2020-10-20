"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Request = void 0;
const axios_1 = __importDefault(require("axios"));
const file_1 = __importDefault(require("./file"));
const response_1 = __importDefault(require("./response"));
const config_1 = require("./config");
class Request {
    constructor(path) {
        let file = new file_1.default(path);
        this.request = this.map(file.readYamlSync());
    }
    map(request) {
        return {
            method: request['method'],
            baseURL: config_1.config('baseUrl'),
            url: request['url'],
            data: request['body'],
            headers: request['headers'],
        };
    }
    execute() {
        console.log(this.request);
        axios_1.default(this.request).then((resp) => {
            let response = new response_1.default(resp);
            response.save();
        }).catch((error) => {
            console.log(error);
        });
    }
}
exports.Request = Request;
