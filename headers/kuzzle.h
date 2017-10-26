#ifndef _KUZZLE_H_
#define _KUZZLE_H_

#include <json-c/json.h>
#include <time.h>
#include <errno.h>
#include <stdbool.h>
#include "../templates/swig.h"

typedef struct {
    void *instance;
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
    unsigned long long expires_in;
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
    bool queuable;
    long from;
    long size;
    char *scroll;
    char *scroll_id;
    char *refresh;
    char *if_exist;
    int retry_on_conflict;
    json_object *volatiles;
} query_options;

//options passed to room constructor
typedef struct {
    char *scope;
    char *state;
    char *user;
    int subscribe_to_self;
    json_object *volatiles;
} room_options;

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
} options;

//meta of a document
typedef struct {
    char *author;
    unsigned long long created_at;
    unsigned long long updated_at;
    char *updater;
    bool active;
    unsigned long long deleted_at;
} meta;

//kuzzle user
typedef struct {
    char *id;
    json_object* source;
    meta* meta;
    char **strategies;
    unsigned long strategies_length;
    int status;
    char *error;
    char *stack;
} user;

/* === Dedicated response structures === */

typedef struct {
  int failed;
  int successful;
  int total;
} shards;

typedef struct {
    char *index;
    char *collection;
    kuzzle *kuzzle;
} collection;

typedef struct {
    char *id;
    char *index;
    meta *meta;
    shards *shards;
    json_object *content;
    int version;
    char *result;
    bool created;
    char *collection;
    collection *_collection;
} document;

typedef struct {
    document *result;
    int status;
    char *error;
    char *stack;
} document_result;

typedef struct {
    char *result;
    int status;
    char *error;
    char *stack;
} string_result;

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
    bool valid;
    char *state;
    unsigned long long expires_at;
    int status;
    char *error;
    char *stack;
} token_validity;

//login
typedef struct {
    char *jwt;
    int status;
    char *error;
    char *stack;
} login_result;

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
    _json_object *result;
    int status;
    char *error;
    char *stack;
} json_result;

//any boolean result
typedef struct {
    bool result;
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

typedef struct {
    char* type;
    json_object* fields;
} field_mapping;

typedef struct {
    json_object* query;
    json_object* sort;
    json_object* aggregations;
    json_object* search_after;
} search_filters;

typedef struct {
    document** hits;
    int length;
    int total;
    char *scrollId;
} search_result;

//any delete* function
typedef struct {
    bool acknowledged;
    bool shards_acknowledged;
    int status;
    char *error;
    char *stack;
} ack_result;

typedef struct {
    shards *result;
    int status;
    char *error;
    char *stack;
} shards_result;

// TODO
typedef struct {

} kuzzle_specification;

typedef struct {
    kuzzle_specification *result;
    int status;
    char *error;
    char *stack;
} kuzzle_specification_result;

typedef struct {
    char *request_id;
    search_result *result;
    char *room_id;
    char *channel;
    int status;
    char *error;
    char *stack;
} kuzzle_search_result;

typedef struct {
    kuzzle_specification** hits;
    int total;
    char *scrollId;
} specification_search_result;

typedef struct {
    specification_search_result *result;
    int status;
    char *error;
    char *stack;
} kuzzle_specification_search_result;

// TODO
typedef struct {

} collection_mapping;

typedef struct {
    collection_mapping *instance;
    int status;
    char *error;
    char *stack;
} collection_mapping_result;

// Kuzzle main object functions
extern void kuzzle_wrapper_new_kuzzle(kuzzle*, char*, char*, options*);
extern char* kuzzle_wrapper_connect(kuzzle*);
extern offline_queue* kuzzle_wrapper_get_offline_queue(kuzzle*);
extern char* kuzzle_wrapper_get_jwt(kuzzle*);
extern token_validity* kuzzle_wrapper_check_token(kuzzle*, char*);
extern ack_result* kuzzle_wrapper_create_index(kuzzle*, char*, query_options*);
extern login_result* kuzzle_wrapper_login(kuzzle*, char*, json_object*, int*);
extern json_result* kuzzle_wrapper_create_my_credentials(kuzzle*, char*, json_object*, query_options*);
extern void kuzzle_wrapper_disconnect(kuzzle*);
extern void kuzzle_wrapper_flush_queue(kuzzle*);
extern json_result* kuzzle_wrapper_get_all_statistics(kuzzle*, query_options*);
extern bool_result* kuzzle_wrapper_get_auto_refresh(kuzzle*, char*, query_options*);
extern json_result* kuzzle_wrapper_get_my_credentials(kuzzle*, char*, query_options*);
extern json_result* kuzzle_wrapper_get_my_rights(kuzzle*, query_options*);
extern json_result* kuzzle_wrapper_get_server_info(kuzzle*, query_options*);
extern statistics* kuzzle_wrapper_get_statistics(kuzzle*, time_t, query_options*);
extern json_result* kuzzle_wrapper_list_collections(kuzzle*, char*, query_options*);
extern string_array_result* kuzzle_wrapper_list_indexes(kuzzle*, query_options*);
extern char* kuzzle_wrapper_logout(kuzzle*);
extern int_result* kuzzle_wrapper_now(kuzzle*, query_options*);
extern shards_result* kuzzle_wrapper_refresh_index(kuzzle*, char*, query_options*);
extern bool_result* kuzzle_wrapper_set_auto_refresh(kuzzle*, char*, unsigned, query_options*);
extern int kuzzle_wrapper_set_default_index(kuzzle*, char*);
extern void kuzzle_wrapper_unset_jwt(kuzzle*);
extern json_result* kuzzle_wrapper_update_self(kuzzle*, json_object*, query_options*);
extern bool_result* kuzzle_wrapper_validate_my_credentials(kuzzle*, char*, json_object*, query_options*);
extern user* kuzzle_wrapper_who_am_i(kuzzle*);
extern kuzzle_response* kuzzle_wrapper_query(kuzzle*, kuzzle_request*, query_options*);
extern void kuzzle_wrapper_set_headers(kuzzle*, json_object*, unsigned);
extern json_object* kuzzle_wrapper_get_headers(kuzzle*);
extern void kuzzle_wrapper_add_listener(kuzzle*, int, void*);
extern void kuzzle_wrapper_remove_listener(kuzzle*, int);
extern void kuzzle_wrapper_replay_queue(kuzzle*);
extern void kuzzle_wrapper_set_jwt(kuzzle*, char*);
extern void kuzzle_wrapper_start_queuing(kuzzle*);
extern void kuzzle_wrapper_stop_queuing(kuzzle*);
extern ack_result* kuzzle_wrapper_delete_my_credentials(kuzzle*, char*, query_options*);
extern json_result* kuzzle_wrapper_update_my_credentials(kuzzle*, char*, json_object*, query_options*);

//Options
extern options* kuzzle_wrapper_new_options();

//Json
extern void kuzzle_wrapper_json_put(json_object*, char*, void*, int);
extern char* kuzzle_wrapper_json_get_string(json_object*, char*);
extern int kuzzle_wrapper_json_get_int(json_object*, char*);
extern double kuzzle_wrapper_json_get_double(json_object*, char*);
extern json_bool kuzzle_wrapper_json_get_bool(json_object*, char*);
extern _json_object kuzzle_wrapper_json_get_json_object(json_object*, char*);
extern void kuzzle_wrapper_json_new(json_object**);

//gc management
extern void unregisterKuzzle(kuzzle*);

#endif
