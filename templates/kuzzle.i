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

    token_validity checkToken(char* token) {
        return kuzzle_wrapper_check_token(token);
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

%javaexception("java.lang.Exception") kuzzle_wrapper_check_token {
  $action
  if (!result) {
    jclass clazz = (*jenv)->FindClass(jenv, "java/lang/Exception");
    (*jenv)->ThrowNew(jenv, clazz, result);
    return $null;
  }
}
