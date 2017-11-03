package main

/*
	#cgo CFLAGS: -I../../headers
	#include <string.h>
	#include "kuzzle.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"encoding/json"
	"unsafe"

	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
)

// Allocates memory
func goToCKuzzleMeta(gMeta *types.Meta) *C.meta {
	if gMeta == nil {
		return nil
	}

	result := (*C.meta)(C.calloc(1, C.sizeof_meta))
	result.author = C.CString(gMeta.Author)
	result.created_at = C.ulonglong(gMeta.CreatedAt)
	result.updated_at = C.ulonglong(gMeta.UpdatedAt)
	result.deleted_at = C.ulonglong(gMeta.DeletedAt)
	result.updater = C.CString(gMeta.Updater)
	result.active = C.bool(gMeta.Active)

	return result
}

// Allocates memory
func goToCShards(gShards *types.Shards) *C.shards {
	if gShards == nil {
		return nil
	}

	result := (*C.shards)(C.calloc(1, C.sizeof_shards))
	result.failed = C.int(gShards.Failed)
	result.successful = C.int(gShards.Successful)
	result.total = C.int(gShards.Total)

	return result
}

// Allocates memory
func goToCDocument(col *C.collection, gDoc *collection.Document) *C.document {
	result := (*C.document)(C.calloc(1, C.sizeof_document))

	result.id = C.CString(gDoc.Id)
	result.index = C.CString(gDoc.Index)
	result.result = C.CString(gDoc.Result)
	result.collection = C.CString(gDoc.Collection)
	result.meta = goToCKuzzleMeta(gDoc.Meta)
	result.shards = goToCShards(gDoc.Shards)
	result._collection = col

	if string(gDoc.Content) != "" {
		buffer := C.CString(string(gDoc.Content))
		result.content = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))
	} else {
		result.content = C.json_object_new_object()
	}

	result.version = C.int(gDoc.Version)
	result.created = C.bool(gDoc.Created)

	return result
}

// Allocates memory
func goToCDocumentResult(col *C.collection, goRes *collection.Document, err error) *C.document_result {
	result := (*C.document_result)(C.calloc(1, C.sizeof_document_result))

	if err != nil {
		Set_document_error(result, err)
		return result
	}

	result.result = goToCDocument(col, goRes)

	return result
}

func goToCPolicyRestriction(restriction *types.PolicyRestriction) *C.policy_restriction {
	crestriction := (*C.policy_restriction)(C.calloc(1, C.sizeof_policy_restriction))
	crestriction.index = C.CString(restriction.Index)
	crestriction.collections_length = C.int(len(restriction.Collections))

	if restriction.Collections != nil {
		crestriction.collections = (**C.char)(C.calloc(C.size_t(len(restriction.Collections)), C.sizeof_char_ptr))
		collections := (*[1<<30 - 1]*C.char)(unsafe.Pointer(crestriction.collections))[:len(restriction.Collections)]

		for i, col := range restriction.Collections {
			collections[i] = C.CString(col)
		}
	}

	return crestriction
}

// Allocates memory
func goToCStringResult(goRes string, err error) *C.string_result {
	result := (*C.string_result)(C.calloc(1, C.sizeof_string_result))

	if err != nil {
		Set_string_result_error(result, err)
		return result
	}

	result.result = C.CString(goRes)

	return result
}

func goToCStringArrayResult(goRes []string, err error) *C.string_array_result {
	result := (*C.string_array_result)(C.calloc(1, C.sizeof_string_array_result))

	if err != nil {
		Set_string_array_result_error(result, err)
		return result
	}

	result.result = (**C.char)(C.calloc(C.size_t(len(goRes)), C.sizeof_char_ptr))
	result.length = C.ulong(len(goRes))

	cArray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(result.result))[:len(goRes):len(goRes)]

	for i, substring := range goRes {
		cArray[i] = C.CString(substring)
	}

	return result
}

// Allocates memory
func goToCIntResult(goRes int, err error) *C.int_result {
	result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

	if err != nil {
		Set_int_result_error(result, err)
		return result
	}

	result.result = C.longlong(goRes)

	return result
}

// Allocates memory
func goToCDoubleResult(goRes float64, err error) *C.double_result {
	result := (*C.double_result)(C.calloc(1, C.sizeof_double_result))

	if err != nil {
		Set_double_result_error(result, err)
		return result
	}

	result.result = C.double(goRes)

	return result
}

// Allocates memory
func goToCBoolResult(goRes bool, err error) *C.bool_result {
	result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))

	if err != nil {
		Set_bool_result_error(result, err)
		return result
	}

	result.result = C.bool(goRes)

	return result
}

func goToCPolicy(policy *types.Policy) *C.policy {
	cpolicy := (*C.policy)(C.calloc(1, C.sizeof_policy))
	cpolicy.role_id = C.CString(policy.RoleId)
	cpolicy.restricted_to_length = C.int(len(policy.RestrictedTo))

	if policy.RestrictedTo != nil {
		cpolicy.restricted_to = (**C.policy_restriction)(C.calloc(C.size_t(len(policy.RestrictedTo)), C.sizeof_policy_restriction_ptr))
		restrictions := (*[1<<30 - 1]*C.policy_restriction)(unsafe.Pointer(cpolicy.restricted_to))[:len(policy.RestrictedTo)]

		for i, restriction := range policy.RestrictedTo {
			restrictions[i] = goToCPolicyRestriction(restriction)
		}
	}

	return cpolicy
}

func goToCProfile(k *C.kuzzle, profile *security.Profile) *C.profile {
	cprofile := (*C.profile)(C.calloc(1, C.sizeof_profile))

	cprofile.id = C.CString(profile.Id)
	cprofile.policies_length = C.int(len(profile.Policies))
	cprofile.kuzzle = k

	if profile.Policies != nil {
		cprofile.policies = (**C.policy)(C.calloc(C.size_t(len(profile.Policies)), C.sizeof_policy_ptr))
		policies := (*[1<<30 - 1]*C.policy)(unsafe.Pointer(cprofile.policies))[:len(profile.Policies)]
		for i, policy := range profile.Policies {
			policies[i] = goToCPolicy(policy)
		}
	}

	return cprofile
}

func goToCProfileSearchResult(k *C.kuzzle, res *security.ProfileSearchResult, err error) *C.search_profiles_result {
	result := (*C.search_profiles_result)(C.calloc(1, C.sizeof_search_profiles_result))

	if err != nil {
		Set_search_profiles_result_error(result, err)
		return result
	}

	result.result = (*C.profile_search)(C.calloc(1, C.sizeof_profile_search))
	result.result.length = C.int(len(res.Hits))
	result.result.total = C.int(res.Total)
	if res.ScrollId != "" {
		result.result.scrollId = C.CString(res.ScrollId)
	}

	if len(res.Hits) > 0 {
		result.result.hits = (**C.profile)(C.calloc(C.size_t(len(res.Hits)), C.sizeof_profile_ptr))
		profiles := (*[1<<30 - 1]*C.profile)(unsafe.Pointer(result.result.hits))[:len(res.Hits)]

		for i, profile := range res.Hits {
			profiles[i] = goToCProfile(k, profile)
		}
	}

	return result
}

func goToCRoleSearchResult(k *C.kuzzle, res *security.RoleSearchResult, err error) *C.search_roles_result {
	result := (*C.search_roles_result)(C.calloc(1, C.sizeof_search_roles_result))

	if err != nil {
		Set_search_roles_result_error(result, err)
		return result
	}

	result.result = (*C.role_search)(C.calloc(1, C.sizeof_role_search))
	result.result.length = C.int(len(res.Hits))
	result.result.total = C.int(res.Total)

	if len(res.Hits) > 0 {
		result.result.hits = (**C.role)(C.calloc(C.size_t(len(res.Hits)), C.sizeof_role_ptr))
		cArray := (*[1<<30 - 1]*C.role)(unsafe.Pointer(result.result.hits))[:len(res.Hits):len(res.Hits)]

		for i, role := range res.Hits {
			cArray[i] = goToCRole(k, role)
		}
	}

	return result
}

// Allocates memory
func goToCSearchResult(col *C.collection, goRes *collection.SearchResult, err error) *C.search_result {
	result := (*C.search_result)(C.calloc(1, C.sizeof_search_result))

	if err != nil {
		Set_search_result_error(result, err)
		return result
	}

	result.result = (*C.document_search)(C.calloc(1, C.sizeof_document_search))
	result.result.length = C.uint(len(goRes.Hits))
	result.result.total = C.uint(goRes.Total)
	if goRes.ScrollId != "" {
		result.result.scrollId = C.CString(goRes.ScrollId)
	}

	if len(goRes.Hits) > 0 {
		result.result.hits = (**C.document)(C.calloc(C.size_t(len(goRes.Hits)), C.sizeof_document_ptr))
		cArray := (*[1<<30 - 1]*C.document)(unsafe.Pointer(result.result.hits))[:len(goRes.Hits):len(goRes.Hits)]

		for i, doc := range goRes.Hits {
			cArray[i] = goToCDocument(col, doc)
		}
	}

	return result
}

// Allocates memory
func goToCMapping(c *C.collection, goMapping *collection.Mapping) *C.mapping {
	result := (*C.mapping)(C.calloc(1, C.sizeof_mapping))

	result.collection = c
	r, _ := json.Marshal(goMapping.Mapping)
	buffer := C.CString(string(r))
	result.mapping = C.json_tokener_parse(buffer)
	C.free(unsafe.Pointer(buffer))

	return result
}

// Allocates memory
func goToCMappingResult(c *C.collection, goRes *collection.Mapping, err error) *C.mapping_result {
	result := (*C.mapping_result)(C.calloc(1, C.sizeof_mapping_result))

	if err != nil {
		Set_mapping_result_error(result, err)
		return result
	}

	result.result = goToCMapping(c, goRes)

	return result
}

func goToCRole(k *C.kuzzle, role *security.Role) *C.role {
	crole := (*C.role)(C.calloc(1, C.sizeof_role))

	crole.id = C.CString(role.Id)
	crole.kuzzle = k

	if role.Controllers != nil {
		j, _ := json.Marshal(role.Controllers)
		buffer := C.CString(string(j))
		crole.controllers = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))
	}

	return crole
}

// Allocates memory
func goToCSpecification(goSpec *types.Specification) *C.specification {
	result := (*C.specification)(C.calloc(1, C.sizeof_specification))

	result.strict = C.bool(goSpec.Strict)

	f, _ := json.Marshal(goSpec.Fields)
	v, _ := json.Marshal(goSpec.Validators)
	bufferFields := C.CString(string(f))
	bufferValidators := C.CString(string(v))

	result.fields = C.json_tokener_parse(bufferFields)
	result.validators = C.json_tokener_parse(bufferValidators)

	C.free(unsafe.Pointer(bufferFields))
	C.free(unsafe.Pointer(bufferValidators))

	return result
}

// Allocates memory
func goToCSpecificationEntry(goEntry *types.SpecificationEntry) *C.specification_entry {
	result := (*C.specification_entry)(C.calloc(1, C.sizeof_specification_entry))
	result.index = C.CString(goEntry.Index)
	result.collection = C.CString(goEntry.Collection)
	result.validation = goToCSpecification(goEntry.Validation)

	return result
}

// Allocates memory
func goToCSpecificationResult(goRes *types.Specification, err error) *C.specification_result {
	result := (*C.specification_result)(C.calloc(1, C.sizeof_specification_result))

	if err != nil {
		Set_specification_result_err(result, err)
		return result
	}

	result.result = goToCSpecification(goRes)

	return result
}

// Allocates memory
func goToCSpecificationSearchResult(goRes *types.SpecificationSearchResult, err error) *C.specification_search_result {
	result := (*C.specification_search_result)(C.calloc(1, C.sizeof_specification_search_result))

	if err != nil {
		Set_specification_search_result_error(result, err)
		return result
	}

	result.result = (*C.specification_search)(C.calloc(1, C.sizeof_specification_search))
	result.result.length = C.uint(len(goRes.Hits))
	result.result.total = C.uint(goRes.Total)
	result.result.scrollId = C.CString(goRes.ScrollId)

	if len(goRes.Hits) > 0 {
		result.result.hits = (**C.specification_entry)(C.calloc(C.size_t(len(goRes.Hits)), C.sizeof_specification_entry_ptr))
		cArray := (*[1<<30 - 1]*C.specification_entry)(unsafe.Pointer(result.result.hits))[:len(goRes.Hits):len(goRes.Hits)]

		for i, spec := range goRes.Hits {
			cArray[i] = goToCSpecificationEntry(&spec.Source)
		}
	}

	return result
}

func goToCJson(data interface{}) (*C.json_object, error) {
	r, err := json.Marshal(data)
	if err != nil {
		return nil, types.NewError(err.Error(), 400)
	}

	buffer := C.CString(string(r))
	defer C.free(unsafe.Pointer(buffer))

	tok := C.json_tokener_new()
	j := C.json_tokener_parse_ex(tok, buffer, C.int(C.strlen(buffer)))
	jerr := C.json_tokener_get_error(tok)
	if jerr != C.json_tokener_success {
		return nil, types.NewError(C.GoString(C.json_tokener_error_desc(jerr)), 400)
	}

	return j, nil
}


func goToCJsonResult(goRes interface{}, err error) *C.json_result {
	result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))

	if err != nil {
		Set_json_result_error(result, err)
		return result
	}

	result.result, err = goToCJson(goRes)
	if err != nil {
		Set_json_result_error(result, err)
		return result
	}

	return result
}

func goToCJsonArrayResult(goRes []interface{}, err error) *C.json_array_result {
	result := (*C.json_array_result)(C.calloc(1, C.sizeof_json_array_result))

	if err != nil {
	  Set_json_array_result_error(result, err)
	  return result
	}

	result.length = C.uint(len(goRes))
	result.result = (**C.json_object)(C.calloc(C.size_t(result.length), C.sizeof_json_object_ptr))
	cArray := (*[1<<30 - 1]*C.json_object)(unsafe.Pointer(result.result))[:len(goRes):len(goRes)]

	for i, res := range(goRes) {
		r, _ := json.Marshal(res)
		buffer := C.CString(string(r))
		cArray[i] = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))
	}

	return result
}

func goToCProfileResult(k *C.kuzzle, res *security.Profile, err error) *C.profile_result {
	result := (*C.profile_result)(C.calloc(1, C.sizeof_profile_result))
	if err != nil {
		Set_profile_result_error(result, err)
		return result
	}

	result.profile = goToCProfile(k, res)
	return result
}

func goToCUserData(data *types.UserData) (*C.user_data, error) {
	if data == nil {
		return nil, nil
	}

	cdata := (*C.user_data)(C.calloc(1, C.sizeof_user_data))

	if data.Content != nil {
		jsonO, err := goToCJson(data.Content)
		if err != nil {
			return nil, err
		}
		cdata.content = jsonO
	}

	if data.ProfileIds != nil {
		cdata.profile_ids_length = C.uint(len(data.ProfileIds))
		cdata.profile_ids = (**C.char)(C.calloc(C.size_t(len(data.ProfileIds)), C.sizeof_char_ptr))
		carray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(cdata.profile_ids))[:len(data.ProfileIds):len(data.ProfileIds)]

		for i, profileId := range data.ProfileIds {
			carray[i] = C.CString(profileId)
		}
	}

	return cdata, nil
}

func goToCUser(k *C.kuzzle, user *security.User) (*C.user, error) {
	if user == nil {
		return nil, nil
	}

	cuser := (*C.user)(C.calloc(1, C.sizeof_user))
	cuser.id = C.CString(user.Id)
	cuser.kuzzle = k

	if user.Content != nil {
		jsonO, err := goToCJson(user.Content)
		if err != nil {
			return nil, err
		}
		cuser.content = jsonO
	}

	if user.ProfileIds != nil {
		cuser.profile_ids_length = C.uint(len(user.ProfileIds))
		cuser.profile_ids = (**C.char)(C.calloc(C.size_t(len(user.ProfileIds)), C.sizeof_char_ptr))
		carray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(cuser.profile_ids))[:len(user.ProfileIds):len(user.ProfileIds)]

		for i, profileId := range user.ProfileIds {
			carray[i] = C.CString(profileId)
		}
	}


	return cuser, nil
}

func goToCUserResult(k *C.kuzzle, user *security.User, err error) *C.user_result {
	result := (*C.user_result)(C.calloc(1, C.sizeof_user_result))
	if err != nil {
		Set_user_result_error(result, err)
		return result
	}

	cuser, err := goToCUser(k, user)
	if err != nil {
		Set_user_result_error(result, err)
		return result
	}

	result.user = cuser

	return result
}

func goToCProfilesResult(k *C.kuzzle, profiles []*security.Profile, err error) *C.profiles_result {
	result := (*C.profiles_result)(C.calloc(1, C.sizeof_profiles_result))
	if err != nil {
		Set_profiles_result_error(result, err)
		return result
	}

	result.profiles_length = C.uint(len(profiles))

	if profiles != nil {
		result.profiles = (**C.profile)(C.calloc(C.size_t(len(profiles)), C.sizeof_profile_ptr))
		carray := (*[1<<30 - 1]*C.profile)(unsafe.Pointer(result.profiles))[:len(profiles):len(profiles)]

		for i, profile := range profiles {
			carray[i] = goToCProfile(k, profile)
		}
	}

	return result
}

func goToCUserRight(right *types.UserRights) *C.user_right {
	if right == nil {
		return nil
	}

	cright := (*C.user_right)(C.calloc(1, C.sizeof_user_right))
	cright.controller = C.CString(right.Controller)
	cright.action = C.CString(right.Action)
	cright.index = C.CString(right.Index)
	cright.collection = C.CString(right.Collection)
	cright.value = C.CString(right.Value)

	return cright
}

func goToCUserRightsResult(rights []*types.UserRights, err error) *C.user_rights_result {
	result := (*C.user_rights_result)(C.calloc(1, C.sizeof_user_rights_result))
	if (err != nil) {
		Set_user_rights_error(result, err)
		return result
	}

	result.user_rights_length = C.uint(len(rights))
	if rights != nil {
		result.user_rights = (**C.user_right)(C.calloc(C.size_t(len(rights)), C.sizeof_user_right_ptr))
		carray := (*[1<<30 - 1]*C.user_right)(unsafe.Pointer(result.user_rights))[:len(rights):len(rights)]

		for i, right := range rights {
			carray[i] = goToCUserRight(right)
		}
	}

	return result
}
