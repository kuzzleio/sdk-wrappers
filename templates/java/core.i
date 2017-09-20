%pragma(java) jniclasscode=%{
  static {
    try {
        System.loadLibrary("kcore");
    } catch (UnsatisfiedLinkError e) {
      System.err.println("Native code library failed to load. \n" + e);
      System.exit(1);
    }
  }
%}

%rename(TokenValidity) token_validity_struct;
%rename(AckResponse) ack_response_struct;
%rename(queueTTL) queue_ttl;

%extend Options {
    Options() {
        Options *o = malloc(sizeof(Options));
        kuzzle_wrapper_new_options(o);
        return o;
    }

    ~Options() {
        free($self);
    }
}

%typemap(javaimports) Kuzzle "
/* The type Kuzzle. */"

%extend Kuzzle {
    // ctors && dtor
    Kuzzle(char* host, Options *opts) {
        Kuzzle *k = malloc(sizeof(Kuzzle));
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

    // checkToken
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

    // connect
    %exception connect {
        $action
        if (result != $null) {
            jclass clazz = (*jenv)->FindClass(jenv, "java/lang/IllegalArgumentException");
            (*jenv)->ThrowNew(jenv, clazz, result);
        }
    }
    char* connect() {
        return kuzzle_wrapper_connect($self);
    }

    // createIndex
    %exception createIndex {
        $action
        if (result == $null) {
            jclass clazz = (*jenv)->FindClass(jenv, "java/lang/IllegalArgumentException");
            (*jenv)->ThrowNew(jenv, clazz, "Kuzzle.createIndex: index required");
            return $null;
        }
    }
    ack_response* createIndex(char* index, query_options* options) {
        static ack_response res;
        int err = kuzzle_wrapper_create_index($self, &res, index, options);

        if (err == 0) {
            return &res;
        }
        return (void*)0;
    }
    ack_response* createIndex(char* index) {
        static ack_response res;
        int err = kuzzle_wrapper_create_index($self, &res, index, (void*)0);

        if (err == 0) {
            return &res;
        }
        return (void*)0;
    }
}

%include "../../kcore.i"
