package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include "sdk_wrappers_internal.h"

	typedef struct {
	} _result;
*/
import "C"
import (
	"encoding/json"
	"unsafe"

	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/security"
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

func goToCPolicyRestriction(restriction *types.PolicyRestriction) *C.policy_restriction {
	crestriction := (*C.policy_restriction)(C.calloc(1, C.sizeof_policy_restriction))
	crestriction.index = C.CString(restriction.Index)
	crestriction.length = C.int(len(restriction.Collections))

	if restriction.Collections != nil {
		crestriction.collections = (**C.char)(C.calloc(C.size_t(len(restriction.Collections)), C.sizeof_char_ptr))
		collections := (*[1<<30 - 1]*C.char)(unsafe.Pointer(crestriction.collections))[:len(restriction.Collections)]

		for i, collection := range restriction.Collections {
			collections[i] = C.CString(collection)
		}
	}

	return crestriction
}

func goToCPolicy(policy *types.Policy) *C.policy {
	cpolicy := (*C.policy)(C.calloc(1, C.sizeof_policy))
	cpolicy.role_id = C.CString(policy.RoleId)
	cpolicy.length = C.int(len(policy.RestrictedTo))

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
	cprofile.length = C.int(len(profile.Policies))
	cprofile.kuzzle = k

	if profile.Policies != nil {
		cprofile.policies = (**C.policy)(C.calloc(C.size_t(len(profile.Policies)), C.sizeof_policy_ptr))
		policies := (*[1<<30 -1]*C.policy)(unsafe.Pointer(cprofile.policies))[:len(profile.Policies)]
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
		cArray := (*[1<<30 -1]*C.role)(unsafe.Pointer(result.result.hits))[:len(res.Hits):len(res.Hits)]

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
	result.result.length = C.int(len(goRes.Hits))
	result.result.total = C.int(goRes.Total)
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

// TODO
func goToCMapping(goMapping *collection.Mapping) *C.mapping {
	result := (*C.mapping)(C.calloc(1, C.sizeof_mapping))

	// TODO

	return result
}

func goToCRole(k *C.kuzzle, role *security.Role) (*C.role) {
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

/*
  TODO: Must be re-done
func goToCSpecificationSearchResult(goRes *types.KuzzleSpecificationSearchResult, cRes *C.specification_search_result) {
	cRes.result.total = C.int(goRes.Total)

	if len(goRes.Hits) > 0 {
		hits := make([]*C.specification, len(goRes.Hits) + 1)

		for i := 0; i < len(goRes.Hits); i++ {
			var spec C.specification
			// TODO register it in global
			t := goRes.Hits[i]
			spec.instance = unsafe.Pointer(&t)
			hits[i] = &spec
		}
		hits[len(goRes.Hits)] = nil

		cRes.result.hits = &hits[0]
	}
}

 */
