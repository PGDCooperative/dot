#ifndef GAME_H
#define GAME_H

#include "faction.h"
#include "player.h"
#include "tile.h"

typedef struct {
    int lastPlayerID;
    int lastFactionID;
    int lastUnitID;
    Player players[8];
    Faction factions[8];
    Tile *tiles;
} Game;

#endif