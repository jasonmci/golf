import { rollDice } from '../utils/dice';

describe('rollDice function', () => {
  it('should return a number between 1 and 6', () => {
    const result = rollDice();
    expect(result).toBeGreaterThanOrEqual(1);
    expect(result).toBeLessThanOrEqual(6);
  });
  
  it('should return an integer', () => {
    const result = rollDice();
    expect(Number.isInteger(result)).toBe(true);
  });
});
