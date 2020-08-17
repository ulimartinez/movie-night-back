/**
 * Create the store with dynamic reducers
 */

import { createStore, applyMiddleware } from 'redux';
import createSagaMiddleware from 'redux-saga';
import { composeWithDevTools } from 'redux-devtools-extension/logOnlyInProduction';
import persistState from 'redux-localstorage'

import reducer from '../reducers/index';
import { Creators as AuthActions } from '../reducers/auth';
import saga from '../sagas';

// Saga middleware
const sagaMiddleware = createSagaMiddleware();

// Add middleware enhancers
const composeEnhancers = composeWithDevTools({
    actionCreators: {
        ...AuthActions
    }
});

export default function configureStore () {
    // Build Redux Store
    const store = createStore(
        reducer,
        composeEnhancers(
            // Apply saga middleware
            applyMiddleware(sagaMiddleware),
            //apply persist enhancer
            persistState(["auth", "groups"])
        )
    );
    // run the root saga
    sagaMiddleware.run(saga);
    return store;
}