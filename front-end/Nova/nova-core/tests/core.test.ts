import NovaCore from "../core";
import HelloWorldPlugin from "../plugins/HelloWorldPlugin";

let core: NovaCore = new NovaCore();

beforeEach(() => {
    let core: NovaCore = new NovaCore();
    return core;
});

//Syntax tree

test('Syntax Tree add', () => {
    expect(core.syntaxTree.root['hello']['plugin'] instanceof HelloWorldPlugin).toBe(true);
});

test('Syntax Tree match', () => {
    expect(core.querySyntaxTree('hello') instanceof HelloWorldPlugin).toBe(true);
});

// Keyword recognition

test('Prevent keyword overwrite', () => {
    expect(() => {
        core.syntaxTree.addPlugin(new HelloWorldPlugin());
    }).toThrow('Keyword Conflict. Plugin already exists at the keyword hello');
});

test('Single word keyord recognition', () => {
    expect(core.invoke('Hello')).toBe('Hello! My name is Nova.');
});

test('Multi-word keyord recognition', () => {
    expect(core.invoke('Hello World')).toBe('Hello! My name is Nova.');
});

// Secondary Command functionality

test('Secondary command recognition', () => {
    core.invoke('Hello');
    expect(core.invoke('Nice to meet you')).toBe('To teach me more fun things to do, go to the plugin store.');
});

test('Command not found', () => {
    expect(core.invoke('adfgartha')).toBe('I\'m sorry, I don\'t understand.');
});