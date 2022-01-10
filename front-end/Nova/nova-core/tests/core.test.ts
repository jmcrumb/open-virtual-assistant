import NovaCore from "../core";

let core: NovaCore = new NovaCore();

test('Hello World Plugin Functionality', () => {
    expect(core.invoke('Hello')).toBe('Hello! My name is Nova.');
});