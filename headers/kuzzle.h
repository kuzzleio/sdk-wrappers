#ifndef _KUZZLE_H_
#define _KUZZLE_H_

#include <json/json.h>
#include <time.h>
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

//used for any boolean result
typedef struct {
    unsigned result;
    char error[2048];
} bool_result;

typedef struct {
    double result;
    char error[2048];
} now_result;

//used for get_statistics
typedef struct {
    json_object* completed_requests;
    json_object* connections;
    json_object* failed_requests;
    json_object* ongoing_requests;
    double timestamp;
    char error[2048];
} statistics;

//used for string array result
typedef struct {
    char **result;
    char error[2048];
} string_array_result;

typedef struct {
    int total;
    int successful;
    int failed;
    char error[2048];
} shards;

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
extern void kuzzle_wrapper_get_all_statistics(json_result*, query_options*);
extern int kuzzle_wrapper_get_auto_refresh(bool_result*, char*, query_options*);
extern int kuzzle_wrapper_get_my_credentials(json_result*, char*, query_options*);
extern void kuzzle_wrapper_get_my_rights(json_result*, query_options*);
extern void kuzzle_wrapper_get_server_info(json_result*, query_options*);
extern void kuzzle_wrapper_get_statistics(statistics*, time_t, query_options*);
extern int kuzzle_wrapper_list_collections(json_result*, char*, query_options*);
extern void kuzzle_wrapper_list_indexes(string_array_result*, query_options*);
extern char* kuzzle_wrapper_logout();
extern void kuzzle_wrapper_now(now_result*, query_options*);
extern void kuzzle_wrapper_refresh_index(shards*, char*, query_options*);
extern int kuzzle_wrapper_set_auto_refresh(bool_result*, char*, unsigned, query_options*);
extern void kuzzle_wrapper_set_default_index();

#endif
