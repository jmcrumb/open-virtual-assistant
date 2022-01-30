import NovaPlugin from '../abstractPlugin';

export default class HelloWorldPlugin extends NovaPlugin {

    getKeywords(): string[] {
        return ['hello', 'hi', 'howdy', 'hello there'];
    }

    execute(command: string): string {
        return 'Hello! My name is Nova.';
    }

    executeSecondaryCommnand(command: string): string {
        return 'To teach me more fun things to do, go to the plugin store.';
    }

}