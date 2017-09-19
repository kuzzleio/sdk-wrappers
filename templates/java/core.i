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
}

%include "../../kcore.i"
