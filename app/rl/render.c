#include "render.h"
#include "raylib.h"
#include "gui.h"
#include "rl/rlassets.h"
#include "uistate.h"
#include <stdbool.h>
#include <stdio.h>

const int gameWidth = 1920;
const int gameHeight = 1080;

const float gameAspectRatio = (float)gameWidth/gameHeight;

void InitializeWindow(int width, int height, bool fullscreen)
{
    InitWindow(width, height, "D.O.T.");
    if (fullscreen)
    {
        ToggleFullscreen();
    }
    else 
    {
        SetWindowState(FLAG_WINDOW_RESIZABLE);
    }
}

int RenderLoop(RLAssets* rlassets, Settings* settings)
{
    UIState uistate = {.mainmenu = true};
    Camera3D camera = {0};
    camera.position = (Vector3){ 0.0f, 0.0f, 1.0f };
    camera.target = (Vector3){ 0.0f, 0.0f, 0.0f };
    camera.up = (Vector3){ 0.0f, 1.0f, 0.0f };
    camera.fovy = 45.0f;
    camera.projection = CAMERA_ORTHOGRAPHIC;

    Texture2D *backgroundtex = GetTexture(rlassets, "assets/background.png");
    if (backgroundtex == NULL)
    {
        return 1;
    }
    
    while (!WindowShouldClose())
    {
        int screenWidth = GetScreenWidth();
        int screenHeight = GetScreenHeight();
        float screenAspectRatio = (float)screenWidth/screenHeight;

        BeginDrawing();
        ClearBackground(BLACK);
        BeginMode3D(camera);
        if (uistate.mainmenu)
        {
            DrawBillboard(camera, *backgroundtex, 
            (Vector3){.x = 0.0f, .y = 0.0f, .z = 0.0f}, 
            45.0f * GetAspectRatio(screenAspectRatio), WHITE);
        }
        EndMode3D();
        if (uistate.mainmenu)
        {
            MainMenu();
        }
        if (uistate.settingsmenu)
        {
            SettingsMenu(settings);
        }
        if (uistate.pausemenu)
        {
            PauseMenu();
        }
        EndDrawing();
    }
    CloseWindow();
    return 0;
}

float GetAspectRatio(float screenAspectRatio)
{
    if (screenAspectRatio > gameAspectRatio)
    {
        return screenAspectRatio/gameAspectRatio;
    }
    return 1.0f;
}