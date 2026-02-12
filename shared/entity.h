#ifndef ENTITY_H
#define ENTITY_H

typedef struct {
    int hp;
    int id;
    struct {float X, Y;} pos;
    int type;
} Entity;

#endif