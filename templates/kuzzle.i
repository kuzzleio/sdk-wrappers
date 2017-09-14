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
}