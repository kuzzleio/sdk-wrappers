#include <stdio.h>
#include <json/json.h>
#include <kuzzle.h>

void print_json_value(json_object *jobj) {
    enum json_type type;
    type = json_object_get_type(jobj);
    switch (type) {
        case json_type_boolean:
            printf("json_type_boolean\n");
            printf("value: %s\n", json_object_get_boolean(jobj)? "true": "false");
        break;
        case json_type_double:
            printf("json_type_double\n");
            printf("          value: %lf\n", json_object_get_double(jobj));
        break;
        case json_type_int:
            printf("json_type_int\n");
            printf("          value: %d\n", json_object_get_int(jobj));
        break;
        case json_type_string:
            printf("json_type_string");
            printf("          value: %s\n", json_object_get_string(jobj));
        break;
    }
}

//void json_parse_array( json_object *jobj, char *key) {
//    void json_parse(json_object * jobj);
//    enum json_type type;
//
//    json_object *jarray = jobj;
//    if(key) {
//        json_object_object_get_ex(jobj, key, &jarray);
//    }
//
//    int arraylen = json_object_array_length(jarray);
//    printf("Array Length: %dn",arraylen);
//    int i;
//    json_object * jvalue;
//
//    for (i=0; i< arraylen; i++){
//        jvalue = json_object_array_get_idx(jarray, i);
//        type = json_object_get_type(jvalue);
//        if (type == json_type_array) {
//            json_parse_array(jvalue, NULL);
//            }
//            else if (type != json_type_object) {
//            printf("value[%d]: ",i);
//            print_json_value(jvalue);
//            }
//            else {
//            json_parse(jvalue);
//        }
//    }
//}

void json_parse(json_object * jobj) {
enum json_type type;

json_object_object_foreach(jobj, key, val) {
//        type = json_object_get_type(val);
//        switch (type) {
//            case json_type_boolean:
//            case json_type_double:
//            case json_type_int:
//            case json_type_string:
//                printf("%s\n", key);
//                print_json_value(val);
//            break;
//            case json_type_object:
//                json_object_object_get_ex(jobj, key, &jobj);
//                json_parse(jobj, reader);
//            break;
//            case json_type_array:
//                json_parse_array(jobj, key);
//            break;
//        }
    }
}
