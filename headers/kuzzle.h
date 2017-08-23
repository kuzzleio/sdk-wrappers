#ifndef _KUZZLE_H_
#define _KUZZLE_H_

#include <json/json.h>
#include <errno.h>

typedef struct {
} kuzzle;

//query object used by query()
typedef struct {
    char   *query;
    char   timestamp[11];
    char   requestId[36];
} query_object;

typedef struct {
    query_object** query;
} offline_queue;

//response of check_token()
typedef struct {
    unsigned valid;
    char state[512];
    int expiresAt;
    char error[2048];
} token_validity;

//response for any delete* function
typedef struct {
    unsigned acknowledged;
    unsigned shardsAcknowledged;
    char error[2048];
} ack_response;

//options passed to query()
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

//result of login()
typedef struct {
    char jwt[512];
    char error[2048];
} login_result;

//any json result
typedef struct {
    json_object *result;
    char error[2048];
} json_result;

typedef struct {
    unsigned result;
    char error[2048];
} bool_result;

extern kuzzle* Kuzzle(char*, char*);
extern char* kuzzle_wrapper_connect();
extern offline_queue* kuzzle_wrapper_get_offline_queue();
extern char* kuzzle_wrapper_get_jwt();
extern int kuzzle_wrapper_check_token(token_validity*, char*);
extern int kuzzle_wrapper_create_index(ack_response*, char*, query_options*);
extern int kuzzle_wrapper_login(login_result*, char*, json_object*, int*);
extern int kuzzle_wrapper_create_my_credentials(json_result*, char*, json_object*, query_options*);
extern void kuzzle_wrapper_disconnect();
extern void kuzzle_wrapper_flush_queue();
extern int kuzzle_wrapper_get_all_statistics(json_result*, query_options*);
extern int kuzzle_wrapper_get_auto_refresh(bool_result*, char*, query_options*);
extern int kuzzle_wrapper_get_my_credentials(json_result*, char*, query_options*);

#endif
