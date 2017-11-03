#ifndef __SDK_WRAPPERS_INTERNAL
#define __SDK_WRAPPERS_INTERNAL

typedef char *char_ptr;
typedef document *document_ptr;
typedef policy *policy_ptr;
typedef policy_restriction *policy_restriction_ptr;
typedef profile *profile_ptr;
typedef role *role_ptr;
typedef user *user_ptr;
typedef user_right *user_right_ptr;
typedef specification_entry *specification_entry_ptr;
typedef json_object *json_object_ptr;

static void set_errno(int err) {
  errno = err;
}

#endif
