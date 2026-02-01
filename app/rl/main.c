#include "settings.h"

int main(void)
{
    Settings settings = {0};
    
    if (ReadSettings(&settings) != 0)
    {
        return 1;
    }
}