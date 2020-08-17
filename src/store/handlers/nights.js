import {INITIAL_STATE} from '../reducers/nights';
import moment from 'moment';

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
export const nightsRequest = (state, {data}) => {
    return {
        ...state,
        fetching: true,
        data,
        success: false,
        payload: undefined,
        loadedAt: moment(),
    };
};
export const nightsSuccess = (state, action) => {
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