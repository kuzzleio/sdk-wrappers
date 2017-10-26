package main

/*
	#cgo CFLAGS: -I../../headers
	#include <errno.h>
	#include <stdlib.h>
	#include "kuzzle.h"
	#include "sdk_wrappers_internal.h"

	void setErrno(int err) {
		errno = err;
	}
 */
import "C"

import (
	"unsafe"

	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

func kuzzle_wrapper_security_new_profile(k *C.kuzzle, id *C.char, policies **C.policy) *C.profile {
	cprofile := (*C.profile)(C.calloc(1, C.sizeof_profile))
	cprofile.id = id
	cprofile.policies = policies
	cprofile.kuzzle = k

	return cprofile
}

func kuzzle_wrapper_security_destroy_profile(p *C.profile) {
	if p == nil {
		return
	}

	if p.policies != nil {
		policies := (*[1<<30 - 1]*C.policy)(unsafe.Pointer(p.policies))[:p.length]
		for _, policy := range policies {
			C.free(unsafe.Pointer(policy.role_id))
			restrictions := (*[1<<30 - 1]*C.policy_restriction)(unsafe.Pointer(policy.restricted_to))[:policy.length]
			for _, restriction := range restrictions {
				C.free(unsafe.Pointer(restriction.index))
				collections := (*[1<<30 - 1]*C.char)(unsafe.Pointer(restriction.collections))[:restriction.length]
				for _, collection := range collections {
					C.free(unsafe.Pointer(collection))
				}
				C.free(unsafe.Pointer(restriction.collections))
			}
			C.free(unsafe.Pointer(policy.restricted_to))
		}
		C.free(unsafe.Pointer(p.policies))
	}

	C.free(unsafe.Pointer(p))
}

func kuzzle_wrapper_security_fetch_profile(k *C.kuzzle, id *C.char, o *C.query_options) *C.profile_result {
	result := (*C.profile_result)(C.calloc(1, C.sizeof_profile_result))
	options := SetQueryOptions(o)

	profile, err := (*kuzzle.Kuzzle)(k.instance).Security.FetchProfile(C.GoString(id), options)
	if err != nil {
		Set_profile_result_error(result, err)
		return result
	}

	result.profile = goToCProfile(k, profile)

	return result
}

func kuzzle_wrapper_security_scroll_profiles(k *C.kuzzle, s *C.char, o *C.query_options) *C.search_profiles_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.ScrollProfiles(C.GoString(s), options)

	return goToCProfileSearchResult(k, res, err)
}

func kuzzle_wrapper_security_search_profiles(k *C.kuzzle, f *C.search_filters, o *C.query_options) *C.search_profiles_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.SearchProfiles(cToGoSearchFilters(f), options)

	return goToCProfileSearchResult(k, res, err)
}

func kuzzle_wrapper_security_new_role(k *C.kuzzle, id *C.char, c *C.controllers) *C.role {
	crole := (*C.role)(C.calloc(1, C.sizeof_role))
	crole.id = id
	crole.controllers = c
	crole.kuzzle = k

	_, err := cToGoRole(crole)
	if err != nil {
		C.setErrno(C.ENOKEY)
		return nil
	}

	return crole
}

func kuzzle_wrapper_security_destroy_role(r *C.role) {
	if r == nil {
		return
	}

	C.json_object_put(r.controllers)
	C.free(unsafe.Pointer(r))
}

func kuzzle_wrapper_security_fetch_role(k *C.kuzzle, id *C.char, o *C.query_options) *C.role_result {
	result := (*C.role_result)(C.calloc(1, C.sizeof_role_result))
	options := SetQueryOptions(o)

	role, err := (*kuzzle.Kuzzle)(k.instance).Security.FetchRole(C.GoString(id), options)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}

	result.role = goToCRole(k, role)

	return result
}

func kuzzle_wrapper_security_search_roles(k *C.kuzzle, f *C.search_filters, o *C.query_options) *C.search_roles_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.SearchRoles(cToGoSearchFilters(f), options)

	return goToCRoleSearchResult(k, res, err)
}

func kuzzle_wrapper_security_role_delete(r *C.role, o *C.query_options) *C.string_result {
	result := (*C.string_result)(C.calloc(1, C.sizeof_string_result))
	opts := SetQueryOptions(o)

	role, err := cToGoRole(r)
	if err != nil {
		Set_string_result_error(result, err)
		return result
	}
	res, err := role.Delete(opts)
	if err != nil {
		Set_string_result_error(result, err)
		return result
	}

	result.result = C.CString(res)

	return result
}

func kuzzle_wrapper_security_role_save(r *C.role, o *C.query_options) *C.role_result {
	result := (*C.role_result)(C.calloc(1, C.sizeof_role_result))
	options := SetQueryOptions(o)

	role, err := cToGoRole(r)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}
	res, err := role.Save(options)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}

	result.role = goToCRole(r.kuzzle, res)

	return result
}

func kuzzle_wrapper_security_role_update(r *C.role, c *C.controllers, o *C.query_options) *C.role_result {
	result := (*C.role_result)(C.calloc(1, C.sizeof_role_result))
	options := SetQueryOptions(o)

	role, err := cToGoRole(r)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}

	var controllers *types.Controllers
	if c != nil {
		controllers, _ = cToGoControllers(c)
	}
	res, err := role.Update(controllers, options)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}

	result.role = goToCRole(r.kuzzle, res)

	return result
}

