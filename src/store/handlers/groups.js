import {INITIAL_STATE} from '../reducers/groups';
import moment from 'moment';

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

export const groupsRequest = (state, {data}) => {
    return {
        ...state,
        fetching: true,
        data,
        success: false,
        payload: undefined,
        loadedAt: moment(),
    };
};

export const groupsSuccess = (state, action) => {
    const {
        payload
    } = action;
    return {
        ...state,
        fetching: false,
        success: true,
        error: false,
        currentGroup: state.currentGroup? state.currentGroup: payload.groups[0],
        ...payload
    };
};
export const setSuccess = (state, action) => {
    const {
        payload
    } = action;
    return {
        ...state,
        fetching: false,
        success: true,
        error: false,
        currentGroup: payload.group,
        ...payload
    };
};
export const addRequest = (state, {data}) => {
    return {
        ...state,
        fetching: true,
        data,
        success: false,
        payload: undefined,
        loadedAt: moment(),
    };
};
export const addSuccess = (state, action) => {
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

export const membersRequest = (state, {data}) => {
    return {
        ...state,
        fetching: true,
        data,
        success: false,
        payload: undefined,
        loadedAt: moment(),
    };
};

export const membersSuccess = (state, action) => {
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
export const logout = (state, action) => {
    return INITIAL_STATE;
}
