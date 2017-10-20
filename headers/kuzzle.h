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
    long from;
    long size;
    char *scroll;
    char *scroll_id;
    char *strategy;
    unsigned long expires_in;
    json_object *volatiles;
    char *scope;
    char *state;
    char *user;
    long start;
    long stop;
    long end;
    unsigned char bit;
    char *member;
    char *member1;
    char *member2;
    char **members;
    unsigned members_length;
    double lon;
    double lat;
    double distance;
    char *unit;
    json_object *options;
    char **keys;
    unsigned keys_length;
    long cursor;
    long offset;
    char *field;
    char **fields;
    unsigned fields_length;
    char *subcommand;
    char *pattern;
    long idx;
    char *min;
    char *max;
    char *limit;
    unsigned long count;
    char *match;
} kuzzle_request;

//query object used by query()
typedef struct {
    json_object *query;
    unsigned long long timestamp;
    char   *request_id;
} query_object;

typedef struct {
    query_object **queries;
    unsigned long length;
} offline_queue;


//options passed to query()
typedef struct {
    unsigned queuable;
    long from;
    long size;
    char *scroll;
    char *scroll_id;
    char *refresh;
    char *if_exist;
    unsigned char retry_on_conflict;
    json_object *volatiles;
} query_options;

enum Mode {AUTO, MANUAL};
//options passed to the Kuzzle() fct
typedef struct {
    unsigned queue_ttl;
    unsigned long queue_max_size;
    unsigned char offline_mode;
    unsigned char auto_queue;
    unsigned char auto_reconnect;
    unsigned char auto_replay;
    unsigned char auto_resubscribe;
    unsigned long reconnection_delay;
    unsigned long replay_interval;
    enum Mode connect;
    char *refresh;
    char *default_index;
    json_object    *headers;
} Options;

//meta of a document
typedef struct {
    char *author;
    unsigned long long created_at;
    unsigned long long updated_at;
    char *updater;
    unsigned char active;
    unsigned long long deleted_at;
} kuzzle_meta;

//kuzzle user
typedef struct {
    char *id;
    json_object* source;
    kuzzle_meta* meta;
    char **strategies;
    unsigned long strategies_length;
    int status;
    char *error;
    char *stack;
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
    char *stack;
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
    char *stack;
} login_result;

//refresh_index
typedef struct {
    int total;
    int successful;
    int failed;
    int status;
    char *error;
    char *stack;
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
    char *stack;
} int_result;

//any array of strings result
typedef struct {
    char **result;
    unsigned long length;
    int status;
    char *error;
    char *stack;
} string_array_result;


// Kuzzle main object functions
extern void kuzzle_wrapper_new_kuzzle(Kuzzle*, char*, char*, Options*);
extern char* kuzzle_wrapper_connect(Kuzzle*);
extern offline_queue* kuzzle_wrapper_get_offline_queue(Kuzzle*);
extern char* kuzzle_wrapper_get_jwt(Kuzzle*);
extern token_validity* kuzzle_wrapper_check_token(Kuzzle*, char*);
extern ack_response* kuzzle_wrapper_create_index(Kuzzle*, char*, query_options*);
extern login_result* kuzzle_wrapper_login(Kuzzle*, char*, json_object*, int*);
extern json_result* kuzzle_wrapper_create_my_credentials(Kuzzle*, char*, json_object*, query_options*);
extern void kuzzle_wrapper_disconnect(Kuzzle*);
extern void kuzzle_wrapper_flush_queue(Kuzzle*);
extern json_result* kuzzle_wrapper_get_all_statistics(Kuzzle*, query_options*);
extern bool_result* kuzzle_wrapper_get_auto_refresh(Kuzzle*, char*, query_options*);
extern json_result* kuzzle_wrapper_get_my_credentials(Kuzzle*, char*, query_options*);
extern json_result* kuzzle_wrapper_get_my_rights(Kuzzle*, query_options*);
extern json_result* kuzzle_wrapper_get_server_info(Kuzzle*, query_options*);
extern statistics* kuzzle_wrapper_get_statistics(Kuzzle*, time_t, query_options*);
extern json_result* kuzzle_wrapper_list_collections(Kuzzle*, char*, query_options*);
extern string_array_result* kuzzle_wrapper_list_indexes(Kuzzle*, query_options*);
extern char* kuzzle_wrapper_logout(Kuzzle*);
extern int_result* kuzzle_wrapper_now(Kuzzle*, query_options*);
extern shards* kuzzle_wrapper_refresh_index(Kuzzle*, char*, query_options*);
extern bool_result* kuzzle_wrapper_set_auto_refresh(Kuzzle*, char*, unsigned, query_options*);
extern int kuzzle_wrapper_set_default_index(Kuzzle*, char*);
extern void kuzzle_wrapper_unset_jwt(Kuzzle*);
extern json_result* kuzzle_wrapper_update_self(Kuzzle*, json_object*, query_options*);
extern bool_result* kuzzle_wrapper_validate_my_credentials(Kuzzle*, char*, json_object*, query_options*);
extern user* kuzzle_wrapper_who_am_i(Kuzzle*);
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
extern json_result* kuzzle_wrapper_update_my_credentials(Kuzzle*, char*, json_object*, query_options*);
//Options
extern Options* kuzzle_wrapper_new_options(void);

#endif
