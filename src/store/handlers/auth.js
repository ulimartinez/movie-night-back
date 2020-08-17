import {INITIAL_STATE} from '../reducers/auth';
import moment from 'moment';

export const logout = (state) => {
    return INITIAL_STATE;
};
/* Request Failure with disconnect */
export const fatalFailure = (state, action) => {
    const {
        payload
    } = action;
    return {
        fetching: false,
        success: false,
        error: true,
        ...payload
    };
};
export const silentRequest = (state, {data}) => ({
    ...state,
    ...data,
    success: false,
    payload: undefined,
    loadedAt: moment()
});
export const loginRequest = (state, {data}) => {
    return {
        ...state,
        fetching: true,
        data,
        success: false,
        payload: undefined,
        loadedAt: moment(),
        connected: false
    };
};
export const loginSuccess = (state, action) => {
    const {
        payload
    } = action;
    return {
        ...state,
        fetching: false,
        success: true,
        error: false,
        connected: true,
        ...payload
    };
};