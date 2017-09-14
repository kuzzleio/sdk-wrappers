%extend Kuzzle {
    Kuzzle(char* host, options *opts) {
        Kuzzle *k;
        k = malloc(sizeof(Kuzzle));
        kuzzle_wrapper_new_kuzzle(k, host, "websocket", opts);
        return k;
    }

    Kuzzle(char* host) {
        Kuzzle *k;
        k = malloc(sizeof(Kuzzle));
        kuzzle_wrapper_new_kuzzle(k, host, "websocket", (void*)0);
        return k;
    }

    ~Kuzzle() {
        free($self);
    }

    %exception checkToken {
      $action
      if (result == $null) {
        jclass clazz = (*jenv)->FindClass(jenv, "java/lang/IllegalArgumentException");
        (*jenv)->ThrowNew(jenv, clazz, "Kuzzle.CheckToken: token required");
        return $null;
      }
    }
    token_validity* checkToken(char* token) {
        static token_validity res;
        int err = kuzzle_wrapper_check_token($self, &res, token);

        if (err == 0) {
            return &res;
        }
        return (void*)0;
    }

    char* connect() {
        kuzzle_wrapper_connect($self);
    }
}