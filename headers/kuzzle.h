#ifndef _KUZZLE_H_
#define _KUZZLE_H_

#include <json/json.h>
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

typedef struct {
    unsigned acknowledged;
    unsigned shardsAcknowledged;
    char error[2048];
} ack_response;

typedef struct {
    unsigned queuable;
    int from;
    int size;
    char scroll[16];
    char scrollId[128];
    char refresh[32];
    char ifExist[32];
    int retryOnConflict;
    json_object *volatiles;
} query_options;

typedef struct {
    char jwt[512];
    char error[2048];
} login_result;

typedef struct {
    json_object *result;
    char error[2048];
} json_result;

extern kuzzle* Kuzzle(char*, char*);
extern char* kuzzle_wrapper_connect();
extern offline_queue* kuzzle_wrapper_get_offline_queue();
extern char* kuzzle_wrapper_get_jwt();
extern int kuzzle_wrapper_check_token(token_validity*, char*);
extern int kuzzle_wrapper_create_index(ack_response*, char*, query_options*);
extern int kuzzle_wrapper_login(login_result*, char*, json_object*, int*);
extern int kuzzle_wrapper_create_my_credentials(json_result*, char*, json_object*, query_options*);

#endif