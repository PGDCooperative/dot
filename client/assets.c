#include "assets.h"
#include "dirent.h"
#include "stdio.h"
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int GetAssetList(char** list, int* size) 
{
    struct dirent *dr;
    DIR *folder = opendir("assets/");
    if (folder == NULL)
    {
        return 1;
    }
    while ((dr = readdir(folder)) != NULL)
    {
        if (!strcmp(dr->d_name, ".") || !strcmp(dr->d_name, ".."))
        {
            (*size)++;
            list = realloc(list, sizeof(int*) * (*size));
            if (list == NULL)
            {
                return 1;
            }
            list[*size] = malloc(sizeof(int) * (strlen(dr->d_name) + 1));
            if (list[*size] == NULL)
            {
                return 1;
            }
            strcpy(list[*size], dr->d_name);
        }
    }
    closedir(folder);
}