import "regenerator-runtime/runtime";
import {
    call,
    put,
    takeEvery,
    takeLatest,
    all
} from 'redux-saga/effects';

import {
    Types as AuthTypes,
    Creators as AuthActions
} from './reducers/auth';

import {
    Types as GroupsTypes,
    Creators as GroupsActions
} from './reducers/groups';

import {
    Types as MovieTypes,
    Creators as MovieActions
} from './reducers/movies';

import {
    Types as NightsTypes,
    Creators as NightsActions
} from './reducers/nights';

import createAPI from './api';
import omdbAPI from './omdb';

// Construct API
const api = createAPI(window.location.origin + "/api");
const ombd = omdbAPI("http://omdbapi.com");

/*------------------Auth sagas ------------------------*/
function * login(api, action){
    console.log("the login saga ran")
    try{
        let response = yield call(api.login, action);
        if(response && response.ok){
            console.log(response);
            yield put(AuthActions.auth_success(response.data));
            localStorage.setItem("token", response.data.user.token);
        }
        else {
            console.log("login not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(AuthActions.auth_failure(e.toString()));
    }
}
function * register(api, action){
    try{
        let response = yield call(api.register, action);
        if(response && response.ok){
            yield put(AuthActions.register_success(response.data));
            let response2 = yield call(api.createGroup, {data:{name: "default", token: response.data.user.token}});
        }
        else {
            console.log("register not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(AuthActions.register_failure(e.toString()));
    }
}

/*--------------------- groups sagas --------------------*/

function * getGroups(api, action){
    try{
        let response = yield call(api.getGroups, action);
        if(response && response.ok){
            console.log(response);
            yield put(GroupsActions.groups_success(response.data));
        }
        else {
            console.log("get groups not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(GroupsActions.groups_failure(e.toString()));
    }
}
function * createGroup(api, action){
    try{
        let response = yield call(api.createGroup, action);
        if(response && response.ok){
            console.log(response);
            yield put(GroupsActions.add_success(response.data));
            yield put(GroupsActions.groups_request({token: action.token}))
        }
        else {
            console.log("create group not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(GroupsActions.add_failure(e.toString()));
    }
}
function * getMembers(api, action){
    try {
        let response = yield call(api.getMembers, action);
        if(response && response.ok){
            console.log(response);
            yield put(GroupsActions.get_members_success(response.data));
        }
        else {
            console.log("get members not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(GroupsActions.get_members_failure(e.toString()));
    }
}
function * setGroup(api, action){
    try {
        let response = yield call(api.setGroup, action);
        if(response && response.ok){
            yield put(GroupsActions.set_groups_success(response.data));
        }
        else {
            console.log("set group not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(GroupsActions.set_groups_failure(e.toString()));
    }
}
function * joinGroup(api, action){
    try {
        let response = yield call(api.joinGroup, action);
        if(response && response.ok){
            yield put(GroupsActions.join_group_success(response.data));
        }
        else {
            console.log("join group not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(GroupsActions.join_group_failure(e.toString()));
    }
}

/*--------------------- movies sagas -----------------*/
function * addMovie(api, action){
    try {
        let response = yield call(api.addMovie, action);
        if(response && response.ok){
            console.log("the action for add movie is");
            console.log(action);
            yield put(MovieActions.add_movie_success(response.data));
            yield put(MovieActions.get_movies_request({"id":action.data.group, "user": {"token": action.data.token}}));
        }
        else {
            console.log("add movie not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(MovieActions.add_movie_failure(e.toString()));
    }
}
function * voteMovie(api, action){
    try {
        let response = yield call(api.voteMovie, action);
        if(response && response.ok){
            yield put(MovieActions.vote_success(response.data));
        }
        else {
            console.log("vote movie not ok?");
            yield put(MovieActions.vote_failure(response.data));
        }
    } catch (e) {
        yield put(MovieActions.vote_failure(e.toString()));
    }
}
function * getMovies(api, action){
    try {
        let response = yield call(api.getMovies, action);
        if(response && response.ok){
            let movies = [];
            for(let i = 0; i < response.data.submissions.length; i++){
                let movie = response.data.submissions[i];
                let omdbResponse = yield call(ombd.getMovie, {t: movie.title});
                if(omdbResponse && omdbResponse.ok){
                    movies.push({...movie, ...omdbResponse.data});
                }
                else {
                    console.log("error contacting omdb api");
                    console.log(omdbResponse);
                }
            }
            yield put(MovieActions.get_movies_success({submissions: movies}));
        }
        else {
            console.log("get movie not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(MovieActions.get_movies_failure(e.toString()));
    }
}


/*--------------------- nights sagas -----------------*/
function * getNights(api, action){
    try {
        let response = yield call(api.getNights, action);
        if(response && response.ok){
            console.log("response of get nights");
            console.log(response);

            yield put(NightsActions.get_nights_success({nights: response.data}));
        }
        else {
            console.log("get nights not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(NightsActions.get_nights_failure(e.toString()));
    }
}
function * addNight(api, action){
    try {
        let response = yield call(api.addNight, action);
        if(response && response.ok){
            yield put(NightsActions.add_night_success(response.data));
            yield put(NightsActions.get_nights_request({id: action.data.group, user: {token: action.data.token}}));
        }
        else {
            console.log("add nights not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(NightsActions.add_night_failure(e.toString()));
    }
}
function * getHistory(api, action){
    try {
        let response = yield call(api.getHistory, action);
        if(response && response.ok){
            yield put(NightsActions.get_history_success(response.data));
        }
        else {
            console.log("get history not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(NightsActions.get_history_failure(e.toString()));
    }
}
function * assignMovie(api, action){
    try {
        let response = yield call(api.assignMovie, action);
        if(response && response.ok){
            yield put(NightsActions.assign_movie_success(response.data));
        }
        else {
            console.log("assign movie not ok?");
            console.log(response);
        }
    } catch (e) {
        yield put(NightsActions.assign_movie_failure(e.toString()));
    }
}

/* ------------- Connect Types To Sagas ------------- */
export default function * root () {
    yield all([
        takeLatest(NightsTypes.GET_NIGHTS_REQUEST, getNights, api),
        takeLatest(NightsTypes.GET_HISTORY_REQUEST, getHistory, api),
        takeEvery(NightsTypes.ADD_NIGHT_REQUEST, addNight, api),
        takeLatest(NightsTypes.ASSIGN_MOVIE_REQUEST, assignMovie, api),
        takeEvery(MovieTypes.GET_MOVIES_REQUEST, getMovies, api),
        takeEvery(MovieTypes.ADD_MOVIE_REQUEST, addMovie, api),
        takeLatest(MovieTypes.VOTE_REQUEST, voteMovie, api),
        takeEvery(AuthTypes.REGISTER_REQUEST, register, api),
        takeLatest(AuthTypes.AUTH_REQUEST, login, api),
        takeLatest(GroupsTypes.GROUPS_REQUEST, getGroups, api),
        takeEvery(GroupsTypes.ADD_REQUEST, createGroup, api),
        takeEvery(GroupsTypes.GET_MEMBERS_REQUEST, getMembers, api),
        takeLatest(GroupsTypes.SET_GROUPS_REQUEST, setGroup, api),
        takeLatest(GroupsTypes.JOIN_GROUP_REQUEST, joinGroup, api)
    ]);
}