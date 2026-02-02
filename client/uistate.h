#ifndef UISTATE_H
#define UISTATE_H

#include <stdbool.h>

typedef struct {
    bool mainmenu, gameloop, settingsmenu, pausemenu;
} UIState;

#endif