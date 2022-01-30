import NovaPlugin from '../abstractPlugin';

export default class CommandNotFoundPlugin extends NovaPlugin {

    executeSecondaryCommnand(command: string): string {
        return 'Can you rephrase what you said?';
    }

    getKeywords(): string[] {
        return [];
    }

    execute(command: string): string {
        return 'I\'m sorry, I don\'t understand.';
    }
}