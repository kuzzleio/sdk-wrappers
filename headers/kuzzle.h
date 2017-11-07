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

/* === Security === */

typedef json_object controllers;

typedef struct {
    char *index;
    char **collections;
    int collections_length;
} policy_restriction;

typedef struct {
    char *role_id;
    policy_restriction **restricted_to;
    int restricted_to_length;
} policy;

typedef struct {
    char *id;
    policy **policies;
    int policies_length;
    kuzzle *kuzzle;
} profile;

typedef struct {
    char *id;
    json_object *controllers;
    kuzzle *kuzzle;
} role;

//kuzzle user
typedef struct {
    char *id;
    json_object *content;
    char **profile_ids;
    uint profile_ids_length;
    kuzzle *kuzzle;
} user;

// user content passed to user constructor
typedef struct {
    json_object *content;
    char **profile_ids;
    uint profile_ids_length;
} user_data;

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
    profile *profile;
    int status;
    char *error;
    char *stack;
} profile_result;

typedef struct {
    profile **profiles;
    uint profiles_length;
    int status;
    char *error;
    char *stack;
} profiles_result;

typedef struct {
    role *role;
    int status;
    char *error;
    char *stack;
} role_result;

typedef struct {
    char *controller;
    char *action;
    char *index;
    char *collection;
    char *value;
} user_right;

typedef struct {
    user_right **user_rights;
    uint user_rights_length;
    int status;
    char *error;
    char *stack;
} user_rights_result;

typedef struct {
    user *user;
    int status;
    char *error;
    char *stack;
} user_result;

enum {
    ALLOWED=0,
    CONDITIONNAL=1,
    DENIED=2
} is_action_allowed;

//statistics
typedef struct {
    json_object* completed_requests;
    json_object* connections;
    json_object* failed_requests;
    json_object* ongoing_requests;
    unsigned long long timestamp;
} statistics;

typedef struct {
    statistics* result;
    int status;
    char *error;
    char *stack;
} statistics_result;

typedef struct {
    statistics* res;
    int res_size;
    int status;
    char *error;
    char *stack;
} all_statistics_result;

//check_token
typedef struct {
    bool valid;
    char *state;
    unsigned long long expires_at;
    int status;
    char *error;
    char *stack;
} token_validity;

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

//any array of json_object result
typedef struct {
    json_object **result;
    unsigned length;
    int status;
    char *error;
    char *stack;
} json_array_result;

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

//any double result
typedef struct {
    double result;
    int status;
    char *error;
    char *stack;
} double_result;

// any string result
typedef struct {
    char *result;
    int status;
    char *error;
    char *stack;
} string_result;

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
    uint length;
    uint total;
    char *scrollId;
} document_search;

typedef struct {
    profile **hits;
    int length;
    int total;
    char *scrollId;
} profile_search;

typedef struct {
    role** hits;
    int length;
    int total;
} role_search;

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

typedef struct {
    bool strict;
    json_object *fields;
    json_object *validators;
} specification;

typedef struct {
    specification *validation;
    char *index;
    char *collection;
} specification_entry;

typedef struct {
    specification *result;
    int status;
    char *error;
    char *stack;
} specification_result;

typedef struct {
    document_search *result;
    int status;
    char *error;
    char *stack;
} search_result;

typedef struct {
    profile_search *result;
    int status;
    char *error;
    char *stack;
} search_profiles_result;

typedef struct {
    role_search *result;
    int status;
    char *error;
    char *stack;
} search_roles_result;

typedef struct {
    specification_entry** hits;
    uint length;
    uint total;
    char *scrollId;
} specification_search;

typedef struct {
    specification_search *result;
    int status;
    char *error;
    char *stack;
} specification_search_result;

typedef struct {
    json_object *mapping;
    collection *collection;
} mapping;

typedef struct {
    mapping *result;
    int status;
    char *error;
    char *stack;
} mapping_result;

