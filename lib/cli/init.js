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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs = __importStar(require("fs"));
const path = __importStar(require("path"));
const chalk_1 = __importDefault(require("chalk"));
const command_1 = __importDefault(require("./command"));
const file_1 = __importDefault(require("../file"));
const config_1 = require("../config");
class Init extends command_1.default {
    constructor() {
        super();
        this.addCommand('init', 'Setup an sendex project');
        this.addAction(this.action.bind(this));
    }
    action() {
        const configFile = new file_1.default('.sendex.yml');
        if (!configFile.exists()) {
            // Make config file
            configFile.writeYamlSync({
                "config": {
                    "path": "_sendex",
                    "baseUrl": "http://domain.com/"
                }
            });
        }
        const folders = this.getFolderStructer(config_1.config('path'));
        this.generateFolders('', folders);
        this.displaySuccess();
    }
    displaySuccess() {
        console.log(chalk_1.default.blue('Created sendex directory'));
    }
    getFolderStructer(root) {
        return {
            [root]: {
                "out": {},
                "requests": {},
                "tests": {}
            }
        };
    }
    generateFolders(root, folders) {
        root = (root != '' ? root + '/' : root);
        for (let [folder, subfolders] of Object.entries(folders)) {
            this.createFolderIfMissing(root + folder);
            if (subfolders != {})
                this.generateFolders(folder, subfolders);
        }
    }
    createFolderIfMissing(folder) {
        const absPath = path.join(process.cwd(), folder);
        if (!fs.existsSync(absPath))
            fs.mkdirSync(absPath);
    }
}
exports.default = Init;
