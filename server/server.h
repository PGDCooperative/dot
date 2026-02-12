#ifndef SERVER_H
#define SERVER_H

#include <stdbool.h>

int RunServer(char* savefile,
    int port,
    int maxPlayers,
    int adminPassword, 
    int serverPassword,
    bool lan,
    bool autoSave,
    bool autoSaveInterval);

#endif