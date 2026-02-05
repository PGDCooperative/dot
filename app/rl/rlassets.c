#include "rlassets.h"
#include "assets.h"
#include "raylib.h"
#include <stdlib.h>

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
    }
    return 0;
}