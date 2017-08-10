#ifndef _KUZZLE_H_
#define _KUZZLE_H_

#define bool char
#define TRUE '1'
#define FALSE '0'

typedef struct {
} kuzzle;

typedef struct {
    char*   query;
    char*   timestamp;
    char*   requestId;
} query_object;

typedef struct {
    query_object** query;
} offline_queue;

typedef struct {
    bool valid;
    char* state;
    int expiresAt;
} token_validity;

extern char* kuzzle_wrapper_connect();
extern offline_queue* kuzzle_wrapper_get_offline_queue();
extern char* kuzzle_wrapper_get_jwt();
extern token_validity kuzzle_wrapper_check_token(char*);

#endif