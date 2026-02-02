#ifndef RENDER_H
#define RENDER_H

#include "settings.h"
#include <stdbool.h>

void InitializeWindow(int width, int height, bool fullscreen);

void RenderLoop(Settings *settings);

#endif