import { createActions, createReducer } from 'reduxsauce';
import * as Immutable from 'seamless-immutable';
import { fatalFailure, movieRequest, movieSuccess, voteSuccess, dismissSnack } from '../handlers/movies';


export const INITIAL_STATE = {
    fetching: false,
    submissions: undefined,
    error: undefined,
    snack: undefined
};

export const {
    Types,
    Creators
} = createActions({
    // Auth Actions
    add_movie_request: ['data'],
    add_movie_success: ['payload'],
    add_movie_failure: ['payload'],
    vote_request: ['data'],
    vote_success: ['payload'],
    vote_failure: ['payload'],
    get_movies_request: ['data'],
    get_movies_success: ['payload'],
    get_movies_failure: ['payload'],
    dismiss_snack: undefined,
});

export default createReducer(INITIAL_STATE, {
    // Auth reducers
    [Types.ADD_MOVIE_REQUEST]: movieRequest,
    [Types.ADD_MOVIE_SUCCESS]: movieSuccess,
    [Types.ADD_MOVIE_FAILURE]: fatalFailure,
    [Types.VOTE_REQUEST]: movieRequest,
    [Types.VOTE_SUCCESS]: voteSuccess,
    [Types.VOTE_FAILURE]: fatalFailure,
    [Types.GET_MOVIES_REQUEST]: movieRequest,
    [Types.GET_MOVIES_SUCCESS]: movieSuccess,
    [Types.GET_MOVIES_FAILURE]: fatalFailure,
    [Types.DISMISS_SNACK]: dismissSnack,
});