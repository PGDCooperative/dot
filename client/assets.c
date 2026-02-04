#include "assets.h"
#include "dirent.h"
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

char** GetAssetList(int* size) 
{
    struct dirent *dr;
    DIR *folder = opendir("assets/");
    if (folder == NULL)
    {
        return NULL;
    }
    char** list = NULL;
    *size = 0;
    while ((dr = readdir(folder)) != NULL)
    {
        if (dr->d_name[0] != '.')
        {
            (*size)++;
            list = (char**)realloc(list, sizeof(char*) * (*size));
            if (list == NULL)
            {
                closedir(folder);
                return NULL;
            }
            list[(*size) - 1] = (char*)malloc(sizeof(char) * (strlen(dr->d_name) + 1));
            if (list[(*size) - 1] == NULL)
            {
                closedir(folder);
                return NULL;
            }
            strcpy(list[(*size) - 1], dr->d_name);
        }
    }
    closedir(folder);
    return list;
}