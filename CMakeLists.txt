cmake_minimum_required(VERSION 3.25)

project(dot)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

set(CMAKE_C_STANDARD 99)

add_executable(dot 
    ${CMAKE_SOURCE_DIR}/app/rl/main.c
    ${CMAKE_SOURCE_DIR}/app/rl/rlassets.c
    ${CMAKE_SOURCE_DIR}/app/rl/gui.c
    ${CMAKE_SOURCE_DIR}/app/rl/render.c
    ${CMAKE_SOURCE_DIR}/app/rl/input.c

    ${CMAKE_SOURCE_DIR}/client/assets.c
    ${CMAKE_SOURCE_DIR}/client/settings.c

    ${CMAKE_SOURCE_DIR}/shared/game.c
    ${CMAKE_SOURCE_DIR}/shared/tile.c

    ${CMAKE_SOURCE_DIR}/utils/file.c
    ${CMAKE_SOURCE_DIR}/deps/cJSON/cJSON.c
)

add_subdirectory(deps/raylib)

target_link_libraries(dot PRIVATE raylib)

include_directories(
    ${CMAKE_SOURCE_DIR}/app
    ${CMAKE_SOURCE_DIR}/client
    ${CMAKE_SOURCE_DIR}/server
    ${CMAKE_SOURCE_DIR}/shared
    ${CMAKE_SOURCE_DIR}/utils
    ${CMAKE_SOURCE_DIR}/deps/raylib/src
    ${CMAKE_SOURCE_DIR}/deps/cJSON
)