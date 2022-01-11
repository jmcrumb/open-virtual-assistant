import NaturalLanguageProcessingAPI from '../api/nlpAPI';
import NovaPlugin from './abstractPlugin';
import HelloWorldPlugin from './HelloWorldPlugin';

export default class NovaCore {
    syntaxTree: {[key: string]: string};

    constructor() {
        this.syntaxTree = {};
        this.initializePlugins();
    }

    initializePlugins() {
        return;
    }

    invoke(input: any) {
        let command: string = NaturalLanguageProcessingAPI.speechToText(input);
        let plugin: NovaPlugin = this.querySyntaxTree(command);
        return plugin.execute(command);
    }

    querySyntaxTree(command: string): NovaPlugin {
        return new HelloWorldPlugin();
    }
}