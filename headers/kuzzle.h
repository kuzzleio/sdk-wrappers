#ifndef _KUZZLE_H_
#define _KUZZLE_H_

#include <json/json.h>
#include <time.h>
#include <errno.h>

typedef struct {
    void* instance;
} kuzzle;

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

//enum used for options.connect
enum {AUTO, MANUAL} connect_mode;
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
    int connect;
    char refresh[64];
    char default_index[128];
    json_object    *headers;
} options;

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

extern void Kuzzle(kuzzle*, char*, char*, options*);
extern char* kuzzle_wrapper_connect(kuzzle*);
extern offline_queue* kuzzle_wrapper_get_offline_queue();
extern char* kuzzle_wrapper_get_jwt(kuzzle*);
extern int kuzzle_wrapper_check_token(kuzzle*, token_validity*, char*);
extern int kuzzle_wrapper_create_index(kuzzle*, ack_response*, char*, query_options*);
extern int kuzzle_wrapper_login(kuzzle*, login_result*, char*, json_object*, int*);
extern int kuzzle_wrapper_create_my_credentials(kuzzle*, json_result*, char*, json_object*, query_options*);
extern void kuzzle_wrapper_disconnect(kuzzle*);
extern void kuzzle_wrapper_flush_queue(kuzzle*);
extern void kuzzle_wrapper_get_all_statistics(kuzzle*, json_result*, query_options*);
extern int kuzzle_wrapper_get_auto_refresh(kuzzle*, bool_result*, char*, query_options*);
extern int kuzzle_wrapper_get_my_credentials(kuzzle*, json_result*, char*, query_options*);
extern void kuzzle_wrapper_get_my_rights(kuzzle*, json_result*, query_options*);
extern void kuzzle_wrapper_get_server_info(kuzzle*, json_result*, query_options*);
extern void kuzzle_wrapper_get_statistics(kuzzle*, statistics*, time_t, query_options*);
extern int kuzzle_wrapper_list_collections(kuzzle*, json_result*, char*, query_options*);
extern void kuzzle_wrapper_list_indexes(kuzzle*, string_array_result*, query_options*);
extern char* kuzzle_wrapper_logout(kuzzle*);
extern void kuzzle_wrapper_now(kuzzle*, now_result*, query_options*);
extern void kuzzle_wrapper_refresh_index(kuzzle*, shards*, char*, query_options*);
extern int kuzzle_wrapper_set_auto_refresh(kuzzle*, bool_result*, char*, unsigned, query_options*);
extern int kuzzle_wrapper_set_default_index(kuzzle*, char*);
extern void kuzzle_wrapper_unset_jwt(kuzzle*);
extern void kuzzle_wrapper_update_self(kuzzle*, json_result*, json_object*, query_options*);
extern void kuzzle_wrapper_validate_my_credentials(kuzzle*, bool_result*, char*, json_object*, query_options*);
extern void kuzzle_wrapper_who_am_i(kuzzle*, user*);
extern void kuzzle_wrapper_query(kuzzle*, kuzzle_response*, kuzzle_request*, query_options*);
extern void kuzzle_wrapper_set_headers(kuzzle*, json_object*, unsigned);
extern json_object* kuzzle_wrapper_get_headers(kuzzle*);
extern void kuzzle_wrapper_set_jwt(kuzzle*, char*);
extern void kuzzle_wrapper_replay_queue(kuzzle*);
extern void kuzzle_wrapper_add_listener(kuzzle*, int, void*);
extern void kuzzle_wrapper_remove_listener(kuzzle*, int);

#endif
