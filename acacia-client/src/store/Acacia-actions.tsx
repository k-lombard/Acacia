import { ThunkAction } from 'redux-thunk'
import { AnyAction } from 'redux'
import { SET_CURRENT_IMAGES } from './Acacia-types'
import { SetCurrentImagesAction } from './Acacia-types'
import { SET_LOADING} from './Acacia-types'
import { SET_LIKES } from './Acacia-types'
import { SetLikesAction } from './Acacia-types'
import { SetLoadingAction } from './Acacia-types'
import { persistStore } from 'redux-persist'
import store from './store'
import { SetCopyImagesAction } from './Acacia-types'
import { SET_COPY_IMAGES } from './Acacia-types'
import { SET_COPY_LIKES } from './Acacia-types'
import { SetCopyLikesAction } from './Acacia-types'
import { SetLikeDictAction } from './Acacia-types'
import { SET_LIKE_DICT } from './Acacia-types'

export function getAllSentries(): ThunkAction<Promise<any>, {}, {}, AnyAction> {
    const requestOptions = {
        method: 'GET',
        headers: new Headers({ 'Content-Type': 'application/json', 'Accept': 'application/json'}),
        
    };
    return async (): Promise<any> => {
        return fetch(`/api/sentries`, requestOptions).then(
            res => {
                return Promise.resolve(res.json())
            }
        ).catch(err => {
            return Promise.reject(err)
        })
    }
}

export function getImagesBySentryId(sentry_id: string): ThunkAction<Promise<any>, {}, {}, AnyAction> {
    const requestOptions = {
        method: 'GET',
        headers: new Headers({ 'Content-Type': 'application/json', 'Accept': 'application/json', 'Access-Control-Allow-Credentials': 'true', "Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST,PUT", "Access-Control-Allow-Headers": "Access-Control-Allow-Headers, Origin, X-Requested-With, Content-Type, Accept, Authorization"}),
        
    };
    return async (): Promise<any> => {
        return fetch(`/api/images/sentry/${sentry_id}`, requestOptions).then(
            res => {
                return Promise.resolve(res.json())
            }
        ).catch(err => {
            return Promise.reject(new Error(err))
        })
    }
}

export function setCurrentImages (currentImages: any): SetCurrentImagesAction {
    return {
        type: SET_CURRENT_IMAGES,
        currentImages
    }
}

export function setLoading (loading: boolean): SetLoadingAction {
    return {
        type: SET_LOADING,
        loading
    }
}

export function setLikes (likes: number[]): SetLikesAction {
    return {
        type: SET_LIKES,
        likes
    }
}

export function setCopyImages (copyImages: any): SetCopyImagesAction {
    return {
        type: SET_COPY_IMAGES,
        copyImages
    }
}

export function setCopyLikes (copyLikes: number[]): SetCopyLikesAction {
    return {
        type: SET_COPY_LIKES,
        copyLikes
    }
}

export function setLikeDict (likeDict: any): SetLikeDictAction {
    return {
        type: SET_LIKE_DICT,
        likeDict
    }
}