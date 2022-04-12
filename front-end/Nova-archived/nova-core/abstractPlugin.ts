export default abstract class NovaPlugin {
    
    abstract getKeywords(): string[];

    abstract execute(command: string): string;

    executeSecondaryCommnand(command: string): string | undefined {
        // Default behavior: return undefined if plugin does not support secondary commands
        // OR secondary command not recognized
        return undefined;
    };

    requestHistory(): string[] {
        return [];
    }

    toString(): string {
        return this.constructor.name;
    }
    
}