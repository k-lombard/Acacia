import { PURGE } from 'redux-persist'
export interface AcaciaState {
    currentImages: any[];
    likes: number[];
    loading: boolean;
    copyImages: any[];
    copyLikes: number[];
    likeDict: any;
}

export const SET_CURRENT_IMAGES = 'SET_CURRENT_IMAGES'
export const SET_LIKES = 'SET_LIKES'
export const SET_LOADING = 'SET_LOADING'
export const SET_COPY_IMAGES = 'SET_COPY_IMAGES'
export const SET_COPY_LIKES = 'SET_COPY_LIKES'
export const SET_LIKE_DICT = 'LIKE_DICT'

export interface SetCurrentImagesAction {
    type: typeof SET_CURRENT_IMAGES;
    currentImages: any[];
}

export interface SetLikesAction {
    type: typeof SET_LIKES;
    likes: number[];
}

export interface SetLoadingAction {
    type: typeof SET_LOADING;
    loading: boolean;
}
export interface PurgeAction {
    type: typeof PURGE;
    AcaciaState: {};
}

export interface SetCopyImagesAction {
    type: typeof SET_COPY_IMAGES;
    copyImages: any[];
}

export interface SetCopyLikesAction {
    type: typeof SET_COPY_LIKES;
    copyLikes: number[];
}

export interface SetLikeDictAction {
    type: typeof SET_LIKE_DICT;
    likeDict: any;
}


export type AcaciaActionTypes = SetCurrentImagesAction | SetLikesAction | 
SetLoadingAction | PurgeAction | SetCopyImagesAction | SetCopyLikesAction
| SetLikeDictAction