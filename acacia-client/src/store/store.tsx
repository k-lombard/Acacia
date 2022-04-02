import { initialAcaciaState, AcaciaReducer } from "./Acacia-reducers";
import { combineReducers, applyMiddleware, createStore, compose } from 'redux'
import { createStoreHook } from "react-redux";
import { persistStore, persistReducer } from 'redux-persist'
import storage from 'redux-persist/lib/storage'
import thunk from 'redux-thunk'
import { PersistPartial } from "redux-persist/lib/persistReducer";
import createSagaMiddleware from "redux-saga";

export const initialState = {
    Acacia: initialAcaciaState
}


export type RootState = ReturnType<typeof rootReducer>
const middleware = createSagaMiddleware()

const rootReducer = combineReducers({
    Acacia: AcaciaReducer
})

const config = {
    key: 'root',
    storage: storage,
};
const persisted = persistReducer<RootState, any>(config, rootReducer);
const store = createStore(persisted, applyMiddleware(thunk, middleware));

export default store
  


// export default createStore(
//     rootReducer,
//     initialState,
//     applyMiddleware(thunk, middleware)
// )