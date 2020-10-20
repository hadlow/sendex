"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const commander_1 = __importDefault(require("commander"));
class Command {
    addCommand(command, description) {
        return this.program = commander_1.default.command(command).description(description);
    }
    addAction(action) {
        return this.program.action(action);
    }
}
exports.default = Command;
