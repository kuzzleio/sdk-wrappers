#ifndef __SDK_WRAPPERS_INTERNAL
#define __SDK_WRAPPERS_INTERNAL

typedef char *char_ptr;
typedef document *document_ptr;
typedef specification_entry *specification_entry_ptr;

// used by memory_storage.geopos
typedef double geopos_arr[2];

static void set_errno(int err) {
  errno = err;
}

#endif
