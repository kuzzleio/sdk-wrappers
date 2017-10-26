%rename(TokenValidity) token_validity;
%rename(AckResponse) ack_response;
%rename(queueTTL) queue_ttl;
%rename(Options) options;
%rename(Kuzzle) kuzzle;
%rename(JsonObject) _json_object;
%rename(JsonResult) json_result;
%rename(LoginResult) login_result;
%rename(BoolResult) bool_result;

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

%ignore _json_object::ptr;
%extend _json_object {
    _json_object() {
        _json_object* j = malloc(sizeof(*j));
        kuzzle_wrapper_json_new(&j->ptr);
        return j;
    }

    ~_json_object() {
        free($self);
    }

    _json_object* put(char* key, char* content) {
        kuzzle_wrapper_json_put($self->ptr, key, content, 0);
        return $self;
    }

    _json_object* put(char* key, int content) {
        kuzzle_wrapper_json_put($self->ptr, key, &content, 1);
        return $self;
    }

    _json_object* put(char* key, double content) {
        kuzzle_wrapper_json_put($self->ptr, key, &content, 2);
        return $self;
    }

    _json_object* put(char* key, bool content) {
        kuzzle_wrapper_json_put($self->ptr, key, &content, 3);
        return $self;
    }

    _json_object* put(char* key, json_object* content) {
        kuzzle_wrapper_json_put($self->ptr, key, content, 4);
        return $self;
    }

    char* getString(char* key) {
        return kuzzle_wrapper_json_get_string($self->ptr, key);
    }

    int getInt(char* key) {
        return kuzzle_wrapper_json_get_int($self->ptr, key);
    }

    double getDouble(char* key) {
        return kuzzle_wrapper_json_get_double($self->ptr, key);
    }

    bool getBoolean(char* key) {
        return kuzzle_wrapper_json_get_bool($self->ptr, key);
    }

    _json_object getJsonObject(char* key) {
        return kuzzle_wrapper_json_get_json_object($self->ptr, key);
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
    kuzzle(char* host) {
        kuzzle *k;
        k = malloc(sizeof(kuzzle));
        kuzzle_wrapper_new_kuzzle(k, host, "websocket", NULL);
        return k;
    }
    ~kuzzle() {
        unregisterKuzzle($self);
        free($self);
    }

    // checkToken
    token_validity* checkToken(char* token) {
        return kuzzle_wrapper_check_token($self, token);
    }

    // connect
    char* connect() {
        return kuzzle_wrapper_connect($self);
    }

    // createIndex
    ack_result* createIndex(char* index, query_options* options) {
        return kuzzle_wrapper_create_index($self, index, options);
    }
    ack_result* createIndex(char* index) {
        return kuzzle_wrapper_create_index($self, index, NULL);
    }

    // createMyCredentials
    json_result* createMyCredentials(char* strategy, _json_object* credentials, query_options* options) {
        return kuzzle_wrapper_create_my_credentials($self, strategy, credentials->ptr, options);
    }
    json_result* createMyCredentials(char* strategy, _json_object* credentials) {
        return kuzzle_wrapper_create_my_credentials($self, strategy, credentials->ptr, NULL);
    }

    // deleteMyCredentials
    ack_result* deleteMyCredentials(char* strategy, query_options *options) {
        return kuzzle_wrapper_delete_my_credentials($self, strategy, options);
    }
    ack_result* deleteMyCredentials(char* strategy) {
        return kuzzle_wrapper_delete_my_credentials($self, strategy, NULL);
    }

    // getMyCredentials
    json_result* getMyCredentials(char *strategy, query_options *options) {
        return kuzzle_wrapper_get_my_credentials($self, strategy, options);
    }
    json_result* getMyCredentials(char *strategy) {
        return kuzzle_wrapper_get_my_credentials($self, strategy, NULL);
    }

    // updateMyCredentials
    json_result* updateMyCredentials(char *strategy, _json_object* credentials, query_options *options) {
        return kuzzle_wrapper_update_my_credentials($self, strategy, credentials->ptr, options);
    }
    json_result* updateMyCredentials(char *strategy, _json_object* credentials) {
        return kuzzle_wrapper_update_my_credentials($self, strategy, credentials->ptr, NULL);
    }

    // validateMyCredentials
    bool_result* validateMyCredentials(char *strategy, _json_object* credentials, query_options* options) {
        return kuzzle_wrapper_validate_my_credentials($self, strategy, credentials->ptr, options);
    }
    bool_result* validateMyCredentials(char *strategy, _json_object* credentials) {
        return kuzzle_wrapper_validate_my_credentials($self, strategy, credentials->ptr, NULL);
    }

    // login
    login_result* login(char* strategy, _json_object* credentials, int expires_in) {
        return kuzzle_wrapper_login($self, strategy, credentials->ptr, &expires_in);
    }
    login_result* login(char* strategy, _json_object* credentials) {
        return kuzzle_wrapper_login($self, strategy, credentials->ptr, NULL);
    }
}
