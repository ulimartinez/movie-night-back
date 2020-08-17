import {INITIAL_STATE} from '../reducers/auth';
import moment from 'moment';

/* Request Failure with disconnect */
export const fatalFailure = (state, action) => {
    const {
        payload
    } = action;
    return {
        ...state,
        fetching: false,
        success: false,
        error: true,
        ...payload
    };
};
export const movieRequest = (state, {data}) => {
    return {
        ...state,
        fetching: true,
        data,
        success: false,
        payload: undefined,
        loadedAt: moment(),
        connected: false,
        snack: undefined
    };
};
export const movieSuccess = (state, action) => {
    const {
        payload
    } = action;
    return {
        ...state,
        fetching: false,
        success: true,
        error: false,
        ...payload
    };
};

export const voteSuccess = (state, action) => {
    const {
        payload
    } = action;
    return {
        ...state,
        fetching: false,
        success: true,
        error: false,
        snack: "Voted for movie",
        ...payload
    };
};

export const dismissSnack = (state, action) => {
    const {
        payload
    } = action;
    return {
        ...state,
        error: false,
        snack: undefined,
        errors: undefined,
        ...payload
    }
}