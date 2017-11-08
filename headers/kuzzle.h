#ifndef _KUZZLE_
#define _KUZZLE_

#include <json-c/json.h>
#include <time.h>
#include <errno.h>
#include <stdbool.h>
#include "structs.h"

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
