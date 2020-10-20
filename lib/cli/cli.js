"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const commander_1 = __importDefault(require("commander"));
const run_1 = __importDefault(require("./run"));
const init_1 = __importDefault(require("./init"));
const test_1 = __importDefault(require("./test"));
class CLI {
    constructor() {
        this.commands = [];
        this.addCommand(new run_1.default());
        this.addCommand(new init_1.default());
        this.addCommand(new test_1.default());
        commander_1.default.parse(process.argv);
    }
    addCommand(command) {
        this.commands.push(command);
    }
}
function main() {
    let cli = new CLI();
}
main();
