import {combineReducers} from "redux";

import auth from "./auth"
import groups from "./groups"
import nights from './nights'
import movies from "./movies";

const rootreducer = combineReducers({
    auth,
    nights,
    groups,
    movies
});
export default rootreducer;