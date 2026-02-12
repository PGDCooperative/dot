#ifndef GAME_H
#define GAME_H

#include "entity.h"
#include "faction.h"
#include "player.h"
#include "tile.h"
#include <stdbool.h>

typedef struct {
    bool isLobby;
    int lastPlayerID;
    int lastFactionID;
    int lastEntityID;
    Player players[10];
    Faction factions[10];
    Entity entities[500];
    Tile tiles[1024];
} Game;

Game NewGame();

int LoadGame(Game* game, char* savefile);

#endif