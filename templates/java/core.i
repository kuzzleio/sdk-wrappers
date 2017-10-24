%rename(TokenValidity) token_validity;
%rename(AckResponse) ack_response;
%rename(queueTTL) queue_ttl;
%rename(Options) options;
%rename(Kuzzle) kuzzle;
%rename(JsonObject) json_object;

%include "../../kcore.i"

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


%extend options {
    options() {
        options *o = kuzzle_wrapper_new_options();
        return o;
    }

    ~options() {
        free($self);
    }
}

%extend json_object {
    json_object() {
        json_object* j = malloc(sizeof(json_object));
        kuzzle_wrapper_json_new(j);
        return j;
    }

    ~json_object() {
        free($self);
    }

    json_object* put(char* key, char* content) {
        kuzzle_wrapper_json_put($self, key, content, 0);
        return $self;
    }

    json_object* put(char* key, int content) {
        kuzzle_wrapper_json_put($self, key, &content, 1);
        return $self;
    }

    json_object* put(char* key, double content) {
        kuzzle_wrapper_json_put($self, key, &content, 2);
        return $self;
    }

    json_object* put(char* key, bool content) {
        kuzzle_wrapper_json_put($self, key, &content, 3);
        return $self;
    }

    json_object* put(char* key, JsonObject* content) {
        kuzzle_wrapper_json_put($self, key, content, 4);
        return $self;
    }

    char* getString(char* key) {
        return kuzzle_wrapper_json_get_string($self, key);
    }

    int getInt(char* key) {
        return kuzzle_wrapper_json_get_int($self, key);
    }

    double getDouble(char* key) {
        return kuzzle_wrapper_json_get_double($self, key);
    }

    bool getBoolean(char* key) {
        return kuzzle_wrapper_json_get_bool($self, key);
    }

    json_object getJsonObject(char* key) {
        return kuzzle_wrapper_json_get_json_object($self->jobj, key);
    }
}

%typemap(javaimports) kuzzle "
/* The type Kuzzle. */"

%extend kuzzle {
    // ctors && dtor
    kuzzle(char* host, options *opts) {
        kuzzle *k = malloc(sizeof(kuzzle));
        kuzzle_wrapper_new_kuzzle(k, host, "websocket", opts);
        return k;
    }
//    _kuzzle(char* host) {
//        kuzzle *k;
//        k = malloc(sizeof(kuzzle));
//        kuzzle_wrapper_new_kuzzle(k, host, "websocket", (void*)0);
//        return k;
//    }
//    ~_kuzzle() {
//        unregisterKuzzle($self);
//        free($self);
//    }

    // checkToken
//    %exception checkToken {
//        $action
//        if (result == $null) {
//            jclass clazz = (*jenv)->FindClass(jenv, "java/lang/IllegalArgumentException");
//            (*jenv)->ThrowNew(jenv, clazz, "Kuzzle.CheckToken: token required");
//            return $null;
//        }
//    }
//    token_validity* checkToken(char* token) {
//        static token_validity res;
//        int err = kuzzle_wrapper_check_token($self, &res, token);
//
//        if (err == 0) {
//            return &res;
//        }
//        return (void*)0;
//    }
//
//    // connect
//    %exception connect {
//        $action
//        if (result != $null) {
//            jclass clazz = (*jenv)->FindClass(jenv, "java/lang/IllegalArgumentException");
//            (*jenv)->ThrowNew(jenv, clazz, result);
//        }
//    }
//    char* connect() {
//        return kuzzle_wrapper_connect($self);
//    }
//
//    // createIndex
//    %exception createIndex {
//        $action
//        if (result == $null) {
//            jclass clazz = (*jenv)->FindClass(jenv, "java/lang/IllegalArgumentException");
//            (*jenv)->ThrowNew(jenv, clazz, "Kuzzle.createIndex: index required");
//            return $null;
//        }
//    }
//    ack_response* createIndex(char* index, query_options* options) {
//        static ack_response res;
//        int err = kuzzle_wrapper_create_index($self, &res, index, options);
//
//        if (err == 0) {
//            return &res;
//        }
//        return (void*)0;
//    }
//    ack_response* createIndex(char* index) {
//        static ack_response res;
//        int err = kuzzle_wrapper_create_index($self, &res, index, (void*)0);
//
//        if (err == 0) {
//            return &res;
//        }
//        return (void*)0;
//    }
//
//    // createMyCredentials
//    %exception createMyCredentials {
//        $action
//        if (result == $null) {
//            jclass clazz = (*jenv)->FindClass(jenv, "java/lang/IllegalArgumentException");
//            (*jenv)->ThrowNew(jenv, clazz, "Kuzzle.CreateMyCredentials: strategy is required");
//            return $null;
//        }
//    }
//    JsonObject* createMyCredentials(char* strategy, JsonObject* credentials, query_options* options) {
//        static json_result res;
//        static JsonObject ret;
//        int err = kuzzle_wrapper_create_my_credentials($self, &res, strategy, credentials->jobj, options);
//
//        if (err == 0) {
//            ret.jobj = res.result;
//        } else {
//            ret.jobj = json_tokener_parse(res.error);
//        }
//        return &ret;
//    }
//    JsonObject* createMyCredentials(char* strategy, JsonObject* credentials) {
//        static json_result res;
//        static JsonObject ret;
//        int err = kuzzle_wrapper_create_my_credentials($self, &res, strategy, credentials->jobj, (void*)0);
//
//        if (err == 0) {
//            ret.jobj = res.result;
//        } else {
//            ret.jobj = json_tokener_parse(res.error);
//        }
//        return &ret;
//    }
}

