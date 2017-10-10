#ifndef _KUZZLE_H_
#define _KUZZLE_H_

#include <json-c/json.h>
#include <time.h>
#include <errno.h>

typedef struct {
    void* instance;
} Kuzzle;

enum {
    CONNECTED,
    DISCARDED,
    DISCONNECTED,
    LOGIN_ATTEMPT,
    NETWORK_ERROR,
    OFFLINE_QUEUE_POP,
    OFFLINE_QUEUE_PUSH,
    QUERY_ERROR,
    RECONNECTED,
    JWT_EXPIRED,
    ERROR
} event;

//define a request
typedef struct {
    char request_id[36];
    char controller[128];
    char action[128];
    char index[128];
    char collection[128];
    json_object *body;
    char id[128];
    int from;
    int size;
    char scroll[32];
    char scroll_id[128];
    char strategy[128];
    int expires_in;
    json_object* volatiles;
    char scope[512];
    char state[512];
    char user[512];
    int start;
    int stop;
    int end;
    int bit;
    char member[512];
    char member1[512];
    char member2[512];
    char **members;
    float lon;
    float lat;
    float distance;
    char unit[128];
    json_object* options;
    char **keys;
    int cursor;
    int offset;
    char field[512];
    char **fields;
    char subcommand[1024];
    char pattern[1024];
    int idx;
    char min[512];
    char max[512];
    char limit[512];
    int count;
    char match[512];
} kuzzle_request;

//query object used by query()
typedef struct {
    json_object *query;
    unsigned long long   timestamp;
    char   request_id[36];
} query_object;

typedef struct {
    query_object** query;
} offline_queue;

//response of check_token()
typedef struct token_validity_struct {
    unsigned valid;
    char state[512];
    int expiresAt;
    char error[2048];
} token_validity;

//response for any delete* function
typedef struct ack_response_struct {
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

enum Mode {AUTO, MANUAL};
//options passed to the Kuzzle() fct
typedef struct {
    double queue_ttl;
    int queue_max_size;
    int offline_mode;
    unsigned auto_queue;
    unsigned auto_reconnect;
    unsigned auto_replay;
    unsigned auto_resubscribe;
    double reconnection_delay;
    double replay_interval;
    enum Mode connect;
    char refresh[64];
    char default_index[128];
    json_object    *headers;
} Options;

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

//used for now result
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

//used for refresh_index
typedef struct {
    int total;
    int successful;
    int failed;
    char error[2048];
} shards;

//meta of a document
typedef struct {
    char author[512];
    int created_at;
    int updated_at;
    char updater[512];
    unsigned active;
    int deleted_at;
} kuzzle_meta;

//kuzzle user
typedef struct {
    char id[512];
    json_object* source;
    kuzzle_meta* meta;
    char **strategies;
    char error[2048];
} user;

typedef struct {
    char request_id[36];
    json_object* result;
    char room_id[36];
    char channel[128];
    char error[2048];
} kuzzle_response;

// Kuzzle main object functions
extern void kuzzle_wrapper_new_kuzzle(Kuzzle*, char*, char*, Options*);
extern char* kuzzle_wrapper_connect(Kuzzle*);
extern void kuzzle_wrapper_get_offline_queue(Kuzzle*, offline_queue*);
extern char* kuzzle_wrapper_get_jwt(Kuzzle*);
extern int kuzzle_wrapper_check_token(Kuzzle*, token_validity*, char*);
extern int kuzzle_wrapper_create_index(Kuzzle*, ack_response*, char*, query_options*);
extern int kuzzle_wrapper_login(Kuzzle*, login_result*, char*, json_object*, int*);
extern int kuzzle_wrapper_create_my_credentials(Kuzzle*, json_result*, char*, json_object*, query_options*);
extern void kuzzle_wrapper_disconnect(Kuzzle*);
extern void kuzzle_wrapper_flush_queue(Kuzzle*);
extern void kuzzle_wrapper_get_all_statistics(Kuzzle*, json_result*, query_options*);
extern int kuzzle_wrapper_get_auto_refresh(Kuzzle*, bool_result*, char*, query_options*);
extern int kuzzle_wrapper_get_my_credentials(Kuzzle*, json_result*, char*, query_options*);
extern void kuzzle_wrapper_get_my_rights(Kuzzle*, json_result*, query_options*);
extern void kuzzle_wrapper_get_server_info(Kuzzle*, json_result*, query_options*);
extern void kuzzle_wrapper_get_statistics(Kuzzle*, statistics*, time_t, query_options*);
extern int kuzzle_wrapper_list_collections(Kuzzle*, json_result*, char*, query_options*);
extern void kuzzle_wrapper_list_indexes(Kuzzle*, string_array_result*, query_options*);
extern char* kuzzle_wrapper_logout(Kuzzle*);
extern void kuzzle_wrapper_now(Kuzzle*, now_result*, query_options*);
extern void kuzzle_wrapper_refresh_index(Kuzzle*, shards*, char*, query_options*);
extern int kuzzle_wrapper_set_auto_refresh(Kuzzle*, bool_result*, char*, unsigned, query_options*);
extern int kuzzle_wrapper_set_default_index(Kuzzle*, char*);
extern void kuzzle_wrapper_unset_jwt(Kuzzle*);
extern void kuzzle_wrapper_update_self(Kuzzle*, json_result*, json_object*, query_options*);
extern void kuzzle_wrapper_validate_my_credentials(Kuzzle*, bool_result*, char*, json_object*, query_options*);
extern void kuzzle_wrapper_who_am_i(Kuzzle*, user*);
extern void kuzzle_wrapper_query(Kuzzle*, kuzzle_response*, kuzzle_request*, query_options*);
extern void kuzzle_wrapper_set_headers(Kuzzle*, json_object*, unsigned);
extern json_object* kuzzle_wrapper_get_headers(Kuzzle*);
extern void kuzzle_wrapper_add_listener(Kuzzle*, int, void*);
extern void kuzzle_wrapper_remove_listener(Kuzzle*, int);
extern void kuzzle_wrapper_replay_queue(Kuzzle*);
extern void kuzzle_wrapper_set_jwt(Kuzzle*, char*);
extern void kuzzle_wrapper_start_queuing(Kuzzle*);
extern void kuzzle_wrapper_stop_queuing(Kuzzle*);

//Options
extern void kuzzle_wrapper_new_options(Options*);

#endif
