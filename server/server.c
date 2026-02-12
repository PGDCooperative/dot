#include "server.h"
#include "game.h"
#include <stdio.h>

int RunServer(char* savefile,
    int port,
    int maxPlayers,
    int adminPassword, 
    int serverPassword,
    bool lan,
    bool autoSave,
    bool autoSaveInterval)
{
    Game game;
    if (savefile == NULL)
    {
        game = NewGame();
    }
    else 
    {
        if (LoadGame(&game, savefile) != 0)
        {
            return 1;
        }
    }
    return 0;
}