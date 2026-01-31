#ifndef SETTINGS_H
#define SETTINGS_H

#include <stdbool.h>

typedef struct {
    int width, height;
    bool fullscreen;
} Settings;

int ReadSettings(Settings *settings);

#endif