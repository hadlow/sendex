"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.env = exports.config = void 0;
const file_1 = __importDefault(require("./file"));
const defaultConfig = {
    path: '_sendex',
};
function config(property) {
    const file = new file_1.default('.sendex.yml');
    const config = file.readYaml()['config'];
    if (!config[property])
        config[property] = defaultConfig[property];
    return config[property];
}
exports.config = config;
function env(property) {
    const file = new file_1.default('.sendex.yml');
    return file.readYaml()['env'][property];
}
exports.env = env;
