#ifndef RENDER_H
#define RENDER_H

#include "rlassets.h"
#include "settings.h"
#include <stdbool.h>

void InitializeWindow(int width, int height, bool fullscreen);

void RenderLoop(RLAssets* rlassets, Settings* settings);

#endif