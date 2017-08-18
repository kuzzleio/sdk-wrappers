%extend kuzzle {
    char* connect() {
        kuzzle_wrapper_connect();
    }

    offline_queue* getOfflineQueue() {
        return kuzzle_wrapper_get_offline_queue();
    }

    char* getJwt() {
        return kuzzle_wrapper_get_jwt();
    }

    %exception checkToken {
      $action
      if (result == $null) {
        jclass clazz = (*jenv)->FindClass(jenv, "java/lang/IllegalArgumentException");
        (*jenv)->ThrowNew(jenv, clazz, "Kuzzle.CheckToken: token required");
        return $null;
      }
    }
    static token_validity* checkToken(char* token) {
        static token_validity res;
        int err = kuzzle_wrapper_check_token(&res, token);

        if (err == 0) {
            return &res;
        }
        return (void*)0;
    }

    %exception createIndex {
      $action
      if (result == $null) {
        jclass clazz = (*jenv)->FindClass(jenv, "java/lang/IllegalArgumentException");
        (*jenv)->ThrowNew(jenv, clazz, "Kuzzle.createIndex: index required");
        return $null;
      }
    }
    static ack_response* createIndex(char* index, query_options *options) {
        static ack_response res;
        int err = kuzzle_wrapper_create_index(&res, index, options);

        if (err == 0) {
            return &res;
        }
        return (void*)0;
    }
}

%javaexception("java.lang.Exception") Kuzzle {
  $action
  if (!result) {
    jclass clazz = (*jenv)->FindClass(jenv, "java/lang/Exception");
    (*jenv)->ThrowNew(jenv, clazz, "Cannot connect");
    return $null;
  }
}
