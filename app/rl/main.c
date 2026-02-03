#include "assets.h"
#include "render.h"
#include "settings.h"
#include <dirent.h>
#include <stdio.h>

int main(void)
{
    //Settings settings = {0};
    
    //if (ReadSettings(&settings) != 0)
    //{
    //    return 1;
    //}
    char** test = NULL;
    int size = 0;
    if (GetAssetList(test, &size) != 0)
    {
        return 1;  
    }
    //printf("%s", test[0]);
    //InitializeWindow(settings.width, settings.height, settings.fullscreen);
    //RenderLoop(&settings);
}