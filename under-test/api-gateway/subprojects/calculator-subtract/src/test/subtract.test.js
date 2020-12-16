const { subtract } = require('../main/subtract');

describe('The subtraction function', () => {
    describe('two numbers are passed', () => {
       it('returns the correct answer', () => {
           expect(subtract(2, 2)).toBe(0);
           expect(subtract(2, 4)).toBe(-2);
           expect(subtract(6, 2)).toBe(4);
           expect(subtract(2, 8)).toBe(-6);
           expect(subtract(2, 200)).toBe(-198);
           expect(subtract(200, -211)).toBe(411);
        })
    });

    describe('when two strings containing numbers are passed', () => {
        expect(subtract('2', '2')).toBe(0);
        expect(subtract('2', '4')).toBe(-2);
        expect(subtract('6', '2')).toBe(4);
        expect(subtract('2', '8')).toBe(-6);
        expect(subtract('2', '200')).toBe(-198);
        expect(subtract('200', '-211')).toBe(411);
    });

    describe('when two strings containing numbers are passed', () => {
        expect(subtract('2e1', '2')).toBe(18);
        expect(subtract('2e0', '4')).toBe(-2);
        expect(subtract('6e0', '2')).toBe(4);
        expect(subtract('2e0', '8')).toBe(-6);
        expect(subtract('2e0', '200')).toBe(-198);
        expect(subtract('200e0', '-211')).toBe(411);
    });

    describe('any non numeric expect number strings are passed', () => {
        it('returns undefined', () => {
            expect(subtract(2, 'a')).toBeUndefined();
            expect(subtract(2, undefined)).toBeUndefined();
            expect(subtract(undefined, 'a')).toBeUndefined();
            expect(subtract(2, null)).toBeUndefined();
            expect(subtract(2, 'kwmekooq')).toBeUndefined();
            expect(subtract('fwefwe', 'feww')).toBeUndefined();
        })
    })
});
