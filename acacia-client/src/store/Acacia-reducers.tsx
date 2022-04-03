import { AcaciaState } from "./Acacia-types";
import { AcaciaActionTypes } from "./Acacia-types";
import { SET_CURRENT_IMAGES } from "./Acacia-types";
import { SET_LOADING } from "./Acacia-types";
import { SET_LIKES } from "./Acacia-types";
import { PURGE } from 'redux-persist';
import { initialState } from "./store";
import { SET_COPY_IMAGES } from "./Acacia-types";
import { SET_COPY_LIKES } from "./Acacia-types";
import { SET_LIKE_DICT } from "./Acacia-types";

export const initialAcaciaState: AcaciaState = {
    currentImages: [],
    likes: [],
    loading: false,
    copyImages: [],
    copyLikes: [],
    likeDict: new Map<String, number>()
}

export function AcaciaReducer (
    state = initialAcaciaState,
    action: AcaciaActionTypes
): AcaciaState {
    switch (action.type) {
        case SET_CURRENT_IMAGES:
            return {
                ...state,
                currentImages: action.currentImages
            }
        case SET_LIKES:
            return {
                ...state,
                likes: action.likes
            }
        case SET_LOADING:
            return {
                ...state,
                loading: action.loading
            }
        case SET_COPY_IMAGES:
            return {
                ...state,
                copyImages: action.copyImages
            }
        case SET_COPY_LIKES:
            return {
                ...state,
                copyLikes: action.copyLikes
            }
        case SET_LIKE_DICT:
            return {
                ...state,
                likeDict: action.likeDict
            }
        case PURGE:
            return initialAcaciaState;
        default:
            return state
    }
}