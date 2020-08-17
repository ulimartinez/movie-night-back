import { createActions, createReducer } from 'reduxsauce';
import * as Immutable from 'seamless-immutable';
import { fatalFailure, nightsRequest, nightsSuccess } from '../handlers/nights';


export const INITIAL_STATE = {
    fetching: false,
    nights: undefined,
    history: undefined,
    error: undefined
};

export const {
    Types,
    Creators
} = createActions({
    // Auth Actions
    get_nights_request: ['data'],
    get_nights_success: ['payload'],
    get_nights_failure: ['payload'],
    get_history_request: ['data'],
    get_history_success: ['payload'],
    get_history_failure: ['payload'],
    add_night_request: ['data'],
    add_night_success: ['payload'],
    add_night_failure: ['payload'],
    assign_movie_request: ['data'],
    assign_movie_success: ['payload'],
    assign_movie_failure: ['payload'],
});

export default createReducer(INITIAL_STATE, {
    // Auth reducers
    [Types.GET_NIGHTS_REQUEST]: nightsRequest,
    [Types.GET_NIGHTS_SUCCESS]: nightsSuccess,
    [Types.GET_NIGHTS_FAILURE]: fatalFailure,
    [Types.GET_HISTORY_REQUEST]: nightsRequest,
    [Types.GET_HISTORY_SUCCESS]: nightsSuccess,
    [Types.GET_HISTORY_FAILURE]: fatalFailure,
    [Types.ADD_NIGHT_REQUEST]: nightsRequest,
    [Types.ADD_NIGHT_SUCCESS]: nightsSuccess,
    [Types.ADD_NIGHT_FAILURE]: fatalFailure,
    [Types.ASSIGN_MOVIE_REQUEST]: nightsRequest,
    [Types.ASSIGN_MOVIE_SUCCESS]: nightsSuccess,
    [Types.ASSIGN_MOVIE_FAILURE]: fatalFailure,
});