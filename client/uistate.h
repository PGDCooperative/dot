#ifndef UISTATE_H
#define UISTATE_H

#include <stdbool.h>

typedef struct {
    bool mainmenu, gameplay, settingsmenu, pausemenu;
} UIState;

#endif