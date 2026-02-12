#include "rlassets.h"
#include "assets.h"
#include "raylib.h"
#include <stdlib.h>
#include <string.h>

int Preload(RLAssets *rlassets)
{
    char** list = NULL;
    int size = 0;
    if (GetAssetList("assets", &list, &size) != 0)
    {
        return 1;
    }
    rlassets->list = list;
    rlassets->size = size;
    rlassets->textures = malloc(sizeof(Texture2D) * size);
    if (rlassets->textures == NULL)
    {
        return 1;
    }
    for (int i = 0; i < size; i++)
    {
        rlassets->textures[i] = LoadTexture(rlassets->list[i]);
        if (rlassets->textures[i].id <= 0)
        {
            return 1;
        }
    }
    return 0;
}

Texture2D* GetTexture(RLAssets rlassets, const char* name)
{
    for (int i = 0; i < rlassets.size; i++)
    {
        if (strcmp(rlassets.list[i], name) == 0)
        {
            return &rlassets.textures[i];
        }
    }
    return NULL;
}