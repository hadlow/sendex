"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const clear = require('clear');
const command_1 = __importDefault(require("./command"));
const config_1 = require("../config");
const request_1 = require("../request");
class Run extends command_1.default {
    constructor() {
        super();
        this.addCommand('run <endpoint>', 'Make a request to an API');
        this.addAction(this.action);
    }
    action(endpoint) {
        let path = '../' + config_1.config('path') + '/requests/' + endpoint + '.yml';
        let request = new request_1.Request(path);
        request.execute();
    }
}
exports.default = Run;
