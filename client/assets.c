#include "assets.h"
#include "dirent.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <unistd.h>

int GetAssetList(const char* p, char*** list, int* size) 
{
    struct dirent *dr;
    DIR *folder = opendir(p);
    if (folder == NULL)
    {
        return 1;
    }
    while ((dr = readdir(folder)) != NULL)
    {
        if (dr->d_name[0] != '.')
        {
            int strsize = strlen(p) + strlen("/") + strlen(dr->d_name) + 1;
            char buf[strsize];
            if (sprintf(buf, "%s/%s", p, dr->d_name) != strsize - 1)
            {
                closedir(folder);
                return 1;
            }
            if (dr->d_type == DT_DIR)
            {
                if (GetAssetList(buf, list, size) != 0)
                {
                    closedir(folder);
                    return 1;
                }
            }
            if (dr->d_type == DT_REG)
            {
                (*size)++;
                *list = (char**)realloc(*list, sizeof(char*) * (*size));
                if (*list == NULL)
                {
                    closedir(folder);
                    return 1;
                }
                (*list)[(*size) - 1] = (char*)malloc(sizeof(char) * strsize);
                if ((*list)[*size - 1] == NULL)
                {
                    free(*list);
                    closedir(folder);
                    return 1;
                }
                strcpy((*list)[(*size) - 1], buf);
            }
        }
    }
    closedir(folder);
    return 0;
}