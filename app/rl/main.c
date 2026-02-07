#include "render.h"
#include "rlassets.h"
#include "settings.h"
#include <dirent.h>
#include "stdio.h"

int main(void)
{
    Settings settings = {0};
    
    if (ReadSettings(&settings) != 0)
    {
        return 1;
    }

    InitializeWindow(settings.width, settings.height, settings.fullscreen);
    RLAssets* rlassets = Preload();
    if (rlassets == NULL)
    {
        return 1;
    }
    RenderLoop(rlassets, &settings);
}