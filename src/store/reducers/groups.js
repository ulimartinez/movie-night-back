import { createActions, createReducer } from 'reduxsauce';
import * as Immutable from 'seamless-immutable';
import { membersRequest, membersSuccess, fatalFailure, groupsRequest, groupsSuccess, addRequest, addSuccess, setSuccess, logout } from '../handlers/groups';


export const INITIAL_STATE = {
    currentGroup: undefined,
    groups: undefined,
    members: undefined,
    fetching: false
};

export const {
    Types,
    Creators
} = createActions({
    get_members_request: ['data'],
    get_members_success: ['payload'],
    get_members_failure: ['payload'],
    set_groups_request: ['data'],
    set_groups_success: ['payload'],
    set_groups_failure: ['payload'],
    groups_request: ['data'],
    groups_success: ['payload'],
    groups_failure: ['payload'],
    add_request: ['data'],
    add_success: ['payload'],
    add_failure: ['payload'],
    join_group_request: ['data'],
    join_group_success: ['payload'],
    join_group_failure: ['payload'],
    logout: undefined
});

export default createReducer(INITIAL_STATE, {
    [Types.GET_MEMBERS_REQUEST]: membersRequest,
    [Types.GET_MEMBERS_SUCCESS]: membersSuccess,
    [Types.GET_MEMBERS_FAILURE]: fatalFailure,
    [Types.SET_GROUPS_REQUEST]: groupsRequest,
    [Types.SET_GROUPS_SUCCESS]: setSuccess,
    [Types.SET_GROUPS_FAILURE]: fatalFailure,
    [Types.JOIN_GROUP_REQUEST]: groupsRequest,
    [Types.JOIN_GROUP_SUCCESS]: groupsSuccess,
    [Types.JOIN_GROUP_FAILURE]: fatalFailure,
    [Types.GROUPS_REQUEST]: groupsRequest,
    [Types.GROUPS_SUCCESS]: groupsSuccess,
    [Types.GROUPS_FAILURE]: fatalFailure,
    [Types.ADD_REQUEST]: addRequest,
    [Types.ADD_SUCCESS]: addSuccess,
    [Types.ADD_FAILURE]: fatalFailure,
    [Types.LOGOUT]: logout
});