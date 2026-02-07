#include "rlassets.h"
#include "assets.h"
#include "raylib.h"
#include <stdlib.h>
#include <string.h>

RLAssets* Preload()
{
    char** list = NULL;
    int size = 0;
    if (GetAssetList("assets", &list, &size) != 0)
    {
        return NULL;
    }
    RLAssets *rlassets = malloc(sizeof(RLAssets));
    if (rlassets == NULL)
    {
        return NULL;
    }
    rlassets->list = list;
    rlassets->size = size;
    rlassets->textures = malloc(sizeof(Texture2D) * size);
    if (rlassets->textures == NULL)
    {
        return NULL;
    }
    for (int i = 0; i < size; i++)
    {
        rlassets->textures[i] = LoadTexture(rlassets->list[i]);
        if (rlassets->textures[i].id <= 0)
        {
            return NULL;
        }
    }
    return rlassets;
}

Texture2D* GetTexture(RLAssets* rlassets, const char* name)
{
    for (int i = 0; i < rlassets->size; i++)
    {
        if (strcmp(rlassets->list[i], name) == 0)
        {
            return &rlassets->textures[i];
        }
    }
    return NULL;
}