#include "render.h"
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
    RenderLoop(&settings);
}