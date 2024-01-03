import re
from dataclasses import dataclass
from queue import PriorityQueue

@dataclass
class Character:
    hp: int
    damage: int = 0
    armor: int = 0
    mana: int = 0
    
    def copy(self):
        return Character(self.hp, self.damage, self.armor, self.mana)

@dataclass
class Spell:
    name: str
    cost: int
    damage: int = 0
    timer: int = 0
    armor: int = 0
    heal: int = 0
    mana: int = 0
    
    def copy(self):
        return Spell(self.name, self.cost, self.damage, self.timer, self.armor, self.heal, self.mana)
    
    def apply_effect(self, player: Character, boss: Character) -> (Character, Character):
        boss.hp -= self.damage
        player.hp += self.heal
        player.mana += self.mana
        return player, boss

@dataclass(eq=True)
class GameState:
    player: Character
    boss: Character
    mana_spent: int
    effects: list
    
    def __lt__(self, other):
        return self.mana_spent < other.mana_spent
    
@dataclass
class Game:
    Player: Character
    Boss_data: str
    Spells = [
        Spell("Magic Missile", 53, 4, 1),
        Spell("Drain", 73, 2, 1, heal= 2),
        Spell("Shield", 113, 0, 6, 7),
        Spell("Poison", 173, 3, 6),
        Spell("Recharge", 229, 0, 5, mana= 101)
    ]
    
    def __post_init__(self):
        self.lowest_mana = 1000000 # Arbitrary large number
        hp, damage = re.findall(r"\d+", self.Boss_data)
        self.Boss = Character(int(hp), int(damage))
    
    # We simulate a turn in the game. We invalidate the game if the player dies, runs out of mana or 
    # uses more mana than the lowest found mana usage.
    def turn(self, game_state: GameState, spell: Spell, hard_mode=False) -> (GameState, int):
        new_effects = []
        p = game_state.player.copy()
        boss = game_state.boss.copy()
        
        if hard_mode:
            p.hp -= 1
            if p.hp < 1:
                return
            
        # Turn start
        for effect in game_state.effects:
            if not effect.timer:
                p.armor -= effect.armor
                continue
            p, boss = effect.apply_effect(p, boss)
            new_effects.append(effect.copy())
            new_effects[-1].timer -= 1
            
        if boss.hp < 1:
            return game_state.mana_spent
        
        # Player turn
        for effect in new_effects:
            if effect.timer and (spell.name == effect.name):
                return # Spell already active
            
        p.mana -= spell.cost
        p.armor += spell.armor
        
        if p.mana < 0:
            return # Not enough mana, invalid state
        mama_spent = game_state.mana_spent + spell.cost
        
        if mama_spent > self.lowest_mana:
            return # No point in continuing, already spending more than the lowest mana usage
        new_effects.append(spell)
        
        # Boss turn
        newer_effects = []
        for effect in new_effects:
            if not effect.timer:
                p.armor -= effect.armor
                continue
            p, boss = effect.apply_effect(p, boss)
            newer_effects.append(effect.copy())
            newer_effects[-1].timer -= 1
        if boss.hp < 1:
            return mama_spent
        p.hp -= max(1, boss.damage - p.armor)
        
        if p.hp < 1:
            return # Player died, invalid state
        
        new_effects = [effect for effect in newer_effects if effect.timer]
        return GameState(p, boss, mama_spent, newer_effects)
    
    # The idea is to keep track of game states, we use a priority queue to keep track of the lowest mana usage.
    # We use a negative mana usage to make sure the priority queue sorts the lowest mana usage first
    # We check the games states in a breadth first manner, so we can be sure that the first time we find a game state 
    # where the boss dies, it is the lowest mana usage.
    def find_lowest_mana_use(self, hard_mode=False):
        game_states = PriorityQueue()
        game_states.put((0, GameState(self.Player, self.Boss, 0, [])))
        
        while not game_states.empty():
            _, game_state = game_states.get()
            for spell in self.Spells:
                new_game_state = self.turn(game_state, spell, hard_mode)
                if isinstance(new_game_state, int):
                    self.lowest_mana = min(self.lowest_mana, new_game_state)
                elif isinstance(new_game_state, GameState):
                    game_states.put((-new_game_state.mana_spent, new_game_state))
                    
        return self.lowest_mana