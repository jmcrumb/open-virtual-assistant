import NaturalLanguageProcessingAPI from '../api/nlpAPI';
import NovaPlugin from './abstractPlugin';
import CommandNotFoundPlugin from './plugins/CommandNotFoundPlugin';
import HelloWorldPlugin from './plugins/HelloWorldPlugin';

export default class NovaCore {
    syntaxTree: SyntaxTree;
    plugins: NovaPlugin[];
    currentPlugin: NovaPlugin | null | undefined;

    constructor() {
        // TODO: Replace mock data with actual data.
        this.plugins = [
            new HelloWorldPlugin(),
        ];
        this.syntaxTree = new SyntaxTree(new CommandNotFoundPlugin());
        this.initializePlugins();
    }

    initializePlugins() {
        this.plugins.forEach((plugin) => {
            this.syntaxTree.addPlugin(plugin);
        });
    }

    invoke(input: any): any {
        let command: string = NaturalLanguageProcessingAPI.speechToText(input).toLowerCase();
        this.currentPlugin = this.querySyntaxTree(command);
        return NaturalLanguageProcessingAPI.textToSpeech(this.currentPlugin.execute(command));
    }

    querySyntaxTree(command: string): NovaPlugin {
        let tokenizedCommand: string[] = command.split(' ');
        return this.syntaxTree.matchCommand(command);
    }
}

class SyntaxTree {
    root: {[key: string]: any};
    notFound: NovaPlugin;
    hits: any;

    constructor(commandNotFoundPlugin: NovaPlugin) {
        this.root = {};
        this.notFound = commandNotFoundPlugin;
    }

    addPlugin(plugin: NovaPlugin) {
        let keywords: string[] = plugin.getKeywords();
        let branch = this.root;
        for(let keyword of keywords) {
            let tokenizedKeywords: string[] = keyword.split(' ')
            for(let i = 0; i < tokenizedKeywords.length - 1; i++) {
                branch = branch[tokenizedKeywords[i]];
            }
            // TODO: Catch this error at the UI layer
            if('plugin' in branch[tokenizedKeywords[tokenizedKeywords.length - 1]]) {
                throw new Error(`Keyword Conflict. Plugin already exists at the keyword ${keyword}`);
            }
            branch[tokenizedKeywords[tokenizedKeywords.length - 1]] = {'plugin': plugin};
        }
    }
    
    matchCommand(command: string): NovaPlugin {
        let tokenizedCommand: string[] = command.split(' ');
        let branch = this.root;
        let max: number = 0;
        this.hits = new Array();

        for(let i = 0; i < tokenizedCommand.length; i++) {
            if(tokenizedCommand[i] in branch) {
                let depth: number = this.matchCommandRec(branch, tokenizedCommand, 0);
                if(depth > max) max = depth;
            }
        }
        // TODO: expand this conditional to check for secondary keywords when implementing
        return this.hits.length > 0 ? this.hits[max][0] : this.notFound;
    }
    
    matchCommandRec(branch: {[key: string]: any}, tokenizedCommand: string[], depth: number): number {
        if(tokenizedCommand.length > 0 && tokenizedCommand[0] in  branch) {
            return this.matchCommandRec(branch[tokenizedCommand[0]], tokenizedCommand.slice(1), depth + 1);
        } else {
            if(this.hits.length <= depth) this.hits.length = depth * 2 + 1;
            if(this.hits[depth] == undefined) this.hits[depth] = new Array();
            this.hits[depth].push(branch['plugin']);
            return depth;
        }
    }
}