// Kuzzle main object functions
extern void kuzzle_wrapper_new_kuzzle(kuzzle*, char*, char*, options*);
extern char* kuzzle_wrapper_connect(kuzzle*);
extern offline_queue* kuzzle_wrapper_get_offline_queue(kuzzle*);
extern char* kuzzle_wrapper_get_jwt(kuzzle*);
extern token_validity* kuzzle_wrapper_check_token(kuzzle*, char*);
extern bool_result* kuzzle_wrapper_create_index(kuzzle*, char*, query_options*);
extern string_result* kuzzle_wrapper_login(kuzzle*, char*, json_object*, int*);
extern json_result* kuzzle_wrapper_create_my_credentials(kuzzle*, char*, json_object*, query_options*);
extern void kuzzle_wrapper_disconnect(kuzzle*);
extern void kuzzle_wrapper_flush_queue(kuzzle*);
extern all_statistics_result* kuzzle_wrapper_get_all_statistics(kuzzle*, query_options*);
extern bool_result* kuzzle_wrapper_get_auto_refresh(kuzzle*, char*, query_options*);
extern json_result* kuzzle_wrapper_get_my_credentials(kuzzle*, char*, query_options*);
extern json_result* kuzzle_wrapper_get_my_rights(kuzzle*, query_options*);
extern json_result* kuzzle_wrapper_get_server_info(kuzzle*, query_options*);
extern statistics_result* kuzzle_wrapper_get_statistics(kuzzle*, time_t, query_options*);
extern json_result* kuzzle_wrapper_list_collections(kuzzle*, char*, query_options*);
extern string_array_result* kuzzle_wrapper_list_indexes(kuzzle*, query_options*);
extern char* kuzzle_wrapper_logout(kuzzle*);
extern int_result* kuzzle_wrapper_now(kuzzle*, query_options*);
extern shards_result* kuzzle_wrapper_refresh_index(kuzzle*, char*, query_options*);
extern bool_result* kuzzle_wrapper_set_auto_refresh(kuzzle*, char*, bool, query_options*);
extern int kuzzle_wrapper_set_default_index(kuzzle*, char*);
extern void kuzzle_wrapper_unset_jwt(kuzzle*);
extern json_result* kuzzle_wrapper_update_self(kuzzle*, user_data*, query_options*);
extern bool_result* kuzzle_wrapper_validate_my_credentials(kuzzle*, char*, json_object*, query_options*);
extern user_result* kuzzle_wrapper_who_am_i(kuzzle*);
extern kuzzle_response* kuzzle_wrapper_query(kuzzle*, kuzzle_request*, query_options*);
extern void kuzzle_wrapper_set_headers(kuzzle*, json_object*, unsigned);
extern json_object* kuzzle_wrapper_get_headers(kuzzle*);
extern void kuzzle_wrapper_add_listener(kuzzle*, int, void*);
extern void kuzzle_wrapper_remove_listener(kuzzle*, int);
extern void kuzzle_wrapper_replay_queue(kuzzle*);
extern void kuzzle_wrapper_set_jwt(kuzzle*, char*);
extern void kuzzle_wrapper_start_queuing(kuzzle*);
extern void kuzzle_wrapper_stop_queuing(kuzzle*);
extern bool_result* kuzzle_wrapper_delete_my_credentials(kuzzle*, char*, query_options*);
extern json_result* kuzzle_wrapper_update_my_credentials(kuzzle*, char*, json_object*, query_options*);

//Options
extern options* kuzzle_wrapper_new_options();

//gc management
extern void unregisterKuzzle(kuzzle*);

//Collection
extern collection* kuzzle_wrapper_new_collection(kuzzle*, char*, char*);
extern bool_result* kuzzle_wrapper_collection_create(collection*, query_options*);
extern bool_result* kuzzle_wrapper_collection_publish_message(collection*, json_object*, query_options*);
extern void kuzzle_wrapper_collection_set_headers(collection*, json_object*, uint);
extern bool_result* kuzzle_wrapper_collection_truncate(collection*, query_options*);

//Collection Document
extern int_result* kuzzle_wrapper_collection_count(collection*, search_filters*, query_options*);
extern document_result* kuzzle_wrapper_collection_create_document(collection*, char*, document*, query_options*);
extern string_result* kuzzle_wrapper_collection_delete_document(collection*, char*, query_options*);
extern bool_result* kuzzle_wrapper_collection_document_exists(collection*, char*, query_options*);
extern document_result* kuzzle_wrapper_collection_fetch_document(collection*, char*, query_options*);
extern document_result* kuzzle_wrapper_collection_replace_document(collection*, char*, document*, query_options*);
extern document_result* kuzzle_wrapper_collection_update_document(collection*, char*, document*, query_options*);
extern search_result* kuzzle_wrapper_collection_scroll(collection*, char*, query_options*);
extern search_result* kuzzle_wrapper_collection_search(collection*, search_filters*, query_options*);
extern search_result* kuzzle_wrapper_collection_m_create_document(collection*, document**, uint, query_options*);
extern search_result* kuzzle_wrapper_collection_m_create_or_replace_document(collection*, document**, uint, query_options*);
extern string_array_result* kuzzle_wrapper_collection_m_delete_document(collection*, char**, uint, query_options*);
extern search_result* kuzzle_wrapper_collection_m_get_document(collection*, char**, uint, query_options*);
extern search_result* kuzzle_wrapper_collection_m_replace_document(collection*, document**, uint, query_options*);
extern search_result* kuzzle_wrapper_collection_m_update_document(collection*, document**, uint, query_options*);

