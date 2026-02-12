#include "render.h"
#include "rlassets.h"
#include "settings.h"
#include <dirent.h>

int main(void)
{
    Settings settings = {0};
    
    if (ReadSettings(&settings) != 0)
    {
        return 1;
    }

    InitializeWindow(settings.width, settings.height, settings.fullscreen);
    RLAssets rlassets = {0};
    if (Preload(&rlassets) != 0)
    {
        return 1;
    }
    if (RenderLoop(rlassets, settings) != 0)
    {
        return 1;
    }
    return 0;
}