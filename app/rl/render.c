#include "render.h"
#include "raylib.h"
#include "gui.h"
#include "uistate.h"
#include <stdbool.h>

void InitializeWindow(int width, int height, bool fullscreen)
{
    InitWindow(width, height, "D.O.T.");
    if (fullscreen)
    {
        ToggleFullscreen();
    }
    SetWindowState(FLAG_WINDOW_RESIZABLE);
}

void RenderLoop(RLAssets* rlassets, Settings* settings)
{
    UIState uistate = {.mainmenu = true};
    Camera3D camera = {0};
    camera.position = (Vector3){ 0.0f, 0.0f, 0.0f };
    camera.target = (Vector3){ 0.0f, 0.0f, 0.0f };
    camera.up = (Vector3){ 0.0f, 0.0f, 0.0f };
    camera.fovy = 45.0f;
    camera.projection = CAMERA_ORTHOGRAPHIC;
    while (!WindowShouldClose())
    {
        BeginDrawing();
        ClearBackground(BLACK);
        BeginMode3D(camera);
        EndMode3D();
        if (uistate.mainmenu)
        {
            MainMenu();
        }
        if (uistate.settingsmenu)
        {
            SettingsMenu(&settings);
        }
        if (uistate.pausemenu)
        {
            PauseMenu();
        }
        EndDrawing();
    }
    CloseWindow();
}