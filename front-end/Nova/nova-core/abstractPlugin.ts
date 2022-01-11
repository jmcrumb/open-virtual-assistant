abstract class NovaPlugin {
    
    abstract getKeywords(): string[];

    abstract execute(command: string): string;

    abstract executeSecondaryCommnand(command: string): string;

    requestHistory(): string[] {
        return [];
    }
    
}