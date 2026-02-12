#ifndef RLASSETS_H
#define RLASSETS_H

#include "raylib.h"

typedef struct {
    char** list;
    Texture2D* textures;
    int size;
} RLAssets;

int Preload(RLAssets *rlassets);

Texture2D* GetTexture(RLAssets rlassets, const char* name);

#endif