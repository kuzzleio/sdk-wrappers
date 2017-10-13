#ifndef _KUZZLE_H_
#define _KUZZLE_H_

#include <json-c/json.h>
#include <time.h>
#include <errno.h>

typedef struct {
    void *instance;
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
    char *request_id;
    char *controller;
    char *action;
    char *index;
    char *collection;
    json_object *body;
    char *id;
    int from;
    int size;
    char *scroll;
    char *scroll_id;
    char *strategy;
    int expires_in;
    json_object *volatiles;
    char *scope;
    char *state;
    char *user;
    int start;
    int stop;
    int end;
    int bit;
    char *member;
    char *member1;
    char *member2;
    char **members;
    float lon;
    float lat;
    float distance;
    char *unit;
    json_object *options;
    char **keys;
    int cursor;
    int offset;
    char *field;
    char **fields;
    char *subcommand;
    char *pattern;
    int idx;
    char *min;
    char *max;
    char *limit;
    int count;
    char *match;
} kuzzle_request;

//query object used by query()
typedef struct {
    json_object *query;
    unsigned long long timestamp;
    char   *request_id;
} query_object;

typedef struct {
    query_object **query;
} offline_queue;


//options passed to query()
typedef struct {
    unsigned queuable;
    int from;
    int size;
    char *scroll;
    char *scroll_id;
    char *refresh;
    char *if_exist;
    int retry_on_conflict;
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
    char *refresh;
    char *default_index;
    json_object    *headers;
} Options;

//meta of a document
typedef struct {
    char *author;
    int created_at;
    int updated_at;
    char *updater;
    unsigned active;
    int deleted_at;
} kuzzle_meta;

//kuzzle user
typedef struct {
    char *id;
    json_object* source;
    kuzzle_meta* meta;
    char **strategies;
    int status;
    char *error;
} user;

/* === Dedicated response structures === */

//statistics
typedef struct {
    json_object* completed_requests;
    json_object* connections;
    json_object* failed_requests;
    json_object* ongoing_requests;
    unsigned long long timestamp;
    int status;
    char *error;
} statistics;

//check_token
typedef struct {
    unsigned valid;
    char *state;
    long long expiresAt;
    int status;
    char *error;
    char *stack;
} token_validity;

//any delete* function
typedef struct {
    unsigned acknowledged;
    unsigned shardsAcknowledged;
    int status;
    char *error;
    char *stack;
} ack_response;

//login
typedef struct {
    char *jwt;
    int status;
    char *error;
} login_result;

//refresh_index
typedef struct {
    int total;
    int successful;
    int failed;
    int status;
    char *error;
} shards;

/* === Generic response structures === */

// raw Kuzzle response
typedef struct {
    char *request_id;
    json_object *result;
    char *room_id;
    char *channel;
    int status;
    char *error;
    char *stack;
} kuzzle_response;

//any json result
typedef struct {
    json_object *result;
    int status;
    char *error;
    char *stack;
} json_result;

//any boolean result
typedef struct {
    unsigned result;
    int status;
    char *error;
    char *stack;
} bool_result;

//any integer result
typedef struct {
    long long result;
    int status;
    char *error;
} now_result;

//used for string array result
typedef struct {
    char **result;
    int status;
    char *error;
} string_array_result;


// Kuzzle main object functions
extern void kuzzle_wrapper_new_kuzzle(Kuzzle*, char*, char*, Options*);
extern char* kuzzle_wrapper_connect(Kuzzle*);
extern void kuzzle_wrapper_get_offline_queue(Kuzzle*, offline_queue*);
extern char* kuzzle_wrapper_get_jwt(Kuzzle*);
extern token_validity* kuzzle_wrapper_check_token(Kuzzle*, char*);
extern ack_response* kuzzle_wrapper_create_index(Kuzzle*, char*, query_options*);
extern int kuzzle_wrapper_login(Kuzzle*, login_result*, char*, json_object*, int*);
extern json_result* kuzzle_wrapper_create_my_credentials(Kuzzle*, char*, json_object*, query_options*);
extern void kuzzle_wrapper_disconnect(Kuzzle*);
extern void kuzzle_wrapper_flush_queue(Kuzzle*);
extern json_result* kuzzle_wrapper_get_all_statistics(Kuzzle*, query_options*);
extern bool_result* kuzzle_wrapper_get_auto_refresh(Kuzzle*, char*, query_options*);
extern json_result* kuzzle_wrapper_get_my_credentials(Kuzzle*, char*, query_options*);
extern json_result kuzzle_wrapper_get_my_rights(Kuzzle*, query_options*);
extern json_result* kuzzle_wrapper_get_server_info(Kuzzle*, query_options*);
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
extern kuzzle_response* kuzzle_wrapper_query(Kuzzle*, kuzzle_request*, query_options*);
extern void kuzzle_wrapper_set_headers(Kuzzle*, json_object*, unsigned);
extern json_object* kuzzle_wrapper_get_headers(Kuzzle*);
extern void kuzzle_wrapper_add_listener(Kuzzle*, int, void*);
extern void kuzzle_wrapper_remove_listener(Kuzzle*, int);
extern void kuzzle_wrapper_replay_queue(Kuzzle*);
extern void kuzzle_wrapper_set_jwt(Kuzzle*, char*);
extern void kuzzle_wrapper_start_queuing(Kuzzle*);
extern void kuzzle_wrapper_stop_queuing(Kuzzle*);
extern ack_response* kuzzle_wrapper_delete_my_credentials(Kuzzle*, char*, query_options*);

//Options
extern void kuzzle_wrapper_new_options(Options*);

#endif
