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
        if (strcmp(dr->d_name, ".") != 0 || strcmp(dr->d_name, "..") != 0)
        {
            printf("%s\n", "test");
            (*size)++;
            list = realloc(list, sizeof(char*) * (*size));
            if (list == NULL)
            {
                return 1;
            }
            list[*size - 1] = malloc(sizeof(char) * (strlen(dr->d_name) + 1));
            if (list[*size - 1] == NULL)
            {
                return 1;
            }
            strcpy(list[*size - 1], dr->d_name);
        }
    }
    closedir(folder);
    return 0;
}