//Collection Mapping
extern mapping_result* kuzzle_wrapper_collection_get_mapping(collection*, query_options*);
extern bool_result* kuzzle_wrapper_mapping_apply(mapping*, query_options*);
extern bool_result* kuzzle_wrapper_mapping_refresh(mapping*, query_options*);
extern void kuzzle_wrapper_mapping_set(mapping*, json_object*);
extern void kuzzle_wrapper_mapping_set_headers(mapping*, json_object*, uint);

//Collection Specification
extern bool_result* kuzzle_wrapper_collection_delete_specifications(collection*, query_options*);
extern specification_result* kuzzle_wrapper_collection_get_specifications(collection*, query_options*);
extern specification_search_result* kuzzle_wrapper_collection_scroll_specifications(collection*, char*, query_options*);
extern specification_search_result* kuzzle_wrapper_collection_search_specifications(collection*, search_filters*, query_options*);
extern specification_result* kuzzle_wrapper_collection_update_specifications(collection*, specification*, query_options*);
extern bool_result* kuzzle_wrapper_collection_validate_specifications(collection*, specification*, query_options*);

//Security
//  profile
extern profile* kuzzle_wrapper_security_new_profile(kuzzle* kuzzle, char* profile_id, policy** policies);
extern void kuzzle_wrapper_security_destroy_profile(profile* profile);
extern profile_result* kuzzle_wrapper_security_fetch_profile(kuzzle* kuzzle, char* profile_id, query_options* options);
extern search_profiles_result* kuzzle_wrapper_security_scroll_profiles(kuzzle* kuzzle, char* scrolle_id, query_options* options);
extern search_profiles_result* kuzzle_wrapper_security_search_profiles(kuzzle* kuzzle, search_filters* filters, query_options* options);
extern profile* kuzzle_wrapper_security_profile_add_policy(profile* profile, policy* policy);
extern string_result* kuzzle_wrapper_security_profile_delete(profile* profile, query_options* options);
extern profile_result* kuzzle_wrapper_security_profile_save(profile* profile, query_options* options);

//  role
extern role* kuzzle_wrapper_security_new_role(kuzzle* kuzzle, char* role_id, controllers* controllers);
extern void kuzzle_wrapper_security_destroy_role(role* role);
extern role_result* kuzzle_wrapper_security_fetch_role(kuzzle* kuzzle, char* role_id, query_options* options);
extern search_roles_result* kuzzle_wrapper_security_search_roles(kuzzle* kuzzle, search_filters* filters, query_options* options);
extern string_result* kuzzle_wrapper_security_role_delete(role* role, query_options* options);
extern role_result* kuzzle_wrapper_security_role_save(role* role, query_options* options);

//  user
extern user* kuzzle_wrapper_security_new_user(kuzzle* kuzzle, char* user_id, user_data* user_data);
extern void kuzzle_wrapper_security_destroy_user(user* user);
extern user_result* kuzzle_wrapper_security_user_create(user* user, query_options* options);
extern json_result* kuzzle_wrapper_security_user_create_credentials(user* user, char* strategy, json_object* credentials, query_options* options);
extern user_result* kuzzle_wrapper_security_user_create_with_credentials(user* user, json_object* credentials, query_options* options);
extern string_result* kuzzle_wrapper_security_user_delete(user* user, query_options* options);
extern bool_result* kuzzle_wrapper_security_user_delete_credentials(user* user, char* strategy, query_options* options);
extern json_result* kuzzle_wrapper_security_user_get_credentials_info(user* user, char* strategy, query_options* options);
extern profiles_result* kuzzle_wrapper_security_user_get_profiles(user* user, query_options* options);
extern user_rights_result* kuzzle_wrapper_security_user_get_rights(user* user, query_options* options);
extern bool_result* kuzzle_wrapper_security_user_has_credentials(user* user, char* strategy, query_options* options);
extern user_result* kuzzle_wrapper_security_user_replace(user* user, query_options* options);
extern json_result* kuzzle_wrapper_security_update_credentials(user* user, char* strategy, json_object* credentials, query_options* options);
extern unsigned int kuzzle_wrapper_security_is_action_allowed(user_right** rights, unsigned int rights_length, char* controller, char* action, char* index, char* collection);

#endif
