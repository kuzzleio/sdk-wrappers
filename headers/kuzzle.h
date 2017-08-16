#ifndef _KUZZLE_H_
#define _KUZZLE_H_

#include <errno.h>

typedef struct {
} kuzzle;

typedef struct {
    char   *query;
    char   *timestamp;
    char   *requestId;
} query_object;

typedef struct {
    query_object** query;
} offline_queue;


typedef struct {
    unsigned valid;
    char state[512];
    int expiresAt;
} token_validity;

extern kuzzle* Kuzzle(char*, char*);
extern char* kuzzle_wrapper_connect();
extern offline_queue* kuzzle_wrapper_get_offline_queue();
extern char* kuzzle_wrapper_get_jwt();
extern int kuzzle_wrapper_check_token(token_validity*, char*);

#endif