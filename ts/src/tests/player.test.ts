import { Player } from '../models/player';

describe('Player class', () => {
  it('should create a player with the correct name and country', () => {
    const player = new Player('Tiger Woods', '🇺🇸', 9, {});

    expect(player.name).toBe('Tiger Woods');
    expect(player.country).toBe('🇺🇸');
    expect(player.qp).toBe(9);
  });
});