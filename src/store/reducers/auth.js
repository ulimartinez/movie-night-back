import { createActions, createReducer } from 'reduxsauce';
import * as Immutable from 'seamless-immutable';
import { fatalFailure, logout, loginRequest, loginSuccess } from '../handlers/auth';


export const INITIAL_STATE = {
    fetching: false,
    User: undefined,
    error: undefined,
    connected: false,
    snack: undefined
};

export const {
    Types,
    Creators
} = createActions({
    // Auth Actions
    auth_request: ['data'],
    auth_success: ['payload'],
    auth_failure: ['payload'],
    register_request: ['data'],
    register_success: ['payload'],
    register_failure: ['payload'],
    logout: undefined
});

export default createReducer(INITIAL_STATE, {
    // Auth reducers
    [Types.AUTH_REQUEST]: loginRequest,
    [Types.AUTH_SUCCESS]: loginSuccess,
    [Types.AUTH_FAILURE]: fatalFailure,
    [Types.REGISTER_REQUEST]: loginRequest,
    [Types.REGISTER_SUCCESS]: loginSuccess,
    [Types.REGISTER_FAILURE]: fatalFailure,
    [Types.LOGOUT]: logout
});